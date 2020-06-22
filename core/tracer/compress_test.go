package tracer

import (
	"reflect"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
)

const cycles = 64
const records = 4 * cycles

var uuids = make([]UUID, cycles)
var trace = &Trace{}
var flag = false

func init() {
	for i := range uuids {
		uuids[i] = NewUUID()
	}

	trace.Records = make([]*Record, records)
	for i := 0; i < records; i += 4 {
		trace.Records[i] = &Record{

			Type:        RecordType_RecordSend,
			Timestamp:   time.Now().UnixNano(),
			MessageName: "Request1",
			Uuid:        uuids[i/4],
			Service:     "ServiceA",
		}

		trace.Records[i+1] = &Record{
			Type:        RecordType_RecordReceive,
			Timestamp:   time.Now().UnixNano(),
			MessageName: "Request1",
			Uuid:        uuids[i/4],
			Service:     "ServiceB",
		}

		trace.Records[i+2] = &Record{
			Type:        RecordType_RecordSend,
			Timestamp:   time.Now().UnixNano(),
			MessageName: "Response1",
			Uuid:        uuids[i/4],
			Service:     "ServiceB",
		}

		trace.Records[i+3] = &Record{
			Type:        RecordType_RecordReceive,
			Timestamp:   time.Now().UnixNano(),
			MessageName: "Response1",
			Uuid:        uuids[i/4],
			Service:     "ServiceA",
		}
	}

	//trace.Records = trace.Records[:255]
}

func BenchmarkCompress(b *testing.B) {
	//c := make([]byte, 6000)
	for i := 0; i < b.N; i++ {
		//a := make([]byte, 6000)
		//copy(a, c)
		h := Compress.Compress(trace)
		_ = h

		Compress.Decompress(h)

		Compress.Recycle(h)
	}
}

func BenchmarkCodec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := proto.Marshal(trace)

		_ = bytes
		newTrace := &Trace{}
		proto.Unmarshal(bytes, newTrace)
	}
}

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, record := range trace.Records {
			if flag {
				ii, rr := i, record

				_, _ = ii, rr
			}
		}
	}
}

func BenchmarkIteration(b *testing.B) {
	h := Compress.Compress(trace)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h.IterateRecords(func(i int, record *Record) {
			if flag {
				ii, rr := i, record

				_, _ = ii, rr
			}
			//fmt.Println(i, record)
		})
	}
}

func TestCompress_Compress(t *testing.T) {
	uuids := make([]UUID, 10)
	for i := range uuids {
		uuids[i] = NewUUID()
	}

	h := Compress.Compress(trace)

	trace_ := Compress.Decompress(h)

	t.Log(reflect.DeepEqual(trace.Records, trace_.Records))
}
