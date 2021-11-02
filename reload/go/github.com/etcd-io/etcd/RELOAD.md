version: 3.4.0

# Preparation (IMPORTANT)

## Git

Under a random empty folder, clone the repo and checkout to the correct branch:

```go
git clone http://github.com/AleckDarcy/reload.git
cd reload
git checkout feature/etcd_3_4_0
```

Before opening ./ as the root folder of a Goland project, please do the following in terminal:

```shell script
cd reload/go/github.com/etcd-io/etcd
# download dependencies
go mod download
go mod vendor
# reload 3milebeach to vendor
chmod +x reload.sh
./reload.sh
```

Bugs may occur if this is done in Goland's terminal (after opening the project), not sure why it's happening.
One possible reason is that the hacking of dependencies under vendor may confuse Goland.
So do all the tricks before Goland can even react and do stupid things. XD

## Use a specific version of reload repo

```shell script
go get github.com/AleckDarcy/reload@commit_id
go mod download
go mod vendor
(WIP)
```

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
	tracer.Tracer // new method
}
```


## (TODO) Extend GRPC/ProtoBuf related interface
