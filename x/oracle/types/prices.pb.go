// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/prices.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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
	TokenID     uint64                   `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	NextRoundID uint64                   `protobuf:"varint,2,opt,name=next_round_id,json=nextRoundId,proto3" json:"next_round_id,omitempty"`
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

func (m *Prices) GetTokenID() uint64 {
	if m != nil {
		return m.TokenID
	}
	return 0
}

func (m *Prices) GetNextRoundID() uint64 {
	if m != nil {
		return m.NextRoundID
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
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xad, 0xc8, 0x4f,
	0xce, 0x2f, 0x4a, 0xd5, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x4c, 0x4e,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x4a, 0xea, 0x41, 0x24, 0xa5, 0xa4,
	0xb0, 0x29, 0x86, 0xa8, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88,
	0xa8, 0xd2, 0x6a, 0x46, 0x2e, 0xb6, 0x00, 0xb0, 0x91, 0x42, 0x6a, 0x5c, 0x1c, 0x25, 0xf9, 0xd9,
	0xa9, 0x79, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x4e, 0xdc, 0x8f, 0xee, 0xc9,
	0xb3, 0x87, 0x80, 0xc4, 0x3c, 0x5d, 0x82, 0xd8, 0xc1, 0x92, 0x9e, 0x29, 0x42, 0xc6, 0x5c, 0xbc,
	0x79, 0xa9, 0x15, 0x25, 0xf1, 0x45, 0xf9, 0xa5, 0x79, 0x29, 0x20, 0xc5, 0x4c, 0x60, 0xc5, 0xfc,
	0x8f, 0xee, 0xc9, 0x73, 0xfb, 0xa5, 0x56, 0x94, 0x04, 0x81, 0xc4, 0x3d, 0x5d, 0x82, 0xb8, 0xf3,
	0xe0, 0x9c, 0x14, 0x21, 0x17, 0x2e, 0x2e, 0xb0, 0x63, 0xe2, 0x73, 0x32, 0x8b, 0x4b, 0x24, 0x98,
	0x15, 0x98, 0x35, 0xb8, 0x8d, 0x54, 0xf5, 0x50, 0x9d, 0xaf, 0x07, 0x76, 0x48, 0x78, 0x66, 0x49,
	0x46, 0x48, 0x66, 0x6e, 0xaa, 0x63, 0x5e, 0x0a, 0x58, 0x73, 0x10, 0x27, 0x58, 0xa3, 0x4f, 0x66,
	0x71, 0x89, 0x93, 0xd7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa4,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xbb, 0x42, 0x4c, 0xf5, 0x4b, 0x2d,
	0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x87, 0x85, 0x49, 0x05, 0x2c, 0x54, 0x4a, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0x01, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8b, 0xfd, 0x9d, 0x61,
	0x01, 0x00, 0x00,
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
	if m.NextRoundID != 0 {
		i = encodeVarintPrices(dAtA, i, uint64(m.NextRoundID))
		i--
		dAtA[i] = 0x10
	}
	if m.TokenID != 0 {
		i = encodeVarintPrices(dAtA, i, uint64(m.TokenID))
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
	if m.TokenID != 0 {
		n += 1 + sovPrices(uint64(m.TokenID))
	}
	if m.NextRoundID != 0 {
		n += 1 + sovPrices(uint64(m.NextRoundID))
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
				return fmt.Errorf("proto: wrong wireType = %d for field TokenID", wireType)
			}
			m.TokenID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRoundID", wireType)
			}
			m.NextRoundID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextRoundID |= uint64(b&0x7F) << shift
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
