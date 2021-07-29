// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/genesis.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the launch module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	ChainList          []*Chain          `protobuf:"bytes,1,rep,name=chainList,proto3" json:"chainList,omitempty"`
	ChainNameCountList []*ChainNameCount `protobuf:"bytes,2,rep,name=chainNameCountList,proto3" json:"chainNameCountList,omitempty"`
	GenesisAccountList []*GenesisAccount `protobuf:"bytes,3,rep,name=genesisAccountList,proto3" json:"genesisAccountList,omitempty"`
	VestedAccountList  []*VestedAccount  `protobuf:"bytes,4,rep,name=vestedAccountList,proto3" json:"vestedAccountList,omitempty"`
	RequestList        []*Request        `protobuf:"bytes,10,rep,name=requestList,proto3" json:"requestList,omitempty"`
	RequestCountList   []*RequestCount   `protobuf:"bytes,11,rep,name=requestCountList,proto3" json:"requestCountList,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cd66d27edc51cd, []int{0}
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

func (m *GenesisState) GetChainList() []*Chain {
	if m != nil {
		return m.ChainList
	}
	return nil
}

func (m *GenesisState) GetChainNameCountList() []*ChainNameCount {
	if m != nil {
		return m.ChainNameCountList
	}
	return nil
}

func (m *GenesisState) GetGenesisAccountList() []*GenesisAccount {
	if m != nil {
		return m.GenesisAccountList
	}
	return nil
}

func (m *GenesisState) GetVestedAccountList() []*VestedAccount {
	if m != nil {
		return m.VestedAccountList
	}
	return nil
}

func (m *GenesisState) GetRequestList() []*Request {
	if m != nil {
		return m.RequestList
	}
	return nil
}

func (m *GenesisState) GetRequestCountList() []*RequestCount {
	if m != nil {
		return m.RequestCountList
	}
	return nil
}

type RequestCount struct {
	ChainID string `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Count   uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *RequestCount) Reset()         { *m = RequestCount{} }
func (m *RequestCount) String() string { return proto.CompactTextString(m) }
func (*RequestCount) ProtoMessage()    {}
func (*RequestCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cd66d27edc51cd, []int{1}
}
func (m *RequestCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCount.Merge(m, src)
}
func (m *RequestCount) XXX_Size() int {
	return m.Size()
}
func (m *RequestCount) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCount.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCount proto.InternalMessageInfo

func (m *RequestCount) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *RequestCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.launch.GenesisState")
	proto.RegisterType((*RequestCount)(nil), "tendermint.spn.launch.RequestCount")
}

func init() { proto.RegisterFile("launch/genesis.proto", fileDescriptor_02cd66d27edc51cd) }

var fileDescriptor_02cd66d27edc51cd = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4b, 0x3a, 0x41,
	0x14, 0xc7, 0x1d, 0xf5, 0xf7, 0x0b, 0x47, 0x0f, 0x35, 0x18, 0x2c, 0x26, 0x83, 0x58, 0xc1, 0x9e,
	0x66, 0xa1, 0x6e, 0x1d, 0xa2, 0x34, 0x90, 0x20, 0x0a, 0x26, 0xea, 0xd0, 0x25, 0xd6, 0x75, 0xd0,
	0x85, 0x9c, 0xdd, 0x9c, 0xd9, 0xa8, 0xff, 0xa2, 0x3f, 0xab, 0xa3, 0xc7, 0x6e, 0x85, 0xfe, 0x23,
	0xe1, 0x9b, 0x31, 0xd7, 0xd4, 0x3d, 0xbe, 0x79, 0x9f, 0xef, 0x67, 0xe7, 0xbd, 0x1d, 0x5c, 0x7d,
	0xf2, 0x13, 0x19, 0x0c, 0xbc, 0xbe, 0x90, 0x42, 0x85, 0x8a, 0xc5, 0xa3, 0x48, 0x47, 0x64, 0x57,
	0x0b, 0xd9, 0x13, 0xa3, 0x61, 0x28, 0x35, 0x53, 0xb1, 0x64, 0x06, 0xaa, 0xcd, 0xe1, 0x91, 0x78,
	0x4e, 0x84, 0xd2, 0x06, 0xae, 0xed, 0xd9, 0xd3, 0x17, 0xa1, 0xb4, 0xe8, 0x3d, 0xfa, 0x41, 0x10,
	0x25, 0x72, 0xde, 0xac, 0x2f, 0xfb, 0xff, 0x74, 0x89, 0xed, 0x06, 0x03, 0x3f, 0x94, 0xe6, 0xac,
	0xf9, 0x55, 0xc0, 0x95, 0x8e, 0xa1, 0x6f, 0xb5, 0xaf, 0x05, 0x39, 0xc1, 0x25, 0xe8, 0x5f, 0x85,
	0x4a, 0x3b, 0xa8, 0x51, 0x70, 0xcb, 0x47, 0x75, 0xb6, 0xf6, 0x82, 0xac, 0x3d, 0xe3, 0xf8, 0x02,
	0x27, 0x77, 0x98, 0x40, 0x71, 0xed, 0x0f, 0x45, 0x7b, 0xf6, 0x61, 0x90, 0xe4, 0x41, 0x72, 0x98,
	0x25, 0xf9, 0x0d, 0xf0, 0x35, 0x82, 0x99, 0xd6, 0x0e, 0x74, 0x6e, 0xe6, 0x01, 0x6d, 0x21, 0x53,
	0xdb, 0x59, 0x0a, 0xf0, 0x35, 0x02, 0xc2, 0xf1, 0x8e, 0x59, 0x62, 0xda, 0x5a, 0x04, 0xeb, 0xc1,
	0x06, 0xeb, 0x7d, 0x9a, 0xe7, 0xab, 0x71, 0x72, 0x86, 0xcb, 0xf6, 0x77, 0x81, 0x0d, 0x83, 0x8d,
	0x6e, 0xb0, 0x71, 0x43, 0xf2, 0x74, 0x84, 0xdc, 0xe0, 0x6d, 0x5b, 0x2e, 0x36, 0x58, 0x06, 0xcd,
	0x7e, 0xb6, 0xc6, 0xec, 0x6f, 0x25, 0xdc, 0x3c, 0xc5, 0x95, 0x34, 0x41, 0x1c, 0xbc, 0x05, 0x3b,
	0xbe, 0xbc, 0x70, 0x50, 0x03, 0xb9, 0x25, 0x3e, 0x2f, 0x49, 0x15, 0xff, 0x83, 0x49, 0x9c, 0x7c,
	0x03, 0xb9, 0x45, 0x6e, 0x8a, 0x56, 0xeb, 0x63, 0x42, 0xd1, 0x78, 0x42, 0xd1, 0xf7, 0x84, 0xa2,
	0xf7, 0x29, 0xcd, 0x8d, 0xa7, 0x34, 0xf7, 0x39, 0xa5, 0xb9, 0x07, 0xb7, 0x1f, 0xea, 0x41, 0xd2,
	0x65, 0x41, 0x34, 0xf4, 0x16, 0x57, 0xf3, 0x54, 0x2c, 0xbd, 0x57, 0xcf, 0xbe, 0x35, 0xfd, 0x16,
	0x0b, 0xd5, 0xfd, 0x0f, 0x8f, 0xed, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x86, 0x50, 0x01, 0x50,
	0x00, 0x03, 0x00, 0x00,
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
	if len(m.RequestCountList) > 0 {
		for iNdEx := len(m.RequestCountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestCountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.RequestList) > 0 {
		for iNdEx := len(m.RequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.VestedAccountList) > 0 {
		for iNdEx := len(m.VestedAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestedAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.GenesisAccountList) > 0 {
		for iNdEx := len(m.GenesisAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ChainNameCountList) > 0 {
		for iNdEx := len(m.ChainNameCountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainNameCountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ChainList) > 0 {
		for iNdEx := len(m.ChainList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *RequestCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainList) > 0 {
		for _, e := range m.ChainList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChainNameCountList) > 0 {
		for _, e := range m.ChainNameCountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GenesisAccountList) > 0 {
		for _, e := range m.GenesisAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VestedAccountList) > 0 {
		for _, e := range m.VestedAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RequestList) > 0 {
		for _, e := range m.RequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RequestCountList) > 0 {
		for _, e := range m.RequestCountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *RequestCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovGenesis(uint64(m.Count))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field ChainList", wireType)
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
			m.ChainList = append(m.ChainList, &Chain{})
			if err := m.ChainList[len(m.ChainList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainNameCountList", wireType)
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
			m.ChainNameCountList = append(m.ChainNameCountList, &ChainNameCount{})
			if err := m.ChainNameCountList[len(m.ChainNameCountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisAccountList", wireType)
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
			m.GenesisAccountList = append(m.GenesisAccountList, &GenesisAccount{})
			if err := m.GenesisAccountList[len(m.GenesisAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestedAccountList", wireType)
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
			m.VestedAccountList = append(m.VestedAccountList, &VestedAccount{})
			if err := m.VestedAccountList[len(m.VestedAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestList", wireType)
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
			m.RequestList = append(m.RequestList, &Request{})
			if err := m.RequestList[len(m.RequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestCountList", wireType)
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
			m.RequestCountList = append(m.RequestCountList, &RequestCount{})
			if err := m.RequestCountList[len(m.RequestCountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *RequestCount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RequestCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
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
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
