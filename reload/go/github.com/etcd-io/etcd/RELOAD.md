version: 3.4.0

go get git.xxxxx/reload@commit_id

go mod vendor



go mod download

go mod vendor

# ProtoBuf

## 3milebeach payload

V2: todo
V3: .proto file of tracer: github.com/AleckDarcy/reload/core/tracer/message.proto

## add payload to messages

For each request and response (currently for all messages), add field *trace*
V2:
```
optional tracer_v2.Trace trace = 100 [(gogoproto.nullable) = true];
```
V3:
```
tracer.Trace trace = 100 [(gogoproto.nullable) = true];
```

## re-generate pb files

Under project root (github.com/AleckDarcy/reload/reload/go/github.com/etcd-io/etcd), execute

```shell script
./scripts/genproto.sh
```

### Modify imports

V2:
```
import tracer_v2 "github.com/AleckDarcy/reload/core/tracer"
```
V3:
```
import tracer "github.com/AleckDarcy/reload/core/tracer"
```

### Add getter setters

### Extend one-of interface

e.g.,
```
type isResponseOp_Response interface {
	isResponseOp_Response()
	MarshalTo([]byte) (int, error)
	Size() int
	GetFI_Trace() *tracer.Trace // new method
	SetTrace(*tracer.Trace)  // new method
}
```


## (TODO) Extend GRPC/ProtoBuf related interface
