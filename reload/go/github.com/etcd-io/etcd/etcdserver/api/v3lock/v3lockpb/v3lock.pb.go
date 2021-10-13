// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v3lock.proto

/*
	Package v3lockpb is a generated protocol buffer package.

	It is generated from these files:
		v3lock.proto

	It has these top-level messages:
		LockRequest
		LockResponse
		UnlockRequest
		UnlockResponse
*/
package v3lockpb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	etcdserverpb "go.etcd.io/etcd/etcdserver/etcdserverpb"

	tracer "github.com/AleckDarcy/reload/core/tracer"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LockRequest struct {
	// name is the identifier for the distributed shared lock to be acquired.
	Name []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// lease is the ID of the lease that will be attached to ownership of the
	// lock. If the lease expires or is revoked and currently holds the lock,
	// the lock is automatically released. Calls to Lock with the same lease will
	// be treated as a single acquisition; locking twice with the same lease is a
	// no-op.
	Lease int64         `protobuf:"varint,2,opt,name=lease,proto3" json:"lease,omitempty"`
	Trace *tracer.Trace `protobuf:"bytes,100,opt,name=trace" json:"trace,omitempty"`
}

func (m *LockRequest) Reset()                    { *m = LockRequest{} }
func (m *LockRequest) String() string            { return proto.CompactTextString(m) }
func (*LockRequest) ProtoMessage()               {}
func (*LockRequest) Descriptor() ([]byte, []int) { return fileDescriptorV3Lock, []int{0} }

func (m *LockRequest) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *LockRequest) GetLease() int64 {
	if m != nil {
		return m.Lease
	}
	return 0
}

func (m *LockRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

type LockResponse struct {
	Header *etcdserverpb.ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// key is a key that will exist on etcd for the duration that the Lock caller
	// owns the lock. Users should not modify this key or the lock may exhibit
	// undefined behavior.
	Key   []byte        `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Trace *tracer.Trace `protobuf:"bytes,100,opt,name=trace" json:"trace,omitempty"`
}

func (m *LockResponse) Reset()                    { *m = LockResponse{} }
func (m *LockResponse) String() string            { return proto.CompactTextString(m) }
func (*LockResponse) ProtoMessage()               {}
func (*LockResponse) Descriptor() ([]byte, []int) { return fileDescriptorV3Lock, []int{1} }

func (m *LockResponse) GetHeader() *etcdserverpb.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LockResponse) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LockResponse) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

type UnlockRequest struct {
	// key is the lock ownership key granted by Lock.
	Key   []byte        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Trace *tracer.Trace `protobuf:"bytes,100,opt,name=trace" json:"trace,omitempty"`
}

func (m *UnlockRequest) Reset()                    { *m = UnlockRequest{} }
func (m *UnlockRequest) String() string            { return proto.CompactTextString(m) }
func (*UnlockRequest) ProtoMessage()               {}
func (*UnlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorV3Lock, []int{2} }

func (m *UnlockRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *UnlockRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

type UnlockResponse struct {
	Header *etcdserverpb.ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Trace  *tracer.Trace                `protobuf:"bytes,100,opt,name=trace" json:"trace,omitempty"`
}

func (m *UnlockResponse) Reset()                    { *m = UnlockResponse{} }
func (m *UnlockResponse) String() string            { return proto.CompactTextString(m) }
func (*UnlockResponse) ProtoMessage()               {}
func (*UnlockResponse) Descriptor() ([]byte, []int) { return fileDescriptorV3Lock, []int{3} }

func (m *UnlockResponse) GetHeader() *etcdserverpb.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UnlockResponse) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func init() {
	proto.RegisterType((*LockRequest)(nil), "v3lockpb.LockRequest")
	proto.RegisterType((*LockResponse)(nil), "v3lockpb.LockResponse")
	proto.RegisterType((*UnlockRequest)(nil), "v3lockpb.UnlockRequest")
	proto.RegisterType((*UnlockResponse)(nil), "v3lockpb.UnlockResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Lock service

type LockClient interface {
	// Lock acquires a distributed shared lock on a given named lock.
	// On success, it will return a unique key that exists so long as the
	// lock is held by the caller. This key can be used in conjunction with
	// transactions to safely ensure updates to etcd only occur while holding
	// lock ownership. The lock is held until Unlock is called on the key or the
	// lease associate with the owner expires.
	Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	// Unlock takes a key returned by Lock and releases the hold on lock. The
	// next Lock caller waiting for the lock will then be woken up and given
	// ownership of the lock.
	Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error)
}

type lockClient struct {
	cc *grpc.ClientConn
}

func NewLockClient(cc *grpc.ClientConn) LockClient {
	return &lockClient{cc}
}

func (c *lockClient) Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := grpc.Invoke(ctx, "/v3lockpb.Lock/Lock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockClient) Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error) {
	out := new(UnlockResponse)
	err := grpc.Invoke(ctx, "/v3lockpb.Lock/Unlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lock service

type LockServer interface {
	// Lock acquires a distributed shared lock on a given named lock.
	// On success, it will return a unique key that exists so long as the
	// lock is held by the caller. This key can be used in conjunction with
	// transactions to safely ensure updates to etcd only occur while holding
	// lock ownership. The lock is held until Unlock is called on the key or the
	// lease associate with the owner expires.
	Lock(context.Context, *LockRequest) (*LockResponse, error)
	// Unlock takes a key returned by Lock and releases the hold on lock. The
	// next Lock caller waiting for the lock will then be woken up and given
	// ownership of the lock.
	Unlock(context.Context, *UnlockRequest) (*UnlockResponse, error)
}

func RegisterLockServer(s *grpc.Server, srv LockServer) {
	s.RegisterService(&_Lock_serviceDesc, srv)
}

func _Lock_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3lockpb.Lock/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServer).Lock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lock_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3lockpb.Lock/Unlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServer).Unlock(ctx, req.(*UnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v3lockpb.Lock",
	HandlerType: (*LockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lock",
			Handler:    _Lock_Lock_Handler,
		},
		{
			MethodName: "Unlock",
			Handler:    _Lock_Unlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v3lock.proto",
}

func (m *LockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Lease != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Lease))
	}
	if m.Trace != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Trace.Size()))
		n1, err := m.Trace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *LockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Header.Size()))
		n2, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Trace != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Trace.Size()))
		n3, err := m.Trace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *UnlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnlockRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Trace != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Trace.Size()))
		n4, err := m.Trace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *UnlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnlockResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Header.Size()))
		n5, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Trace != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintV3Lock(dAtA, i, uint64(m.Trace.Size()))
		n6, err := m.Trace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func encodeVarintV3Lock(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LockRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovV3Lock(uint64(l))
	}
	if m.Lease != 0 {
		n += 1 + sovV3Lock(uint64(m.Lease))
	}
	if m.Trace != nil {
		l = m.Trace.Size()
		n += 2 + l + sovV3Lock(uint64(l))
	}
	return n
}

func (m *LockResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovV3Lock(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovV3Lock(uint64(l))
	}
	if m.Trace != nil {
		l = m.Trace.Size()
		n += 2 + l + sovV3Lock(uint64(l))
	}
	return n
}

func (m *UnlockRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovV3Lock(uint64(l))
	}
	if m.Trace != nil {
		l = m.Trace.Size()
		n += 2 + l + sovV3Lock(uint64(l))
	}
	return n
}

func (m *UnlockResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovV3Lock(uint64(l))
	}
	if m.Trace != nil {
		l = m.Trace.Size()
		n += 2 + l + sovV3Lock(uint64(l))
	}
	return n
}

func sovV3Lock(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozV3Lock(x uint64) (n int) {
	return sovV3Lock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Lock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lease |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trace == nil {
				m.Trace = &tracer.Trace{}
			}
			if err := m.Trace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Lock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Lock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Lock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &etcdserverpb.ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trace == nil {
				m.Trace = &tracer.Trace{}
			}
			if err := m.Trace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Lock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Lock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnlockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Lock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trace == nil {
				m.Trace = &tracer.Trace{}
			}
			if err := m.Trace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Lock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Lock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnlockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Lock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &etcdserverpb.ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV3Lock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trace == nil {
				m.Trace = &tracer.Trace{}
			}
			if err := m.Trace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Lock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Lock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipV3Lock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowV3Lock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowV3Lock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthV3Lock
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowV3Lock
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipV3Lock(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthV3Lock = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowV3Lock   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v3lock.proto", fileDescriptorV3Lock) }

var fileDescriptorV3Lock = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xbd, 0x4e, 0xeb, 0x30,
	0x18, 0xbd, 0xee, 0x9f, 0xae, 0xdc, 0xf4, 0xde, 0xca, 0xca, 0xbd, 0x44, 0x51, 0x15, 0xaa, 0x0c,
	0xa8, 0x30, 0xc4, 0x52, 0xcb, 0xc4, 0xd8, 0x89, 0xa1, 0x12, 0x52, 0x04, 0x62, 0x76, 0xd2, 0x4f,
	0xa1, 0x6a, 0x1a, 0xa7, 0x71, 0x5a, 0x89, 0x11, 0x5e, 0x81, 0x85, 0xc7, 0xe0, 0x31, 0x3a, 0x22,
	0xb1, 0x23, 0x54, 0x78, 0x10, 0x14, 0xdb, 0x85, 0x00, 0x0b, 0x3f, 0x8b, 0x7d, 0xfc, 0x7d, 0x27,
	0xe7, 0xf8, 0x38, 0x1f, 0x36, 0x96, 0x83, 0x98, 0x87, 0x53, 0x2f, 0xcd, 0x78, 0xce, 0xc9, 0x6f,
	0x75, 0x4a, 0x03, 0xdb, 0x8c, 0x78, 0xc4, 0x65, 0x91, 0x16, 0x48, 0xf5, 0xed, 0x1d, 0xc8, 0xc3,
	0x31, 0x2d, 0x16, 0x01, 0xd9, 0x12, 0xb2, 0x12, 0x4c, 0x03, 0x9a, 0xa5, 0xa1, 0xe6, 0x99, 0x79,
	0xc6, 0x42, 0xc8, 0xe8, 0x0c, 0x84, 0x60, 0x11, 0xe8, 0x6a, 0x27, 0xe2, 0x3c, 0x8a, 0x81, 0xb2,
	0x74, 0x42, 0x59, 0x92, 0xf0, 0x9c, 0xe5, 0x13, 0x9e, 0x08, 0xd5, 0x75, 0x03, 0xdc, 0x1c, 0xf1,
	0x70, 0xea, 0xc3, 0x7c, 0x01, 0x22, 0x27, 0x04, 0xd7, 0x12, 0x36, 0x03, 0x0b, 0x75, 0x51, 0xcf,
	0xf0, 0x25, 0x26, 0x26, 0xae, 0xc7, 0xc0, 0x04, 0x58, 0x95, 0x2e, 0xea, 0x55, 0x7d, 0x75, 0x20,
	0xbb, 0xb8, 0x2e, 0xed, 0xac, 0x71, 0x17, 0xf5, 0x9a, 0xfd, 0x96, 0xa7, 0xcc, 0xbd, 0xe3, 0x62,
	0x1b, 0xd6, 0x56, 0xf7, 0xdb, 0xc8, 0x57, 0x0c, 0xf7, 0x02, 0x61, 0x43, 0x99, 0x88, 0x94, 0x27,
	0x02, 0xc8, 0x3e, 0x6e, 0x9c, 0x01, 0x1b, 0x43, 0x26, 0x7d, 0x9a, 0xfd, 0x8e, 0x57, 0x4e, 0xe4,
	0x6d, 0x78, 0x87, 0x92, 0xe3, 0x6b, 0x2e, 0x69, 0xe3, 0xea, 0x14, 0xce, 0xe5, 0x2d, 0x0c, 0xbf,
	0x80, 0x5f, 0xb9, 0xc3, 0x08, 0xb7, 0x4e, 0x92, 0xb8, 0x94, 0x54, 0xab, 0xa1, 0x6f, 0xa9, 0xcd,
	0xf1, 0x9f, 0x8d, 0xda, 0x8f, 0x22, 0x7d, 0xde, 0xb2, 0x7f, 0x83, 0x70, 0xad, 0x78, 0x44, 0x72,
	0xa4, 0xf7, 0x7f, 0xde, 0x66, 0x6c, 0xbc, 0xd2, 0x1f, 0xb4, 0xff, 0xbf, 0x2f, 0x2b, 0x63, 0xd7,
	0xba, 0xbc, 0x7b, 0xba, 0xaa, 0x10, 0xb7, 0x45, 0x97, 0x03, 0x5a, 0x10, 0xe4, 0x72, 0x80, 0xf6,
	0xc8, 0x29, 0x6e, 0xa8, 0x30, 0x64, 0xeb, 0xf5, 0xdb, 0x37, 0x8f, 0x65, 0x5b, 0x1f, 0x1b, 0x5a,
	0xd6, 0x96, 0xb2, 0xa6, 0xfb, 0xf7, 0x45, 0x76, 0x91, 0x68, 0xe1, 0x61, 0x7b, 0xb5, 0x76, 0xd0,
	0xed, 0xda, 0x41, 0x0f, 0x6b, 0x07, 0x5d, 0x3f, 0x3a, 0xbf, 0x82, 0x86, 0x1c, 0xba, 0xc1, 0x73,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x55, 0x50, 0xe3, 0x00, 0x03, 0x00, 0x00,
}