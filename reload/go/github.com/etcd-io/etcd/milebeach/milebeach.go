package milebeach

import (
	"time"

	"github.com/AleckDarcy/reload/core/tracer"
)

type Messager interface {
	GetFI_Name() string
	GetFI_Trace() *tracer.Trace
	SetFI_Trace(*tracer.Trace)
}

func ReceiveRequest(request Messager) {
	if trace := request.GetFI_Trace(); trace != nil {
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
	if trace := response.GetFI_Trace(); trace != nil {
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

				response.SetFI_Trace(trace)
			}
		}
	}
}
