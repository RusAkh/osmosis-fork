// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/twap/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the twap module
type Params struct {
	PruneEpochIdentifier    string        `protobuf:"bytes,1,opt,name=prune_epoch_identifier,json=pruneEpochIdentifier,proto3" json:"prune_epoch_identifier,omitempty"`
	RecordHistoryKeepPeriod time.Duration `protobuf:"bytes,2,opt,name=record_history_keep_period,json=recordHistoryKeepPeriod,proto3,stdduration" json:"record_history_keep_period" yaml:"record_history_keep_period"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f4bdf49b69bd63c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPruneEpochIdentifier() string {
	if m != nil {
		return m.PruneEpochIdentifier
	}
	return ""
}

func (m *Params) GetRecordHistoryKeepPeriod() time.Duration {
	if m != nil {
		return m.RecordHistoryKeepPeriod
	}
	return 0
}

// GenesisState defines the twap module's genesis state.
type GenesisState struct {
	// twaps is the collection of all twap records.
	Twaps []TwapRecord `protobuf:"bytes,1,rep,name=twaps,proto3" json:"twaps"`
	// params is the container of twap parameters.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f4bdf49b69bd63c, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetTwaps() []TwapRecord {
	if m != nil {
		return m.Twaps
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.twap.v1beta1.Params")
	proto.RegisterType((*GenesisState)(nil), "osmosis.twap.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/twap/v1beta1/genesis.proto", fileDescriptor_3f4bdf49b69bd63c)
}

var fileDescriptor_3f4bdf49b69bd63c = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0xce, 0xd3, 0x30,
	0x14, 0xc5, 0x63, 0xfe, 0x54, 0x22, 0x1f, 0x53, 0x54, 0x41, 0xbf, 0x0a, 0xa5, 0x21, 0x03, 0xea,
	0x52, 0x9b, 0xb6, 0x4c, 0x15, 0x53, 0x05, 0xe2, 0xdf, 0x52, 0x05, 0x26, 0x96, 0xc8, 0x49, 0x6e,
	0x53, 0x8b, 0x26, 0xb6, 0x6c, 0xa7, 0x25, 0x0f, 0x80, 0xc4, 0xc8, 0xc8, 0x13, 0xa1, 0x8e, 0x1d,
	0x99, 0x0a, 0x6a, 0xdf, 0x80, 0x27, 0x40, 0x89, 0x5d, 0x06, 0xd4, 0x6f, 0xcb, 0xd5, 0xef, 0x9c,
	0xa3, 0x93, 0x7b, 0xed, 0x86, 0x5c, 0x15, 0x5c, 0x31, 0x45, 0xf4, 0x96, 0x0a, 0xb2, 0x19, 0x27,
	0xa0, 0xe9, 0x98, 0xe4, 0x50, 0x82, 0x62, 0x0a, 0x0b, 0xc9, 0x35, 0xf7, 0xba, 0x56, 0x83, 0x1b,
	0x0d, 0xb6, 0x9a, 0x7e, 0x37, 0xe7, 0x39, 0x6f, 0x05, 0xa4, 0xf9, 0x32, 0xda, 0xfe, 0x93, 0x8b,
	0x79, 0xcd, 0x10, 0x4b, 0x48, 0xb9, 0xcc, 0xac, 0xee, 0x3a, 0xe7, 0x3c, 0x5f, 0x03, 0x69, 0xa7,
	0xa4, 0x5a, 0x12, 0x5a, 0xd6, 0x67, 0x94, 0xb6, 0x19, 0xb1, 0xc9, 0x36, 0x83, 0x45, 0xfe, 0xff,
	0xae, 0xac, 0x92, 0x54, 0x33, 0x5e, 0x1a, 0x1e, 0xfe, 0x40, 0x6e, 0x67, 0x41, 0x25, 0x2d, 0x94,
	0xf7, 0xcc, 0x7d, 0x20, 0x64, 0x55, 0x42, 0x0c, 0x82, 0xa7, 0xab, 0x98, 0x65, 0x50, 0x6a, 0xb6,
	0x64, 0x20, 0x7b, 0x28, 0x40, 0xc3, 0x7b, 0x51, 0xb7, 0xa5, 0x2f, 0x1b, 0xf8, 0xe6, 0x1f, 0xf3,
	0xbe, 0x20, 0xb7, 0x6f, 0x7a, 0xc6, 0x2b, 0xa6, 0x34, 0x97, 0x75, 0xfc, 0x09, 0x40, 0xc4, 0x02,
	0x24, 0xe3, 0x59, 0xef, 0x56, 0x80, 0x86, 0x57, 0x93, 0x6b, 0x6c, 0x6a, 0xe0, 0x73, 0x0d, 0xfc,
	0xc2, 0xd6, 0x98, 0x8f, 0x76, 0x87, 0x81, 0xf3, 0xe7, 0x30, 0x78, 0x5c, 0xd3, 0x62, 0x3d, 0x0b,
	0x6f, 0x8e, 0x0a, 0xbf, 0xff, 0x1a, 0xa0, 0xe8, 0xa1, 0x11, 0xbc, 0x36, 0xfc, 0x1d, 0x80, 0x58,
	0x18, 0xfa, 0x15, 0xb9, 0xf7, 0x5f, 0x99, 0x23, 0xbc, 0xd7, 0x54, 0x83, 0xf7, 0xdc, 0xbd, 0xdb,
	0x2c, 0x51, 0xf5, 0x50, 0x70, 0x7b, 0x78, 0x35, 0x09, 0xf0, 0xa5, 0x9b, 0xe0, 0x0f, 0x5b, 0x2a,
	0xa2, 0x36, 0x72, 0x7e, 0xa7, 0x69, 0x12, 0x19, 0x93, 0x37, 0x73, 0x3b, 0xa2, 0x5d, 0x8b, 0xfd,
	0x83, 0x47, 0x97, 0xed, 0x66, 0x75, 0xd6, 0x6a, 0x1d, 0xf3, 0xb7, 0xbb, 0xa3, 0x8f, 0xf6, 0x47,
	0x1f, 0xfd, 0x3e, 0xfa, 0xe8, 0xdb, 0xc9, 0x77, 0xf6, 0x27, 0xdf, 0xf9, 0x79, 0xf2, 0x9d, 0x8f,
	0x4f, 0x73, 0xa6, 0x57, 0x55, 0x82, 0x53, 0x5e, 0x10, 0x9b, 0x37, 0x5a, 0xd3, 0x44, 0x9d, 0x07,
	0xb2, 0x19, 0x4f, 0xc9, 0x67, 0xf3, 0x12, 0x74, 0x2d, 0x40, 0x25, 0x9d, 0x76, 0x63, 0xd3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xba, 0xd9, 0x22, 0x76, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RecordHistoryKeepPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.RecordHistoryKeepPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.PruneEpochIdentifier) > 0 {
		i -= len(m.PruneEpochIdentifier)
		copy(dAtA[i:], m.PruneEpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PruneEpochIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Twaps) > 0 {
		for iNdEx := len(m.Twaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Twaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PruneEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RecordHistoryKeepPeriod)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Twaps) > 0 {
		for _, e := range m.Twaps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PruneEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PruneEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordHistoryKeepPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RecordHistoryKeepPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Twaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Twaps = append(m.Twaps, TwapRecord{})
			if err := m.Twaps[len(m.Twaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
