// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rafthttp

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/AleckDarcy/reload/core/log"

	"github.com/AleckDarcy/reload/core/tracer"

	stats "go.etcd.io/etcd/etcdserver/api/v2stats"
	"go.etcd.io/etcd/pkg/pbutil"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft/raftpb"
)

const (
	msgTypeLinkHeartbeat uint8 = 0
	msgTypeAppEntries    uint8 = 1
	msgTypeApp           uint8 = 2

	msgAppV2BufSize = 1024 * 1024
)

// msgappv2 stream sends three types of message: linkHeartbeatMessage,
// AppEntries and MsgApp. AppEntries is the MsgApp that is sent in
// replicate state in raft, whose index and term are fully predictable.
//
// Data format of linkHeartbeatMessage:
// | offset | bytes | description |
// +--------+-------+-------------+
// | 0      | 1     | \x00        |
//
// Data format of AppEntries:
// | offset | bytes | description |
// +--------+-------+-------------+
// | 0      | 1     | \x01        |
// | 1      | 8     | length of entries |
// | 9      | 8     | length of first entry |
// | 17     | n1    | first entry |
// ...
// | x      | 8     | length of k-th entry data |
// | x+8    | nk    | k-th entry data |
// | x+8+nk | 8     | commit index |
//
// Data format of MsgApp:
// | offset | bytes | description |
// +--------+-------+-------------+
// | 0      | 1     | \x02        |
// | 1      | 8     | length of encoded message |
// | 9      | n     | encoded message |
type msgAppV2Encoder struct {
	TMB *tracer.Plugin // 3milebeach

	w  io.Writer
	fs *stats.FollowerStats

	term      uint64
	index     uint64
	buf       []byte
	uint64buf []byte
	uint8buf  []byte
}

func newMsgAppV2Encoder(w io.Writer, fs *stats.FollowerStats) *msgAppV2Encoder {
	return &msgAppV2Encoder{
		w:         w,
		fs:        fs,
		buf:       make([]byte, msgAppV2BufSize),
		uint64buf: make([]byte, 8),
		uint8buf:  make([]byte, 1),
	}
}

func (enc *msgAppV2Encoder) setServerID(id tracer.UUID) { // 3milebeach begins
	enc.TMB = tracer.GetPlugin(id)
} // 3milebeach ends

func (enc *msgAppV2Encoder) encode(m *raftpb.Message) error {
	start := time.Now()
	switch {
	case isLinkHeartbeatMessage(m):
		enc.uint8buf[0] = msgTypeLinkHeartbeat
		if _, err := enc.w.Write(enc.uint8buf); err != nil {
			return err
		}
	case enc.index == m.Index && enc.term == m.LogTerm && m.LogTerm == m.Term:
		enc.uint8buf[0] = msgTypeAppEntries
		if _, err := enc.w.Write(enc.uint8buf); err != nil {
			return err
		}
		// write length of entries
		binary.BigEndian.PutUint64(enc.uint64buf, uint64(len(m.Entries)))
		if _, err := enc.w.Write(enc.uint64buf); err != nil {
			return err
		}
		for i := 0; i < len(m.Entries); i++ {
			// write length of entry
			binary.BigEndian.PutUint64(enc.uint64buf, uint64(m.Entries[i].Size()))
			if _, err := enc.w.Write(enc.uint64buf); err != nil {
				return err
			}
			if n := m.Entries[i].Size(); n < msgAppV2BufSize {
				if _, err := m.Entries[i].MarshalTo(enc.buf); err != nil {
					return err
				}
				if _, err := enc.w.Write(enc.buf[:n]); err != nil {
					return err
				}
			} else {
				if _, err := enc.w.Write(pbutil.MustMarshal(&m.Entries[i])); err != nil {
					return err
				}
			}
			enc.index++
		}
		// write commit index
		binary.BigEndian.PutUint64(enc.uint64buf, m.Commit)
		if _, err := enc.w.Write(enc.uint64buf); err != nil {
			return err
		}

		// 3milebeach note:
		// Unlike rafthttp.messageEncoder's directly using built-in Marshal function, msgAppV2Encoder "manually" encodes
		// essential fields of pb.Message to bytes. The order of encoding determines the order of decoding.
		if trace := m.Trace; trace == nil { // 3milebeach begins
			// write size of trace
			binary.BigEndian.PutUint64(enc.uint64buf, 0)
			if _, err := enc.w.Write(enc.uint64buf); err != nil {
				return err
			}
		} else {
			// 3milebeach todo: capture event && fault injection

			uuid := tracer.NewUUID()
			event := &tracer.Record{
				Type:        tracer.RecordType_RecordSend,
				Timestamp:   time.Now().UnixNano(),
				MessageName: m.Type.String(),
				Uuid:        uuid,
				Service:     enc.TMB.ServerID,
			}

			log.Debug.PrintlnWithCaller("%s capture event: %s", enc.TMB, log.Stringer.JSON(event))

			// write size of trace
			binary.BigEndian.PutUint64(enc.uint64buf, uint64(trace.Size()))
			if _, err := enc.w.Write(enc.uint64buf); err != nil {
				return err
			}
			// write trace
			if _, err := enc.w.Write(pbutil.MustMarshal(trace)); err != nil {
				return err
			}
		} // 3milebeach ends

		enc.fs.Succ(time.Since(start))
	default:
		if err := binary.Write(enc.w, binary.BigEndian, msgTypeApp); err != nil {
			return err
		}
		// write size of message
		if err := binary.Write(enc.w, binary.BigEndian, uint64(m.Size())); err != nil {
			return err
		}
		// write message
		if _, err := enc.w.Write(pbutil.MustMarshal(m)); err != nil {
			return err
		}

		enc.term = m.Term
		enc.index = m.Index
		if l := len(m.Entries); l > 0 {
			enc.index = m.Entries[l-1].Index
		}
		enc.fs.Succ(time.Since(start))
	}
	return nil
}

type msgAppV2Decoder struct {
	TMB *tracer.Plugin // 3milebeach

	r             io.Reader
	local, remote types.ID

	term      uint64
	index     uint64
	buf       []byte
	uint64buf []byte
	uint8buf  []byte
}

func newMsgAppV2Decoder(r io.Reader, local, remote types.ID) *msgAppV2Decoder {
	return &msgAppV2Decoder{
		TMB:       tracer.GetPlugin(local.Decimal()), // 3milebeach
		r:         r,
		local:     local,
		remote:    remote,
		buf:       make([]byte, msgAppV2BufSize),
		uint64buf: make([]byte, 8),
		uint8buf:  make([]byte, 1),
	}
}

func (dec *msgAppV2Decoder) decode() (raftpb.Message, error) {
	var (
		m   raftpb.Message
		typ uint8
	)
	if _, err := io.ReadFull(dec.r, dec.uint8buf); err != nil {
		return m, err
	}
	typ = dec.uint8buf[0]
	switch typ {
	case msgTypeLinkHeartbeat:
		return linkHeartbeatMessage, nil
	case msgTypeAppEntries:
		m = raftpb.Message{
			Type:    raftpb.MsgApp,
			From:    uint64(dec.remote),
			To:      uint64(dec.local),
			Term:    dec.term,
			LogTerm: dec.term,
			Index:   dec.index,
		}

		// decode entries
		if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil {
			return m, err
		}
		l := binary.BigEndian.Uint64(dec.uint64buf)
		m.Entries = make([]raftpb.Entry, int(l))
		for i := 0; i < int(l); i++ {
			if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil {
				return m, err
			}
			size := binary.BigEndian.Uint64(dec.uint64buf)
			var buf []byte
			if size < msgAppV2BufSize {
				buf = dec.buf[:size]
				if _, err := io.ReadFull(dec.r, buf); err != nil {
					return m, err
				}
			} else {
				buf = make([]byte, int(size))
				if _, err := io.ReadFull(dec.r, buf); err != nil {
					return m, err
				}
			}
			dec.index++
			// 1 alloc
			pbutil.MustUnmarshal(&m.Entries[i], buf)
		}
		// decode commit index
		if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil {
			return m, err
		}
		m.Commit = binary.BigEndian.Uint64(dec.uint64buf)

		// 3milebeach note:
		// Unlike rafthttp.messageDecoder's directly using built-in Unmarshal function, msgAppV2Decoder "manually"
		// decodes essential fields of pb.Message out of bytes. The order of decoding should be according to
		// msgAppV2Encoder.
		if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil { // 3milebeach begins
			return m, err
		}

		if size := binary.BigEndian.Uint64(dec.uint64buf); size != 0 {
			buf := dec.buf[:size]
			if _, err := io.ReadFull(dec.r, buf); err != nil {
				return m, err
			}

			trace := &tracer.Trace{}
			pbutil.MustUnmarshal(trace, buf)

			// todo capture event
			if lastEvent, ok := trace.GetLastEvent(); !ok {
				log.Error.PrintlnWithCaller("%s trace with no events", dec.TMB)
			} else {
				event := &tracer.Record{
					Type:        tracer.RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: lastEvent.MessageName,
					Uuid:        lastEvent.Uuid,
					Service:     dec.TMB.ServerID,
				}

				log.Debug.PrintlnWithCaller("%s capture event: %s from event: %s", dec.TMB, log.Stringer.JSON(event), log.Stringer.JSON(lastEvent))
			}

			m.Trace = trace
		} // 3milebeach ends
	case msgTypeApp:
		var size uint64
		if err := binary.Read(dec.r, binary.BigEndian, &size); err != nil {
			return m, err
		}
		buf := make([]byte, int(size))
		if _, err := io.ReadFull(dec.r, buf); err != nil {
			return m, err
		}
		pbutil.MustUnmarshal(&m, buf)

		dec.term = m.Term
		dec.index = m.Index
		if l := len(m.Entries); l > 0 {
			dec.index = m.Entries[l-1].Index
		}
	default:
		return m, fmt.Errorf("failed to parse type %d in msgappv2 stream", typ)
	}
	return m, nil
}
