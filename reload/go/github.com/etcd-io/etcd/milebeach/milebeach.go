package milebeach

import (
	"time"

	"github.com/AleckDarcy/reload/core/tracer"
)

type Messager interface {
	//MessageName() string
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func ReceiveRequest(request Messager) {
	if trace := request.GetTrace(); trace != nil {
		records := trace.GetRecords()
		if count := len(records); count != 0 {
			if lastEvent := records[count-1]; lastEvent.Type == tracer.RecordType_RecordSend {
				records = append(records, &tracer.Record{
					Type:        tracer.RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: lastEvent.MessageName,
					Uuid:        lastEvent.GetUuid(),
					Service:     "todo",
				})

				trace.Records = records
			}
		}
	}
}

func SendResponse(request, response Messager) {
	if trace := response.GetTrace(); trace != nil {
		records := trace.GetRecords()
		if count := len(records); count != 0 {
			if lastEvent := records[count-1]; lastEvent.Type == tracer.RecordType_RecordSend {
				records = append(records, &tracer.Record{
					Type:      tracer.RecordType_RecordReceive,
					Timestamp: time.Now().UnixNano(),
					//MessageName: response.Name(),
					Uuid:    lastEvent.GetUuid(),
					Service: "todo",
				})

				trace.Records = records

				response.SetTrace(trace)
			}
		}
	}
}
