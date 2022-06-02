// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stratos/register/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// GenesisState defines the register module's genesis state.
type GenesisState struct {
	Params              *Params                                `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty" yaml:"params"`
	ResourceNodes       ResourceNodes                          `protobuf:"bytes,2,rep,name=resource_nodes,json=resourceNodes,proto3,castrepeated=ResourceNodes" json:"resource_nodes" yaml:"resource_nodes"`
	MetaNodes           MetaNodes                              `protobuf:"bytes,3,rep,name=meta_nodes,json=metaNodes,proto3,castrepeated=MetaNodes" json:"meta_nodes" yaml:"meta_nodes"`
	InitialUozPrice     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=initial_uoz_price,json=initialUozPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"initial_uoz_price" yaml:"initial_uoz_price"`
	TotalUnissuedPrepay github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_unissued_prepay,json=totalUnissuedPrepay,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_unissued_prepay" yaml:"total_unissued_prepay"`
	Slashing            []*Slashing                            `protobuf:"bytes,6,rep,name=slashing,proto3" json:"slashing,omitempty" yaml:"slashing_info"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bdab54ebea9e48e, []int{0}
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

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetResourceNodes() ResourceNodes {
	if m != nil {
		return m.ResourceNodes
	}
	return nil
}

func (m *GenesisState) GetMetaNodes() MetaNodes {
	if m != nil {
		return m.MetaNodes
	}
	return nil
}

func (m *GenesisState) GetSlashing() []*Slashing {
	if m != nil {
		return m.Slashing
	}
	return nil
}

type GenesisMetaNode struct {
	NetworkAddress string                                 `protobuf:"bytes,1,opt,name=network_address,json=networkAddress,proto3" json:"network_address,omitempty" yaml:"network_address"`
	Pubkey         *types.Any                             `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty" yaml:"pubkey"`
	Suspend        bool                                   `protobuf:"varint,3,opt,name=suspend,proto3" json:"suspend,omitempty" yaml:"suspend"`
	Status         types1.BondStatus                      `protobuf:"varint,4,opt,name=status,proto3,enum=cosmos.staking.v1beta1.BondStatus" json:"status,omitempty" yaml:"status"`
	Tokens         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=tokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"tokens" yaml:"token"`
	OwnerAddress   string                                 `protobuf:"bytes,6,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty" yaml:"owner_address"`
	Description    *Description                           `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty" yaml:"description",omitempty`
}

func (m *GenesisMetaNode) Reset()         { *m = GenesisMetaNode{} }
func (m *GenesisMetaNode) String() string { return proto.CompactTextString(m) }
func (*GenesisMetaNode) ProtoMessage()    {}
func (*GenesisMetaNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bdab54ebea9e48e, []int{1}
}
func (m *GenesisMetaNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisMetaNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisMetaNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisMetaNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisMetaNode.Merge(m, src)
}
func (m *GenesisMetaNode) XXX_Size() int {
	return m.Size()
}
func (m *GenesisMetaNode) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisMetaNode.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisMetaNode proto.InternalMessageInfo

func (m *GenesisMetaNode) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *GenesisMetaNode) GetPubkey() *types.Any {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *GenesisMetaNode) GetSuspend() bool {
	if m != nil {
		return m.Suspend
	}
	return false
}

func (m *GenesisMetaNode) GetStatus() types1.BondStatus {
	if m != nil {
		return m.Status
	}
	return types1.Unspecified
}

func (m *GenesisMetaNode) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *GenesisMetaNode) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stratos.register.v1.GenesisState")
	proto.RegisterType((*GenesisMetaNode)(nil), "stratos.register.v1.GenesisMetaNode")
}

func init() { proto.RegisterFile("stratos/register/v1/genesis.proto", fileDescriptor_5bdab54ebea9e48e) }

var fileDescriptor_5bdab54ebea9e48e = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xeb, 0x46,
	0x14, 0x8e, 0xf9, 0x09, 0x64, 0x20, 0x41, 0x98, 0x50, 0x19, 0xda, 0xc6, 0x61, 0x54, 0xa1, 0x54,
	0x02, 0x5b, 0xa1, 0x5d, 0x55, 0x6a, 0x25, 0x5c, 0xd4, 0xaa, 0xad, 0x40, 0x91, 0x23, 0x5a, 0xa9,
	0x9b, 0x68, 0x62, 0x0f, 0xc6, 0x4a, 0x3c, 0x63, 0x79, 0xc6, 0x01, 0xb3, 0xec, 0xaa, 0xcb, 0xbe,
	0x46, 0xbb, 0xee, 0x43, 0xa0, 0xae, 0x58, 0x56, 0x5d, 0xb8, 0x57, 0xf0, 0x06, 0x79, 0x82, 0xab,
	0xcc, 0x8c, 0x89, 0xb9, 0x37, 0xba, 0x12, 0x2b, 0xcf, 0x39, 0xe7, 0x3b, 0xdf, 0x37, 0x73, 0x7e,
	0x0c, 0x0e, 0x18, 0x4f, 0x10, 0xa7, 0xcc, 0x4e, 0x70, 0x10, 0x32, 0x8e, 0x13, 0x7b, 0xd2, 0xb5,
	0x03, 0x4c, 0x30, 0x0b, 0x99, 0x15, 0x27, 0x94, 0x53, 0x7d, 0x47, 0x41, 0xac, 0x02, 0x62, 0x4d,
	0xba, 0xfb, 0x7b, 0x01, 0xa5, 0xc1, 0x18, 0xdb, 0x02, 0x32, 0x4c, 0xaf, 0x6c, 0x44, 0x32, 0x89,
	0xdf, 0x6f, 0x06, 0x34, 0xa0, 0xe2, 0x68, 0xcf, 0x4e, 0xca, 0xbb, 0xe7, 0x51, 0x16, 0x51, 0x36,
	0x90, 0x01, 0x69, 0xa8, 0xd0, 0x67, 0xd2, 0xb2, 0x19, 0x47, 0xa3, 0x90, 0x04, 0xf6, 0xa4, 0x3b,
	0xc4, 0x1c, 0x75, 0x0b, 0x5b, 0xa1, 0xe0, 0xa2, 0x9b, 0x3e, 0x5f, 0x49, 0x60, 0xe0, 0xef, 0xab,
	0x60, 0xf3, 0x7b, 0x79, 0xf9, 0x3e, 0x47, 0x1c, 0xeb, 0xdf, 0x81, 0x6a, 0x8c, 0x12, 0x14, 0x31,
	0x43, 0x6b, 0x6b, 0x9d, 0x8d, 0x93, 0x8f, 0xad, 0x05, 0x8f, 0xb1, 0x7a, 0x02, 0xe2, 0x6c, 0x4f,
	0x73, 0xb3, 0x9e, 0xa1, 0x68, 0xfc, 0x15, 0x94, 0x49, 0xd0, 0x55, 0xd9, 0xfa, 0x2d, 0x68, 0x24,
	0x98, 0xd1, 0x34, 0xf1, 0xf0, 0x80, 0x50, 0x1f, 0x33, 0x63, 0xa9, 0xbd, 0xdc, 0xd9, 0x38, 0x39,
	0x58, 0xc8, 0xe7, 0x2a, 0xe8, 0x05, 0xf5, 0xb1, 0x63, 0xdd, 0xe7, 0x66, 0x65, 0x9a, 0x9b, 0xbb,
	0x92, 0xf9, 0x25, 0x0d, 0xfc, 0xeb, 0x7f, 0xb3, 0x5e, 0x86, 0x33, 0xb7, 0x9e, 0x94, 0x4d, 0xdd,
	0x07, 0x20, 0xc2, 0x1c, 0x29, 0xd5, 0x65, 0xa1, 0xfa, 0xe9, 0x42, 0xd5, 0x73, 0xcc, 0x91, 0x50,
	0x3c, 0x54, 0x8a, 0xdb, 0x52, 0x71, 0x9e, 0x3e, 0x53, 0xab, 0x15, 0x30, 0xe6, 0xd6, 0xa2, 0xe2,
	0xa8, 0x4f, 0xc0, 0x76, 0x48, 0x42, 0x1e, 0xa2, 0xf1, 0x20, 0xa5, 0x77, 0x83, 0x38, 0x09, 0x3d,
	0x6c, 0xac, 0xb4, 0xb5, 0x4e, 0xcd, 0xf9, 0x71, 0xc6, 0xf6, 0x5f, 0x6e, 0x1e, 0x06, 0x21, 0xbf,
	0x4e, 0x87, 0x96, 0x47, 0x23, 0xd5, 0x3e, 0xf5, 0x39, 0x66, 0xfe, 0xc8, 0xe6, 0x59, 0x8c, 0x99,
	0x75, 0x86, 0xbd, 0x69, 0x6e, 0x1a, 0x52, 0xf7, 0x3d, 0x42, 0xe8, 0x6e, 0x29, 0xdf, 0x25, 0xbd,
	0xeb, 0xcd, 0x3c, 0xfa, 0x6f, 0x1a, 0xd8, 0xe5, 0x94, 0xcf, 0x50, 0x24, 0x64, 0x2c, 0xc5, 0xfe,
	0x20, 0x4e, 0x70, 0x8c, 0x32, 0x63, 0x55, 0x88, 0x5f, 0xbc, 0x42, 0xfc, 0x07, 0xc2, 0xa7, 0xb9,
	0xf9, 0x89, 0x14, 0x5f, 0x48, 0x0a, 0xdd, 0x1d, 0xe1, 0xbf, 0x54, 0xee, 0x9e, 0xf0, 0xea, 0x7d,
	0xb0, 0xce, 0xc6, 0x88, 0x5d, 0x87, 0x24, 0x30, 0xaa, 0x1f, 0x28, 0x70, 0x5f, 0x81, 0x1c, 0x63,
	0x9a, 0x9b, 0x4d, 0xa9, 0x53, 0x24, 0x0e, 0x42, 0x72, 0x45, 0xa1, 0xfb, 0x4c, 0x04, 0xff, 0x5c,
	0x01, 0x5b, 0x6a, 0x14, 0x8b, 0x8a, 0xeb, 0xdf, 0x82, 0x2d, 0x82, 0xf9, 0x0d, 0x4d, 0x46, 0x03,
	0xe4, 0xfb, 0x09, 0x66, 0x72, 0x2c, 0x6b, 0xce, 0xfe, 0x34, 0x37, 0x3f, 0x92, 0x84, 0xef, 0x00,
	0xa0, 0xdb, 0x50, 0x9e, 0x53, 0xe9, 0xd0, 0x7f, 0x01, 0xd5, 0x38, 0x1d, 0x8e, 0x70, 0x66, 0x2c,
	0x89, 0x91, 0x6e, 0x5a, 0x72, 0x15, 0xad, 0x62, 0x15, 0xad, 0x53, 0x92, 0x39, 0x9f, 0x97, 0x66,
	0x59, 0xa0, 0xe1, 0x3f, 0x7f, 0x1f, 0x37, 0xd5, 0xda, 0x79, 0x49, 0x16, 0x73, 0x6a, 0xf5, 0xd2,
	0xe1, 0x4f, 0x38, 0x73, 0x15, 0x9d, 0x7e, 0x04, 0xd6, 0x58, 0xca, 0x62, 0x4c, 0x7c, 0x63, 0xb9,
	0xad, 0x75, 0xd6, 0x1d, 0x7d, 0x9a, 0x9b, 0x0d, 0xf5, 0x4c, 0x19, 0x80, 0x6e, 0x01, 0xd1, 0xcf,
	0x41, 0x95, 0x71, 0xc4, 0x53, 0x26, 0xc6, 0xa4, 0x71, 0x02, 0x2d, 0x45, 0x5e, 0x6c, 0xad, 0xda,
	0x62, 0xcb, 0xa1, 0xc4, 0xef, 0x0b, 0x64, 0x79, 0xc1, 0x64, 0x2e, 0x74, 0x15, 0x89, 0xfe, 0x33,
	0xa8, 0x72, 0x3a, 0xc2, 0x84, 0xa9, 0xc6, 0x7f, 0xf3, 0xea, 0xc6, 0x6f, 0x16, 0x8d, 0x1f, 0x61,
	0x02, 0x5d, 0xc5, 0xa6, 0x7f, 0x0d, 0xea, 0xf4, 0x86, 0xe0, 0xe4, 0xb9, 0xe0, 0x55, 0x41, 0x5f,
	0xea, 0xe0, 0x8b, 0x30, 0x74, 0x37, 0x85, 0x5d, 0x14, 0xdb, 0x07, 0x1b, 0x3e, 0x66, 0x5e, 0x12,
	0xc6, 0x3c, 0xa4, 0xc4, 0x58, 0x13, 0x15, 0x6f, 0x2f, 0x9c, 0x8e, 0xb3, 0x39, 0xce, 0x69, 0xcf,
	0x07, 0xb1, 0x94, 0x0e, 0x8f, 0x68, 0x14, 0x72, 0x1c, 0xc5, 0x3c, 0x73, 0xcb, 0xb4, 0xce, 0xc5,
	0xfd, 0x63, 0x4b, 0x7b, 0x78, 0x6c, 0x69, 0x6f, 0x1e, 0x5b, 0xda, 0x1f, 0x4f, 0xad, 0xca, 0xc3,
	0x53, 0xab, 0xf2, 0xef, 0x53, 0xab, 0xf2, 0xeb, 0x97, 0xa5, 0xe7, 0x2b, 0x51, 0x82, 0x79, 0x71,
	0x3c, 0xf6, 0xae, 0x51, 0x48, 0xec, 0xdb, 0xf9, 0x2f, 0x51, 0x14, 0x64, 0x58, 0x15, 0xa3, 0xf0,
	0xc5, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xd5, 0x28, 0x12, 0xdd, 0x05, 0x00, 0x00,
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
	if len(m.Slashing) > 0 {
		for iNdEx := len(m.Slashing) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Slashing[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.TotalUnissuedPrepay.Size()
		i -= size
		if _, err := m.TotalUnissuedPrepay.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InitialUozPrice.Size()
		i -= size
		if _, err := m.InitialUozPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MetaNodes) > 0 {
		for iNdEx := len(m.MetaNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MetaNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ResourceNodes) > 0 {
		for iNdEx := len(m.ResourceNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisMetaNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisMetaNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisMetaNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Description != nil {
		{
			size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Tokens.Size()
		i -= size
		if _, err := m.Tokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Status != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Suspend {
		i--
		if m.Suspend {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Pubkey != nil {
		{
			size, err := m.Pubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.NetworkAddress) > 0 {
		i -= len(m.NetworkAddress)
		copy(dAtA[i:], m.NetworkAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.NetworkAddress)))
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ResourceNodes) > 0 {
		for _, e := range m.ResourceNodes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MetaNodes) > 0 {
		for _, e := range m.MetaNodes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.InitialUozPrice.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TotalUnissuedPrepay.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Slashing) > 0 {
		for _, e := range m.Slashing {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisMetaNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NetworkAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Suspend {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovGenesis(uint64(m.Status))
	}
	l = m.Tokens.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Description != nil {
		l = m.Description.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceNodes", wireType)
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
			m.ResourceNodes = append(m.ResourceNodes, ResourceNode{})
			if err := m.ResourceNodes[len(m.ResourceNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaNodes", wireType)
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
			m.MetaNodes = append(m.MetaNodes, MetaNode{})
			if err := m.MetaNodes[len(m.MetaNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialUozPrice", wireType)
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
			if err := m.InitialUozPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUnissuedPrepay", wireType)
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
			if err := m.TotalUnissuedPrepay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slashing", wireType)
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
			m.Slashing = append(m.Slashing, &Slashing{})
			if err := m.Slashing[len(m.Slashing)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisMetaNode) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisMetaNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisMetaNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkAddress", wireType)
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
			m.NetworkAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
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
			if m.Pubkey == nil {
				m.Pubkey = &types.Any{}
			}
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspend", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Suspend = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types1.BondStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
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
			if err := m.Tokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			if m.Description == nil {
				m.Description = &Description{}
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
