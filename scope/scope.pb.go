// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scope.proto

package scope

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/metaverse/truss/deftree/googlethirdparty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UserScopeRequest struct {
}

func (m *UserScopeRequest) Reset()         { *m = UserScopeRequest{} }
func (m *UserScopeRequest) String() string { return proto.CompactTextString(m) }
func (*UserScopeRequest) ProtoMessage()    {}
func (*UserScopeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67276d5d71daf81, []int{0}
}
func (m *UserScopeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserScopeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserScopeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserScopeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserScopeRequest.Merge(m, src)
}
func (m *UserScopeRequest) XXX_Size() int {
	return m.Size()
}
func (m *UserScopeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserScopeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserScopeRequest proto.InternalMessageInfo

type UserScopeResponse struct {
	Code    int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Scopes  []string `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (m *UserScopeResponse) Reset()         { *m = UserScopeResponse{} }
func (m *UserScopeResponse) String() string { return proto.CompactTextString(m) }
func (*UserScopeResponse) ProtoMessage()    {}
func (*UserScopeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67276d5d71daf81, []int{1}
}
func (m *UserScopeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserScopeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserScopeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserScopeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserScopeResponse.Merge(m, src)
}
func (m *UserScopeResponse) XXX_Size() int {
	return m.Size()
}
func (m *UserScopeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserScopeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserScopeResponse proto.InternalMessageInfo

func (m *UserScopeResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserScopeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserScopeResponse) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func init() {
	proto.RegisterType((*UserScopeRequest)(nil), "scope.UserScopeRequest")
	proto.RegisterType((*UserScopeResponse)(nil), "scope.UserScopeResponse")
}

func init() { proto.RegisterFile("scope.proto", fileDescriptor_c67276d5d71daf81) }

var fileDescriptor_c67276d5d71daf81 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x4a, 0xf4, 0x30,
	0x10, 0xc7, 0x9b, 0xdd, 0xaf, 0xfb, 0xd1, 0x78, 0x59, 0x73, 0xd0, 0xb0, 0x87, 0x50, 0x7a, 0xea,
	0xa9, 0x01, 0x7d, 0x03, 0xc1, 0x17, 0xa8, 0x08, 0x7a, 0xcc, 0xb6, 0x63, 0xb7, 0x60, 0x9b, 0x9a,
	0x99, 0x0a, 0x22, 0x5e, 0x7c, 0x02, 0xc1, 0x97, 0xf2, 0xb8, 0xe0, 0xc5, 0xa3, 0xb4, 0x3e, 0x88,
	0x98, 0xaa, 0x88, 0xde, 0xf2, 0xcf, 0x0c, 0xbf, 0xf9, 0xfd, 0xf9, 0x0e, 0x16, 0xb6, 0x83, 0xac,
	0x73, 0x96, 0xac, 0x08, 0x7d, 0x58, 0x1d, 0x57, 0x35, 0x6d, 0xfa, 0x75, 0x56, 0xd8, 0x46, 0x37,
	0x40, 0xe6, 0x1a, 0x1c, 0x82, 0x26, 0xd7, 0x23, 0xea, 0x12, 0x2e, 0xc8, 0x01, 0xe8, 0xca, 0xda,
	0xea, 0x12, 0x68, 0x53, 0xbb, 0xb2, 0x33, 0x8e, 0x6e, 0xb4, 0x69, 0x5b, 0x4b, 0x86, 0x6a, 0xdb,
	0xe2, 0x44, 0x4b, 0x04, 0x5f, 0x9e, 0x22, 0xb8, 0x93, 0x0f, 0x66, 0x0e, 0x57, 0x3d, 0x20, 0x25,
	0xe7, 0x7c, 0xf7, 0xc7, 0x1f, 0x76, 0xb6, 0x45, 0x10, 0x82, 0xff, 0x2b, 0x6c, 0x09, 0x92, 0xc5,
	0x2c, 0x0d, 0x73, 0xff, 0x16, 0x92, 0xff, 0x6f, 0x00, 0xd1, 0x54, 0x20, 0x67, 0x31, 0x4b, 0xa3,
	0xfc, 0x2b, 0x8a, 0x3d, 0xbe, 0xf0, 0x9a, 0x28, 0xe7, 0xf1, 0x3c, 0x8d, 0xf2, 0xcf, 0x74, 0x60,
	0x78, 0xe8, 0xb1, 0xe2, 0x8c, 0x47, 0xdf, 0x37, 0xc4, 0x7e, 0x36, 0x15, 0xfc, 0x6d, 0xb2, 0x92,
	0x7f, 0x07, 0x93, 0x4e, 0x22, 0xef, 0x9f, 0xdf, 0x1e, 0x67, 0x42, 0x2c, 0xb5, 0xdf, 0xd0, 0x3d,
	0x82, 0xd3, 0xb7, 0x75, 0x79, 0x77, 0x24, 0x9f, 0x06, 0xc5, 0xb6, 0x83, 0x62, 0xaf, 0x83, 0x62,
	0x0f, 0xa3, 0x0a, 0xb6, 0xa3, 0x0a, 0x5e, 0x46, 0x15, 0xac, 0x17, 0xbe, 0xf2, 0xe1, 0x7b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x01, 0xdb, 0x3b, 0xe4, 0x4f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScopeClient is the client API for Scope service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScopeClient interface {
	UserScope(ctx context.Context, in *UserScopeRequest, opts ...grpc.CallOption) (*UserScopeResponse, error)
}

type scopeClient struct {
	cc *grpc.ClientConn
}

func NewScopeClient(cc *grpc.ClientConn) ScopeClient {
	return &scopeClient{cc}
}

func (c *scopeClient) UserScope(ctx context.Context, in *UserScopeRequest, opts ...grpc.CallOption) (*UserScopeResponse, error) {
	out := new(UserScopeResponse)
	err := c.cc.Invoke(ctx, "/scope.Scope/UserScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopeServer is the server API for Scope service.
type ScopeServer interface {
	UserScope(context.Context, *UserScopeRequest) (*UserScopeResponse, error)
}

// UnimplementedScopeServer can be embedded to have forward compatible implementations.
type UnimplementedScopeServer struct {
}

func (*UnimplementedScopeServer) UserScope(ctx context.Context, req *UserScopeRequest) (*UserScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserScope not implemented")
}

func RegisterScopeServer(s *grpc.Server, srv ScopeServer) {
	s.RegisterService(&_Scope_serviceDesc, srv)
}

func _Scope_UserScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeServer).UserScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.Scope/UserScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeServer).UserScope(ctx, req.(*UserScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scope_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scope.Scope",
	HandlerType: (*ScopeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserScope",
			Handler:    _Scope_UserScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scope.proto",
}

func (m *UserScopeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserScopeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserScopeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserScopeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintScope(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintScope(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.Scopes) > 0 {
		for _, s := range m.Scopes {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintScope(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserScopeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UserScopeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovScope(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovScope(uint64(l))
	}
	if len(m.Scopes) > 0 {
		for _, s := range m.Scopes {
			l = len(s)
			n += 1 + l + sovScope(uint64(l))
		}
	}
	return n
}

func sovScope(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScope(x uint64) (n int) {
	return sovScope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserScopeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScope
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserScopeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserScopeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScope
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScope
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
func (m *UserScopeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScope
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserScopeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserScopeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scopes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scopes = append(m.Scopes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScope
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScope
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
func skipScope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScope
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
					return 0, ErrIntOverflowScope
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
					return 0, ErrIntOverflowScope
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
			if length < 0 {
				return 0, ErrInvalidLengthScope
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthScope
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowScope
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
				next, err := skipScope(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthScope
				}
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
	ErrInvalidLengthScope = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScope   = fmt.Errorf("proto: integer overflow")
)
