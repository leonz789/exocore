// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/round_info.proto

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

type RoundInfo struct {
	TokenId      int32  `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	BasedBlock   uint64 `protobuf:"varint,2,opt,name=based_block,json=basedBlock,proto3" json:"based_block,omitempty"`
	NexitRoundId uint64 `protobuf:"varint,3,opt,name=nexit_roundId,json=nexitRoundId,proto3" json:"nexit_roundId,omitempty"`
	FeederId     int32  `protobuf:"varint,4,opt,name=feeder_id,json=feederId,proto3" json:"feeder_id,omitempty"`
	Status       int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *RoundInfo) Reset()         { *m = RoundInfo{} }
func (m *RoundInfo) String() string { return proto.CompactTextString(m) }
func (*RoundInfo) ProtoMessage()    {}
func (*RoundInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bbd3567e4c61954, []int{0}
}
func (m *RoundInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoundInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoundInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoundInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundInfo.Merge(m, src)
}
func (m *RoundInfo) XXX_Size() int {
	return m.Size()
}
func (m *RoundInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoundInfo proto.InternalMessageInfo

func (m *RoundInfo) GetTokenId() int32 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *RoundInfo) GetBasedBlock() uint64 {
	if m != nil {
		return m.BasedBlock
	}
	return 0
}

func (m *RoundInfo) GetNexitRoundId() uint64 {
	if m != nil {
		return m.NexitRoundId
	}
	return 0
}

func (m *RoundInfo) GetFeederId() int32 {
	if m != nil {
		return m.FeederId
	}
	return 0
}

func (m *RoundInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*RoundInfo)(nil), "exocore.oracle.RoundInfo")
}

func init() { proto.RegisterFile("exocore/oracle/round_info.proto", fileDescriptor_1bbd3567e4c61954) }

var fileDescriptor_1bbd3567e4c61954 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0x87, 0x63, 0x68, 0x4b, 0x6b, 0xfe, 0x0c, 0x1e, 0x50, 0x10, 0x92, 0x5b, 0xc1, 0xd2, 0x29,
	0x41, 0xe2, 0x06, 0x95, 0x18, 0xc2, 0xc0, 0x90, 0x91, 0x25, 0x4a, 0xe2, 0x17, 0x88, 0x52, 0xfc,
	0x2a, 0xc7, 0x11, 0xe1, 0x16, 0x5c, 0x81, 0xdb, 0x30, 0x76, 0x64, 0x44, 0xc9, 0x45, 0xaa, 0x3c,
	0x37, 0xe3, 0xfb, 0xfc, 0x49, 0x9f, 0xfc, 0xe3, 0x4b, 0x68, 0x31, 0x47, 0x03, 0x21, 0x9a, 0x34,
	0xdf, 0x42, 0x68, 0xb0, 0xd1, 0x2a, 0x29, 0x75, 0x81, 0xc1, 0xce, 0xa0, 0x45, 0x71, 0x75, 0x14,
	0x02, 0x27, 0xdc, 0xfd, 0x30, 0xbe, 0x88, 0x07, 0x29, 0xd2, 0x05, 0x8a, 0x1b, 0x3e, 0xb7, 0x58,
	0x81, 0x4e, 0x4a, 0xe5, 0xb3, 0x15, 0x5b, 0x4f, 0xe3, 0x33, 0xba, 0x23, 0x25, 0x96, 0xfc, 0x3c,
	0x4b, 0x6b, 0x50, 0x49, 0xb6, 0xc5, 0xbc, 0xf2, 0x4f, 0x56, 0x6c, 0x3d, 0x89, 0x39, 0xa1, 0xcd,
	0x40, 0xc4, 0x3d, 0xbf, 0xd4, 0xd0, 0x96, 0x36, 0xa1, 0x66, 0xa4, 0xfc, 0x53, 0x52, 0x2e, 0x08,
	0xba, 0x84, 0x12, 0xb7, 0x7c, 0x51, 0x00, 0x28, 0x30, 0x43, 0x61, 0x42, 0x85, 0xb9, 0x03, 0x91,
	0x12, 0xd7, 0x7c, 0x56, 0xdb, 0xd4, 0x36, 0xb5, 0x3f, 0xa5, 0x97, 0xe3, 0xb5, 0x79, 0xfe, 0xed,
	0x24, 0xdb, 0x77, 0x92, 0xfd, 0x77, 0x92, 0x7d, 0xf7, 0xd2, 0xdb, 0xf7, 0xd2, 0xfb, 0xeb, 0xa5,
	0xf7, 0xfa, 0xf0, 0x56, 0xda, 0xf7, 0x26, 0x0b, 0x72, 0xfc, 0x08, 0x9f, 0xdc, 0xc7, 0x5e, 0xc0,
	0x7e, 0xa2, 0xa9, 0xc2, 0x71, 0x88, 0x76, 0x9c, 0xc2, 0x7e, 0xed, 0xa0, 0xce, 0x66, 0x34, 0xc3,
	0xe3, 0x21, 0x00, 0x00, 0xff, 0xff, 0x66, 0xb4, 0x0d, 0xef, 0x29, 0x01, 0x00, 0x00,
}

func (m *RoundInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoundInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoundInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintRoundInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.FeederId != 0 {
		i = encodeVarintRoundInfo(dAtA, i, uint64(m.FeederId))
		i--
		dAtA[i] = 0x20
	}
	if m.NexitRoundId != 0 {
		i = encodeVarintRoundInfo(dAtA, i, uint64(m.NexitRoundId))
		i--
		dAtA[i] = 0x18
	}
	if m.BasedBlock != 0 {
		i = encodeVarintRoundInfo(dAtA, i, uint64(m.BasedBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.TokenId != 0 {
		i = encodeVarintRoundInfo(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoundInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoundInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoundInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenId != 0 {
		n += 1 + sovRoundInfo(uint64(m.TokenId))
	}
	if m.BasedBlock != 0 {
		n += 1 + sovRoundInfo(uint64(m.BasedBlock))
	}
	if m.NexitRoundId != 0 {
		n += 1 + sovRoundInfo(uint64(m.NexitRoundId))
	}
	if m.FeederId != 0 {
		n += 1 + sovRoundInfo(uint64(m.FeederId))
	}
	if m.Status != 0 {
		n += 1 + sovRoundInfo(uint64(m.Status))
	}
	return n
}

func sovRoundInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoundInfo(x uint64) (n int) {
	return sovRoundInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoundInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoundInfo
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
			return fmt.Errorf("proto: RoundInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoundInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoundInfo
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
				return fmt.Errorf("proto: wrong wireType = %d for field BasedBlock", wireType)
			}
			m.BasedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoundInfo
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NexitRoundId", wireType)
			}
			m.NexitRoundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoundInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NexitRoundId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederId", wireType)
			}
			m.FeederId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoundInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeederId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoundInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoundInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoundInfo
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
func skipRoundInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoundInfo
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
					return 0, ErrIntOverflowRoundInfo
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
					return 0, ErrIntOverflowRoundInfo
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
				return 0, ErrInvalidLengthRoundInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoundInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoundInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoundInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoundInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoundInfo = fmt.Errorf("proto: unexpected end of group")
)