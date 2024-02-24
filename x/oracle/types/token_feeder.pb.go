// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/token_feeder.proto

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

// n out of m required source
type NOMSource struct {
	//required source set, refer to params.sourceList, 1st set to 0 means all valid sources
	SourceIds []int32 `protobuf:"varint,1,rep,packed,name=source_ids,json=sourceIds,proto3" json:"source_ids,omitempty"`
	//minimum number from the required sources to be fullfiled
	Minimum int32 `protobuf:"varint,2,opt,name=minimum,proto3" json:"minimum,omitempty"`
}

func (m *NOMSource) Reset()         { *m = NOMSource{} }
func (m *NOMSource) String() string { return proto.CompactTextString(m) }
func (*NOMSource) ProtoMessage()    {}
func (*NOMSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cc86055064704d5, []int{0}
}
func (m *NOMSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NOMSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NOMSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NOMSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NOMSource.Merge(m, src)
}
func (m *NOMSource) XXX_Size() int {
	return m.Size()
}
func (m *NOMSource) XXX_DiscardUnknown() {
	xxx_messageInfo_NOMSource.DiscardUnknown(m)
}

var xxx_messageInfo_NOMSource proto.InternalMessageInfo

func (m *NOMSource) GetSourceIds() []int32 {
	if m != nil {
		return m.SourceIds
	}
	return nil
}

func (m *NOMSource) GetMinimum() int32 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

// specify data from which source is needed
// rule_1: specified sources
// rule_2: n out of total sources are required
type RuleWithSource struct {
	//refer to params.sourceList.ID, when length>0, ignore the other field, when 1st set to 0, means all valid sources, length==0->check next field:minimum
	SourceIds []int32 `protobuf:"varint,1,rep,packed,name=source_ids,json=sourceIds,proto3" json:"source_ids,omitempty"`
	//n out of total sources are required
	Nom *NOMSource `protobuf:"bytes,2,opt,name=nom,proto3" json:"nom,omitempty"`
}

func (m *RuleWithSource) Reset()         { *m = RuleWithSource{} }
func (m *RuleWithSource) String() string { return proto.CompactTextString(m) }
func (*RuleWithSource) ProtoMessage()    {}
func (*RuleWithSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cc86055064704d5, []int{1}
}
func (m *RuleWithSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuleWithSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuleWithSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuleWithSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleWithSource.Merge(m, src)
}
func (m *RuleWithSource) XXX_Size() int {
	return m.Size()
}
func (m *RuleWithSource) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleWithSource.DiscardUnknown(m)
}

var xxx_messageInfo_RuleWithSource proto.InternalMessageInfo

func (m *RuleWithSource) GetSourceIds() []int32 {
	if m != nil {
		return m.SourceIds
	}
	return nil
}

func (m *RuleWithSource) GetNom() *NOMSource {
	if m != nil {
		return m.Nom
	}
	return nil
}

// Tokenfeeder represents a price oracle for one token
type TokenFeeder struct {
	//refer to params.tokenList, from 1
	TokenId int32 `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	//refer to params.ruleList, 0 means no restriction, accept any source including customer defined
	RuleId int32 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	//include, from 1, when some token's feeder had been stop and then restart, the token_id will be continuous from previous one
	StartRoundId int64 `protobuf:"varint,3,opt,name=start_round_id,json=startRoundId,proto3" json:"start_round_id,omitempty"`
	//include, first block which start_round_id can be settled is at least start_base_block+1
	StartBaseBlock int64 `protobuf:"varint,4,opt,name=start_base_block,json=startBaseBlock,proto3" json:"start_base_block,omitempty"`
	//set as count of blocks, for how many blocks interval the price will be update once
	Interval int64 `protobuf:"varint,5,opt,name=interval,proto3" json:"interval,omitempty"`
	//tokenfeeder is initialized with forever live, update the End parameters by voting, and will off service by the end
	//this is set by updateParams, and the EndRoundID will be update by related. excluded, will not work if current height >=EndBlock
	EndBlock int64 `protobuf:"varint,6,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (m *TokenFeeder) Reset()         { *m = TokenFeeder{} }
func (m *TokenFeeder) String() string { return proto.CompactTextString(m) }
func (*TokenFeeder) ProtoMessage()    {}
func (*TokenFeeder) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cc86055064704d5, []int{2}
}
func (m *TokenFeeder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenFeeder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenFeeder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenFeeder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenFeeder.Merge(m, src)
}
func (m *TokenFeeder) XXX_Size() int {
	return m.Size()
}
func (m *TokenFeeder) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenFeeder.DiscardUnknown(m)
}

var xxx_messageInfo_TokenFeeder proto.InternalMessageInfo

func (m *TokenFeeder) GetTokenId() int32 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *TokenFeeder) GetRuleId() int32 {
	if m != nil {
		return m.RuleId
	}
	return 0
}

func (m *TokenFeeder) GetStartRoundId() int64 {
	if m != nil {
		return m.StartRoundId
	}
	return 0
}

func (m *TokenFeeder) GetStartBaseBlock() int64 {
	if m != nil {
		return m.StartBaseBlock
	}
	return 0
}

func (m *TokenFeeder) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *TokenFeeder) GetEndBlock() int64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*NOMSource)(nil), "exocore.oracle.NOMSource")
	proto.RegisterType((*RuleWithSource)(nil), "exocore.oracle.RuleWithSource")
	proto.RegisterType((*TokenFeeder)(nil), "exocore.oracle.TokenFeeder")
}

func init() { proto.RegisterFile("exocore/oracle/token_feeder.proto", fileDescriptor_1cc86055064704d5) }

var fileDescriptor_1cc86055064704d5 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0x8e, 0xe6, 0xe5, 0x4f, 0x19, 0x61, 0x88, 0xc1, 0x9c, 0x8c, 0x99, 0x2c, 0xec, 0xc2, 0x30,
	0xb0, 0xc7, 0xf6, 0x06, 0x61, 0x1b, 0x78, 0xb0, 0x0c, 0xbc, 0x41, 0xa1, 0x14, 0x8c, 0x6d, 0x9d,
	0x24, 0x22, 0xb6, 0x14, 0x64, 0xb9, 0x4d, 0xdf, 0xa2, 0x8f, 0x55, 0x7a, 0x95, 0xcb, 0x5e, 0x96,
	0xe4, 0x45, 0x8a, 0x24, 0x27, 0xd0, 0x5e, 0xf5, 0xee, 0x7c, 0x3f, 0xe7, 0xe8, 0xfc, 0x08, 0x7f,
	0x82, 0xad, 0xc8, 0x85, 0x84, 0x50, 0xc8, 0x34, 0x2f, 0x20, 0x54, 0x62, 0x0d, 0x3c, 0x59, 0x00,
	0x50, 0x90, 0xc1, 0x46, 0x0a, 0x25, 0xc8, 0xb0, 0xb1, 0x04, 0xd6, 0x32, 0x7e, 0xb7, 0x14, 0x4b,
	0x61, 0xa4, 0x50, 0x47, 0xd6, 0x35, 0x1e, 0x3d, 0x2b, 0xc4, 0xf8, 0xa2, 0x91, 0xa6, 0x3f, 0x70,
	0x7f, 0xfe, 0xf7, 0xcf, 0x3f, 0x51, 0xcb, 0x1c, 0xc8, 0x47, 0x8c, 0x2b, 0x13, 0x25, 0x8c, 0x56,
	0x2e, 0x9a, 0x38, 0x7e, 0x3b, 0xee, 0x5b, 0x26, 0xa2, 0x15, 0x71, 0x71, 0xb7, 0x64, 0x9c, 0x95,
	0x75, 0xe9, 0xbe, 0x9a, 0x20, 0xbf, 0x1d, 0x1f, 0xe1, 0xf4, 0x02, 0x0f, 0xe3, 0xba, 0x80, 0x33,
	0xa6, 0x56, 0x2f, 0x2b, 0xf5, 0x05, 0x3b, 0x5c, 0xd8, 0x32, 0x83, 0x6f, 0xa3, 0xe0, 0xe9, 0x14,
	0xc1, 0xa9, 0xa3, 0x58, 0xbb, 0xa6, 0x77, 0x08, 0x0f, 0xfe, 0xeb, 0xd9, 0x7f, 0x99, 0xd1, 0xc9,
	0x08, 0xf7, 0xec, 0x2a, 0x18, 0x75, 0x91, 0x6d, 0xc4, 0xe0, 0x88, 0x92, 0xf7, 0xb8, 0x2b, 0xeb,
	0x42, 0x3f, 0xda, 0xb4, 0xd8, 0xd1, 0x30, 0xa2, 0xe4, 0x33, 0x1e, 0x56, 0x2a, 0x95, 0x2a, 0x91,
	0xa2, 0xe6, 0x54, 0xeb, 0xce, 0x04, 0xf9, 0x4e, 0xfc, 0xc6, 0xb0, 0xb1, 0x26, 0x23, 0x4a, 0x7c,
	0xfc, 0xd6, 0xba, 0xb2, 0xb4, 0x82, 0x24, 0x2b, 0x44, 0xbe, 0x76, 0x5f, 0x1b, 0x9f, 0xcd, 0x9e,
	0xa5, 0x15, 0xcc, 0x34, 0x4b, 0xc6, 0xb8, 0xc7, 0xb8, 0x02, 0x79, 0x99, 0x16, 0x6e, 0xdb, 0x38,
	0x4e, 0x98, 0x7c, 0xc0, 0x7d, 0xe0, 0xb4, 0x49, 0xef, 0x58, 0x11, 0x38, 0x35, 0x89, 0xb3, 0xdf,
	0xb7, 0x7b, 0x0f, 0xed, 0xf6, 0x1e, 0x7a, 0xd8, 0x7b, 0xe8, 0xe6, 0xe0, 0xb5, 0x76, 0x07, 0xaf,
	0x75, 0x7f, 0xf0, 0x5a, 0xe7, 0x5f, 0x97, 0x4c, 0xad, 0xea, 0x2c, 0xc8, 0x45, 0x19, 0xfe, 0xb4,
	0x0b, 0x99, 0x83, 0xba, 0x12, 0x72, 0x1d, 0x1e, 0xef, 0xb7, 0x3d, 0x7d, 0x85, 0xeb, 0x0d, 0x54,
	0x59, 0xc7, 0xdc, 0xf0, 0xfb, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xf1, 0x12, 0xb2, 0x29,
	0x02, 0x00, 0x00,
}

func (m *NOMSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NOMSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NOMSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Minimum != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.Minimum))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SourceIds) > 0 {
		dAtA2 := make([]byte, len(m.SourceIds)*10)
		var j1 int
		for _, num1 := range m.SourceIds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTokenFeeder(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RuleWithSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuleWithSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuleWithSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nom != nil {
		{
			size, err := m.Nom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTokenFeeder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourceIds) > 0 {
		dAtA5 := make([]byte, len(m.SourceIds)*10)
		var j4 int
		for _, num1 := range m.SourceIds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		i -= j4
		copy(dAtA[i:], dAtA5[:j4])
		i = encodeVarintTokenFeeder(dAtA, i, uint64(j4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenFeeder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenFeeder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenFeeder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlock != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x30
	}
	if m.Interval != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.Interval))
		i--
		dAtA[i] = 0x28
	}
	if m.StartBaseBlock != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.StartBaseBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.StartRoundId != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.StartRoundId))
		i--
		dAtA[i] = 0x18
	}
	if m.RuleId != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.RuleId))
		i--
		dAtA[i] = 0x10
	}
	if m.TokenId != 0 {
		i = encodeVarintTokenFeeder(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenFeeder(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenFeeder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NOMSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SourceIds) > 0 {
		l = 0
		for _, e := range m.SourceIds {
			l += sovTokenFeeder(uint64(e))
		}
		n += 1 + sovTokenFeeder(uint64(l)) + l
	}
	if m.Minimum != 0 {
		n += 1 + sovTokenFeeder(uint64(m.Minimum))
	}
	return n
}

func (m *RuleWithSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SourceIds) > 0 {
		l = 0
		for _, e := range m.SourceIds {
			l += sovTokenFeeder(uint64(e))
		}
		n += 1 + sovTokenFeeder(uint64(l)) + l
	}
	if m.Nom != nil {
		l = m.Nom.Size()
		n += 1 + l + sovTokenFeeder(uint64(l))
	}
	return n
}

func (m *TokenFeeder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenId != 0 {
		n += 1 + sovTokenFeeder(uint64(m.TokenId))
	}
	if m.RuleId != 0 {
		n += 1 + sovTokenFeeder(uint64(m.RuleId))
	}
	if m.StartRoundId != 0 {
		n += 1 + sovTokenFeeder(uint64(m.StartRoundId))
	}
	if m.StartBaseBlock != 0 {
		n += 1 + sovTokenFeeder(uint64(m.StartBaseBlock))
	}
	if m.Interval != 0 {
		n += 1 + sovTokenFeeder(uint64(m.Interval))
	}
	if m.EndBlock != 0 {
		n += 1 + sovTokenFeeder(uint64(m.EndBlock))
	}
	return n
}

func sovTokenFeeder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenFeeder(x uint64) (n int) {
	return sovTokenFeeder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NOMSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenFeeder
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
			return fmt.Errorf("proto: NOMSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NOMSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTokenFeeder
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.SourceIds = append(m.SourceIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTokenFeeder
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTokenFeeder
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTokenFeeder
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.SourceIds) == 0 {
					m.SourceIds = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTokenFeeder
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.SourceIds = append(m.SourceIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceIds", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minimum", wireType)
			}
			m.Minimum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minimum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokenFeeder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenFeeder
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
func (m *RuleWithSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenFeeder
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
			return fmt.Errorf("proto: RuleWithSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuleWithSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTokenFeeder
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.SourceIds = append(m.SourceIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTokenFeeder
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTokenFeeder
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTokenFeeder
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.SourceIds) == 0 {
					m.SourceIds = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTokenFeeder
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.SourceIds = append(m.SourceIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceIds", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
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
				return ErrInvalidLengthTokenFeeder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenFeeder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nom == nil {
				m.Nom = &NOMSource{}
			}
			if err := m.Nom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenFeeder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenFeeder
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
func (m *TokenFeeder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenFeeder
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
			return fmt.Errorf("proto: TokenFeeder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenFeeder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
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
				return fmt.Errorf("proto: wrong wireType = %d for field RuleId", wireType)
			}
			m.RuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RuleId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartRoundId", wireType)
			}
			m.StartRoundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartRoundId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBaseBlock", wireType)
			}
			m.StartBaseBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBaseBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			m.Interval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Interval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokenFeeder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenFeeder
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
func skipTokenFeeder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenFeeder
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
					return 0, ErrIntOverflowTokenFeeder
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
					return 0, ErrIntOverflowTokenFeeder
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
				return 0, ErrInvalidLengthTokenFeeder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenFeeder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenFeeder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenFeeder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenFeeder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenFeeder = fmt.Errorf("proto: unexpected end of group")
)
