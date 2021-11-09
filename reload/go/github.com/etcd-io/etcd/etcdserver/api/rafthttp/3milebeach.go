package rafthttp

import (
	"time"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"

	"go.etcd.io/etcd/raft/raftpb"
)

func beforeEncode(plugin *tracer.Plugin, m *raftpb.Message) {
	if trace := m.Trace; trace != nil {
		var uuid string
		if m.Type.TMBType() == tracer.MessageType_Message_Request {
			uuid = tracer.NewUUID()
		} else if lastEvent, ok := trace.GetLastEvent(); ok {
			uuid = lastEvent.Uuid
		} else {
			log.Error.PrintlnWithCaller("%s trace with no events", plugin)

			return
		}

		event := &tracer.Record{
			Type:        tracer.RecordType_RecordSend,
			Timestamp:   time.Now().UnixNano(),
			MessageName: m.Type.String(),
			Uuid:        uuid,
			Service:     plugin.ServerID,
		}

		updateFunction := func(trace *tracer.Trace) {
			trace.Records = append(trace.Records, event)
		}

		meta := tracer.NewContextMeta(trace.Id, uuid, plugin.ServerID)
		if trace, ok := plugin.Store.UpdateFunctionByContextMeta(meta, updateFunction); ok {
			if m.Type.TMBType() == tracer.MessageType_Message_Request {
				// 3milebeach todo: fault injection

				trace = &tracer.Trace{
					Id:      trace.Id,
					Records: []*tracer.Record{event},
					Rlfis:   trace.Rlfis,
					Tfis:    trace.Tfis,
				}
			} else {
				plugin.Store.DeleteByContextMeta(meta)
			}

			log.Debug.PrintlnWithCaller("%s capture event: %s", plugin, log.Stringer.JSON(event))
			m.SetFI_Trace(trace)
		} else {
			log.Error.PrintlnWithCaller("%s UpdateFunctionByContextMeta failed", plugin)
		}
	}
}

func afterDecode(plugin *tracer.Plugin, m *raftpb.Message) {
	if trace := m.Trace; trace != nil {
		if m.Type.TMBType() == tracer.MessageType_Message_Request {
			if events := trace.Records; len(events) != 1 {
				log.Error.PrintlnWithCaller("invalid trace: %s", log.Stringer.JSON(trace))
			} else if lastEvent := events[0]; lastEvent.Uuid == "" {
				log.Error.PrintlnWithCaller("invalid trace: %s", log.Stringer.JSON(trace))
			} else {
				lastEvent.Type = tracer.RecordType_RecordReceive
				lastEvent.Timestamp = time.Now().UnixNano()
				lastEvent.Service = plugin.ServerID

				meta := tracer.NewContextMeta(trace.Id, lastEvent.Uuid, plugin.ServerID)
				plugin.Store.SetByContextMeta(meta, trace)
			}
		} else {
			if events := trace.Records; len(events) == 0 {

			} else if lastEvent := events[0]; lastEvent.Uuid == "" {
				log.Error.PrintlnWithCaller("invalid trace: %s", log.Stringer.JSON(trace))
			} else {
				event := &tracer.Record{
					Type:        tracer.RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: m.Type.String(),
					Uuid:        lastEvent.Uuid,
					Service:     plugin.ServerID,
				}

				trace.Records = append(trace.Records, event)
				log.Debug.PrintlnWithCaller("%s capture event: %s from event: %s", plugin, log.Stringer.JSON(event), log.Stringer.JSON(lastEvent))
				m.SetFI_Trace(trace)
			}
		}
	}
}
