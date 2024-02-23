// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreatePrice struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	//refer to id from Params.TokenList, 0 is reserved, invalid to use
	TokenId int32              `protobuf:"varint,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Prices  []*PriceWithSource `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty"`
	//on which block commit does this message be built on
	BasedBlock uint64 `protobuf:"varint,4,opt,name=based_block,json=basedBlock,proto3" json:"based_block,omitempty"`
}

func (m *MsgCreatePrice) Reset()         { *m = MsgCreatePrice{} }
func (m *MsgCreatePrice) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePrice) ProtoMessage()    {}
func (*MsgCreatePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cf64aff79d2288, []int{0}
}
func (m *MsgCreatePrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePrice.Merge(m, src)
}
func (m *MsgCreatePrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePrice proto.InternalMessageInfo

func (m *MsgCreatePrice) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePrice) GetTokenId() int32 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *MsgCreatePrice) GetPrices() []*PriceWithSource {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *MsgCreatePrice) GetBasedBlock() uint64 {
	if m != nil {
		return m.BasedBlock
	}
	return 0
}

type MsgCreatePriceResponse struct {
}

func (m *MsgCreatePriceResponse) Reset()         { *m = MsgCreatePriceResponse{} }
func (m *MsgCreatePriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePriceResponse) ProtoMessage()    {}
func (*MsgCreatePriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cf64aff79d2288, []int{1}
}
func (m *MsgCreatePriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePriceResponse.Merge(m, src)
}
func (m *MsgCreatePriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePrice)(nil), "exocore.oracle.MsgCreatePrice")
	proto.RegisterType((*MsgCreatePriceResponse)(nil), "exocore.oracle.MsgCreatePriceResponse")
}

func init() { proto.RegisterFile("exocore/oracle/tx.proto", fileDescriptor_02cf64aff79d2288) }

var fileDescriptor_02cf64aff79d2288 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xad, 0xc8, 0x4f,
	0xce, 0x2f, 0x4a, 0xd5, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x4a, 0xe8, 0x41, 0x24, 0xa4, 0xa4, 0xd0, 0x14, 0x16, 0x14,
	0x65, 0x26, 0xa7, 0x42, 0xd4, 0x2a, 0xcd, 0x67, 0xe4, 0xe2, 0xf3, 0x2d, 0x4e, 0x77, 0x2e, 0x4a,
	0x4d, 0x2c, 0x49, 0x0d, 0x00, 0x49, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x83, 0xb8, 0xf9, 0x45, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x24, 0x17, 0x47, 0x49, 0x7e, 0x76, 0x6a,
	0x5e, 0x7c, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x3b, 0x98, 0xef, 0x99, 0x22,
	0x64, 0xce, 0xc5, 0x06, 0x36, 0xb6, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5e, 0x0f,
	0xd5, 0x11, 0x7a, 0x60, 0xb3, 0xc3, 0x33, 0x4b, 0x32, 0x82, 0xf3, 0x4b, 0x8b, 0x92, 0x53, 0x83,
	0xa0, 0xca, 0x85, 0xe4, 0xb9, 0xb8, 0x93, 0x12, 0x8b, 0x53, 0x53, 0xe2, 0x93, 0x72, 0xf2, 0x93,
	0xb3, 0x25, 0x58, 0x14, 0x18, 0x35, 0x58, 0x82, 0xb8, 0xc0, 0x42, 0x4e, 0x20, 0x11, 0x25, 0x09,
	0x2e, 0x31, 0x54, 0x07, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0xc5, 0x70, 0x31,
	0xfb, 0x16, 0xa7, 0x0b, 0x85, 0x72, 0x71, 0x23, 0x3b, 0x5f, 0x0e, 0xdd, 0x66, 0x54, 0xdd, 0x52,
	0x6a, 0xf8, 0xe5, 0x61, 0xa6, 0x3b, 0x79, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x2b, 0xc4,
	0x2c, 0xbf, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0x7d, 0x58, 0x48, 0x57, 0xc0, 0x23, 0xa5, 0xb2,
	0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0xd8, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xb1,
	0x60, 0x86, 0xb3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	//create price for a new oracle round
	CreatePrice(ctx context.Context, in *MsgCreatePrice, opts ...grpc.CallOption) (*MsgCreatePriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePrice(ctx context.Context, in *MsgCreatePrice, opts ...grpc.CallOption) (*MsgCreatePriceResponse, error) {
	out := new(MsgCreatePriceResponse)
	err := c.cc.Invoke(ctx, "/exocore.oracle.Msg/CreatePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	//create price for a new oracle round
	CreatePrice(context.Context, *MsgCreatePrice) (*MsgCreatePriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePrice(ctx context.Context, req *MsgCreatePrice) (*MsgCreatePriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.oracle.Msg/CreatePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePrice(ctx, req.(*MsgCreatePrice))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exocore.oracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrice",
			Handler:    _Msg_CreatePrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exocore/oracle/tx.proto",
}

func (m *MsgCreatePrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BasedBlock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BasedBlock))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.TokenId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TokenId != 0 {
		n += 1 + sovTx(uint64(m.TokenId))
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.BasedBlock != 0 {
		n += 1 + sovTx(uint64(m.BasedBlock))
	}
	return n
}

func (m *MsgCreatePriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, &PriceWithSource{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasedBlock", wireType)
			}
			m.BasedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasedBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
