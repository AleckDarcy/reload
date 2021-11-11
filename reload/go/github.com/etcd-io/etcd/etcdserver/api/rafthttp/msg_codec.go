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
	"errors"
	"io"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"

	"go.etcd.io/etcd/pkg/pbutil"
	"go.etcd.io/etcd/raft/raftpb"
)

// messageEncoder is a encoder that can encode all kinds of messages.
// It MUST be used with a paired messageDecoder.
type messageEncoder struct {
	TMB *tracer.Plugin // 3milebeach

	w io.Writer
}

func (enc *messageEncoder) setServerID(id tracer.UUID) { // 3milebeach begins
	enc.TMB = tracer.GetPlugin(id)
} // 3milebeach ends

func (enc *messageEncoder) encode(m *raftpb.Message) error {
	log.Logger.PrintlnWithCaller("%s message: %s", enc.TMB, log.Stringer.JSON(m))
	if err := binary.Write(enc.w, binary.BigEndian, uint64(m.Size())); err != nil {
		return err
	}

	beforeEncode(enc.TMB, m) // 3milebeach

	_, err := enc.w.Write(pbutil.MustMarshal(m))
	return err
}

// messageDecoder is a decoder that can decode all kinds of messages.
type messageDecoder struct {
	TMB *tracer.Plugin // 3milebeach

	r io.Reader
}

var (
	readBytesLimit     uint64 = 512 * 1024 * 1024 // 512 MB
	ErrExceedSizeLimit        = errors.New("rafthttp: error limit exceeded")
)

func (dec *messageDecoder) decode() (raftpb.Message, error) {
	return dec.decodeLimit(readBytesLimit)
}

func (dec *messageDecoder) decodeLimit(numBytes uint64) (raftpb.Message, error) {
	var m raftpb.Message
	var l uint64
	if err := binary.Read(dec.r, binary.BigEndian, &l); err != nil {
		return m, err
	}
	if l > numBytes {
		return m, ErrExceedSizeLimit
	}
	buf := make([]byte, int(l))
	if _, err := io.ReadFull(dec.r, buf); err != nil {
		return m, err
	}

	err := m.Unmarshal(buf) // 3milebeach begins
	afterDecode(dec.TMB, &m)

	return m, err // 3milebeach ends

	// return m, m.Unmarshal(buf)
}
