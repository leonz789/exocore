// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/prices.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Prices struct {
	TokenId     int32                    `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	NextRoundId uint64                   `protobuf:"varint,2,opt,name=next_round_id,json=nextRoundId,proto3" json:"next_round_id,omitempty"`
	PriceList   []*PriceWithTimeAndRound `protobuf:"bytes,3,rep,name=price_list,json=priceList,proto3" json:"price_list,omitempty"`
}

func (m *Prices) Reset()         { *m = Prices{} }
func (m *Prices) String() string { return proto.CompactTextString(m) }
func (*Prices) ProtoMessage()    {}
func (*Prices) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cc9ccc8f92b87a, []int{0}
}
func (m *Prices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Prices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Prices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Prices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prices.Merge(m, src)
}
func (m *Prices) XXX_Size() int {
	return m.Size()
}
func (m *Prices) XXX_DiscardUnknown() {
	xxx_messageInfo_Prices.DiscardUnknown(m)
}

var xxx_messageInfo_Prices proto.InternalMessageInfo

func (m *Prices) GetTokenId() int32 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *Prices) GetNextRoundId() uint64 {
	if m != nil {
		return m.NextRoundId
	}
	return 0
}

func (m *Prices) GetPriceList() []*PriceWithTimeAndRound {
	if m != nil {
		return m.PriceList
	}
	return nil
}

func init() {
	proto.RegisterType((*Prices)(nil), "exocore.oracle.Prices")
}

func init() { proto.RegisterFile("exocore/oracle/prices.proto", fileDescriptor_50cc9ccc8f92b87a) }

var fileDescriptor_50cc9ccc8f92b87a = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xad, 0xc8, 0x4f,
	0xce, 0x2f, 0x4a, 0xd5, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x4c, 0x4e,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x4a, 0xea, 0x41, 0x24, 0xa5, 0xa4,
	0xb0, 0x29, 0x86, 0xa8, 0x55, 0xea, 0x65, 0xe4, 0x62, 0x0b, 0x00, 0x6b, 0x16, 0x92, 0xe4, 0xe2,
	0x28, 0xc9, 0xcf, 0x4e, 0xcd, 0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62,
	0x07, 0xf3, 0x3d, 0x53, 0x84, 0x94, 0xb8, 0x78, 0xf3, 0x52, 0x2b, 0x4a, 0xe2, 0x8b, 0xf2, 0x4b,
	0xf3, 0x52, 0x40, 0xf2, 0x4c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0xdc, 0x20, 0xc1, 0x20, 0x90, 0x98,
	0x67, 0x8a, 0x90, 0x0b, 0x17, 0x17, 0xd8, 0xe0, 0xf8, 0x9c, 0xcc, 0xe2, 0x12, 0x09, 0x66, 0x05,
	0x66, 0x0d, 0x6e, 0x23, 0x55, 0x3d, 0x54, 0xa7, 0xe8, 0x81, 0xad, 0x0a, 0xcf, 0x2c, 0xc9, 0x08,
	0xc9, 0xcc, 0x4d, 0x75, 0xcc, 0x4b, 0x01, 0x6b, 0x0e, 0xe2, 0x04, 0x6b, 0xf4, 0xc9, 0x2c, 0x2e,
	0x71, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x83, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x57, 0x88, 0xa9, 0x7e, 0xa9, 0x25, 0xe5,
	0xf9, 0x45, 0xd9, 0xfa, 0x30, 0xff, 0x55, 0xc0, 0x7c, 0x58, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0xf6, 0xa2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xea, 0xd2, 0xd3, 0x40, 0x2d, 0x01, 0x00,
	0x00,
}

func (m *Prices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Prices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Prices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PriceList) > 0 {
		for iNdEx := len(m.PriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPrices(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.NextRoundId != 0 {
		i = encodeVarintPrices(dAtA, i, uint64(m.NextRoundId))
		i--
		dAtA[i] = 0x10
	}
	if m.TokenId != 0 {
		i = encodeVarintPrices(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrices(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrices(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Prices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenId != 0 {
		n += 1 + sovPrices(uint64(m.TokenId))
	}
	if m.NextRoundId != 0 {
		n += 1 + sovPrices(uint64(m.NextRoundId))
	}
	if len(m.PriceList) > 0 {
		for _, e := range m.PriceList {
			l = e.Size()
			n += 1 + l + sovPrices(uint64(l))
		}
	}
	return n
}

func sovPrices(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrices(x uint64) (n int) {
	return sovPrices(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Prices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: Prices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Prices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRoundId", wireType)
			}
			m.NextRoundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextRoundId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
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
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceList = append(m.PriceList, &PriceWithTimeAndRound{})
			if err := m.PriceList[len(m.PriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func skipPrices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrices
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
					return 0, ErrIntOverflowPrices
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
					return 0, ErrIntOverflowPrices
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
				return 0, ErrInvalidLengthPrices
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrices
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrices
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrices        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrices          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrices = fmt.Errorf("proto: unexpected end of group")
)
