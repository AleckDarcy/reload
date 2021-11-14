package rafthttp

import (
	"time"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"

	"go.etcd.io/etcd/raft/raftpb"
)

// before sending something
func beforeEncode(plugin *tracer.Plugin, m *raftpb.Message) {
	if trace := m.PrepareTrace().Trace; trace != nil { // 3milebeach todo: deprecate PrepareTrace()
		// 3milebeach note:
		// 1. Send Request: firstEvent is recorded when receiving the upstream request;
		// 2. Send Response: firstEvent is recorded when receiving the corresponding request
		firstEvent, ok := trace.GetFirstEvent()
		if !ok {
			log.Error.PrintlnWithCaller("%s trace with no events", plugin)

			return
		}

		var event *tracer.Record
		msgType, msgName := m.Type.TMBType(), m.Type.String()
		if msgType == tracer.MessageType_Message_Request {
			uuid := tracer.NewUUID()

			event = &tracer.Record{
				Type:        tracer.RecordType_RecordSend,
				Timestamp:   time.Now().UnixNano(),
				MessageName: msgName,
				Uuid:        uuid,
				Service:     plugin.ServerID,
			}

			meta := tracer.NewContextMeta1(trace.Id, uuid, msgName, plugin.ServerID)
			plugin.Store.DoRequest(meta, trace, event, &tracer.Refer{UUID: firstEvent.Uuid})
		} else { // response
			uuid := firstEvent.Uuid

			event = &tracer.Record{
				Type:        tracer.RecordType_RecordSend,
				Timestamp:   time.Now().UnixNano(),
				MessageName: msgName,
				Uuid:        uuid,
				Service:     plugin.ServerID,
			}

			meta := tracer.NewContextMeta1(trace.Id, uuid, msgName, plugin.ServerID)
			plugin.Store.DoResponse(meta, trace, event)
		}

		log.Trace.PrintlnWithCaller("%s capture event: %s", plugin, log.Stringer.JSON(event))
		m.SetFI_Trace(trace)
	} else {
		log.Trace.PrintlnWithCaller("%s no trace", plugin)
	}
}

// after having received something
func afterDecode(plugin *tracer.Plugin, m *raftpb.Message) {
	if trace := m.PrepareTrace().Trace; trace != nil { // 3milebeach todo: deprecate PrepareTrace()
		log.Trace.PrintlnWithCaller("%s lllll", plugin)

		firstEvent, ok := trace.GetFirstEvent()
		if !ok {
			log.Error.PrintlnWithCaller("%s trace with no events", plugin)

			return
		}

		var event *tracer.Record
		msgType, msgName := m.Type.TMBType(), m.Type.String()
		if msgType == tracer.MessageType_Message_Request {
			if events := trace.Records; len(events) != 1 {
				log.Error.PrintlnWithCaller("invalid trace: %s", log.Stringer.JSON(trace))
			} else {
				event = &tracer.Record{
					Type:        tracer.RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: msgName,
					Uuid:        firstEvent.Uuid,
					Service:     plugin.ServerID,
				}
				events[0] = event
				trace.Records = events

				meta := tracer.NewContextMeta1(trace.Id, firstEvent.Uuid, msgName, plugin.ServerID)
				plugin.Store.DoRequest(meta, trace, event, nil)
			}
		} else {
			event = &tracer.Record{
				Type:        tracer.RecordType_RecordReceive,
				Timestamp:   time.Now().UnixNano(),
				MessageName: msgName,
				Uuid:        firstEvent.Uuid,
				Service:     plugin.ServerID,
			}

			meta := tracer.NewContextMeta1(trace.Id, firstEvent.Uuid, msgName, plugin.ServerID)
			plugin.Store.DoResponse(meta, trace, event)
		}

		log.Trace.PrintlnWithCaller("%s capture event: %s from event: %s", plugin, log.Stringer.JSON(event), log.Stringer.JSON(firstEvent))
		m.SetFI_Trace(trace)
	} else {
		log.Trace.PrintlnWithCaller("%s no trace", plugin)
	}
}
