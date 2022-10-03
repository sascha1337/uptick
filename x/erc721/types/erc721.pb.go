// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uptick/erc721/v1/erc721.proto

package types

import (
	fmt "fmt"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Owner enumerates the ownership of a ERC721 contract.
type Owner int32

const (
	// OWNER_UNSPECIFIED defines an invalid/undefined owner.
	OWNER_UNSPECIFIED Owner = 0
	// OWNER_MODULE erc721 is owned by the erc721 module account.
	OWNER_MODULE Owner = 1
	// EXTERNAL erc721 is owned by an external account.
	OWNER_EXTERNAL Owner = 2
)

var Owner_name = map[int32]string{
	0: "OWNER_UNSPECIFIED",
	1: "OWNER_MODULE",
	2: "OWNER_EXTERNAL",
}

var Owner_value = map[string]int32{
	"OWNER_UNSPECIFIED": 0,
	"OWNER_MODULE":      1,
	"OWNER_EXTERNAL":    2,
}

func (x Owner) String() string {
	return proto.EnumName(Owner_name, int32(x))
}

func (Owner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{0}
}

// TokenPair defines an instance that records a pairing consisting of a native
// Cosmos Coin and an ERC721 token address.
type TokenPair struct {
	// address of ERC721 contract token
	Erc721Address string `protobuf:"bytes,1,opt,name=erc721_address,json=erc721Address,proto3" json:"erc721_address,omitempty"`
	// cosmos nft class ID to be mapped to
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// shows token mapping enable status
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// ERC721 owner address ENUM (0 invalid, 1 ModuleAccount, 2 external address)
	ContractOwner Owner `protobuf:"varint,4,opt,name=contract_owner,json=contractOwner,proto3,enum=uptick.erc721.v1.Owner" json:"contract_owner,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetErc721Address() string {
	if m != nil {
		return m.Erc721Address
	}
	return ""
}

func (m *TokenPair) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *TokenPair) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *TokenPair) GetContractOwner() Owner {
	if m != nil {
		return m.ContractOwner
	}
	return OWNER_UNSPECIFIED
}

// RegisterNFTProposal is a gov Content type to register a token pair for a
// native Cosmos coin.
type RegisterNFTProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// nft class of the native Cosmos nft class
	Class nft.Class `protobuf:"bytes,3,opt,name=class,proto3" json:"class"`
}

func (m *RegisterNFTProposal) Reset()         { *m = RegisterNFTProposal{} }
func (m *RegisterNFTProposal) String() string { return proto.CompactTextString(m) }
func (*RegisterNFTProposal) ProtoMessage()    {}
func (*RegisterNFTProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{1}
}
func (m *RegisterNFTProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterNFTProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterNFTProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterNFTProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNFTProposal.Merge(m, src)
}
func (m *RegisterNFTProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterNFTProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNFTProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNFTProposal proto.InternalMessageInfo

func (m *RegisterNFTProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterNFTProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterNFTProposal) GetClass() nft.Class {
	if m != nil {
		return m.Class
	}
	return nft.Class{}
}

// RegisterERC721Proposal is a gov Content type to register a token pair for an
// ERC721 class
type RegisterERC721Proposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address of ERC721 token
	Erc721Address string `protobuf:"bytes,3,opt,name=erc721address,proto3" json:"erc721address,omitempty"`
}

func (m *RegisterERC721Proposal) Reset()         { *m = RegisterERC721Proposal{} }
func (m *RegisterERC721Proposal) String() string { return proto.CompactTextString(m) }
func (*RegisterERC721Proposal) ProtoMessage()    {}
func (*RegisterERC721Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{2}
}
func (m *RegisterERC721Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterERC721Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterERC721Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterERC721Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterERC721Proposal.Merge(m, src)
}
func (m *RegisterERC721Proposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterERC721Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterERC721Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterERC721Proposal proto.InternalMessageInfo

func (m *RegisterERC721Proposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterERC721Proposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterERC721Proposal) GetErc721Address() string {
	if m != nil {
		return m.Erc721Address
	}
	return ""
}

// ToggleTokenConversionProposal is a gov Content type to toggle the conversion
// of a token pair.
type ToggleTokenConversionProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// token identifier can be either the hex contract address of the ERC721 or
	// the Cosmos nft class
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *ToggleTokenConversionProposal) Reset()         { *m = ToggleTokenConversionProposal{} }
func (m *ToggleTokenConversionProposal) String() string { return proto.CompactTextString(m) }
func (*ToggleTokenConversionProposal) ProtoMessage()    {}
func (*ToggleTokenConversionProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{3}
}
func (m *ToggleTokenConversionProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ToggleTokenConversionProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ToggleTokenConversionProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ToggleTokenConversionProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToggleTokenConversionProposal.Merge(m, src)
}
func (m *ToggleTokenConversionProposal) XXX_Size() int {
	return m.Size()
}
func (m *ToggleTokenConversionProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ToggleTokenConversionProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ToggleTokenConversionProposal proto.InternalMessageInfo

func (m *ToggleTokenConversionProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToggleTokenConversionProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToggleTokenConversionProposal) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("uptick.erc721.v1.Owner", Owner_name, Owner_value)
	proto.RegisterType((*TokenPair)(nil), "uptick.erc721.v1.TokenPair")
	proto.RegisterType((*RegisterNFTProposal)(nil), "uptick.erc721.v1.RegisterNFTProposal")
	proto.RegisterType((*RegisterERC721Proposal)(nil), "uptick.erc721.v1.RegisterERC721Proposal")
	proto.RegisterType((*ToggleTokenConversionProposal)(nil), "uptick.erc721.v1.ToggleTokenConversionProposal")
}

func init() { proto.RegisterFile("uptick/erc721/v1/erc721.proto", fileDescriptor_e4208f03f5270a65) }

var fileDescriptor_e4208f03f5270a65 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x8d, 0xb7, 0x96, 0x6d, 0x1e, 0xab, 0x82, 0x19, 0x90, 0x55, 0x2c, 0xab, 0x2a, 0x90, 0x2a,
	0x0e, 0x89, 0x52, 0x84, 0x26, 0x71, 0x40, 0xda, 0xba, 0x4c, 0x14, 0x8d, 0xb4, 0x0a, 0xad, 0x40,
	0x5c, 0xaa, 0x34, 0x31, 0xc1, 0x6a, 0x16, 0x47, 0xb6, 0xd7, 0x81, 0xc4, 0x0f, 0x40, 0x9c, 0xf8,
	0x09, 0x48, 0x9c, 0xf9, 0x1f, 0x3b, 0xee, 0xc8, 0x09, 0xa1, 0xf6, 0xc2, 0xcf, 0x40, 0xb1, 0x1d,
	0x69, 0xda, 0x75, 0xb7, 0xef, 0xbd, 0xef, 0xe5, 0xf3, 0xcb, 0xf3, 0x67, 0xb8, 0x7b, 0x56, 0x08,
	0x12, 0xcf, 0x5c, 0xcc, 0xe2, 0xfd, 0xae, 0xe7, 0xce, 0x3d, 0x5d, 0x39, 0x05, 0xa3, 0x82, 0x22,
	0x53, 0xb5, 0x1d, 0x4d, 0xce, 0xbd, 0xe6, 0x76, 0x4a, 0x53, 0x2a, 0x9b, 0x6e, 0x59, 0x29, 0x5d,
	0xf3, 0xa1, 0x1e, 0x93, 0x7f, 0x10, 0xee, 0xdc, 0x9b, 0x62, 0x11, 0x79, 0x65, 0xad, 0xba, 0xed,
	0x5f, 0x00, 0x6e, 0x8c, 0xe8, 0x0c, 0xe7, 0xc3, 0x88, 0x30, 0xf4, 0x18, 0x36, 0xd4, 0xb8, 0x49,
	0x94, 0x24, 0x0c, 0x73, 0x6e, 0x81, 0x16, 0xe8, 0x6c, 0x84, 0x5b, 0x8a, 0x3d, 0x50, 0x24, 0xda,
	0x81, 0xeb, 0x71, 0x16, 0x71, 0x3e, 0x21, 0x89, 0xb5, 0x22, 0x05, 0x6b, 0x12, 0xf7, 0x13, 0x64,
	0xc1, 0x35, 0x9c, 0x47, 0xd3, 0x0c, 0x27, 0xd6, 0x6a, 0x0b, 0x74, 0xd6, 0xc3, 0x0a, 0xa2, 0x17,
	0xb0, 0x11, 0xd3, 0x5c, 0xb0, 0x28, 0x16, 0x13, 0x7a, 0x9e, 0x63, 0x66, 0xd5, 0x5a, 0xa0, 0xd3,
	0xe8, 0x3e, 0x70, 0xae, 0xff, 0x88, 0x33, 0x28, 0xdb, 0xe1, 0x56, 0x25, 0x97, 0xf0, 0x79, 0xed,
	0xdf, 0x8f, 0x3d, 0xd0, 0xfe, 0x06, 0xe0, 0xdd, 0x10, 0xa7, 0x84, 0x0b, 0xcc, 0x82, 0xe3, 0xd1,
	0x90, 0xd1, 0x82, 0xf2, 0x28, 0x43, 0xdb, 0xb0, 0x2e, 0x88, 0xc8, 0xb0, 0x36, 0xac, 0x00, 0x6a,
	0xc1, 0xcd, 0x04, 0xf3, 0x98, 0x91, 0x42, 0x10, 0x9a, 0x6b, 0xaf, 0x57, 0x29, 0xf4, 0x0c, 0xd6,
	0xa5, 0x75, 0xe9, 0x76, 0xb3, 0xbb, 0xe3, 0xc4, 0x94, 0x9f, 0x52, 0xee, 0x94, 0x09, 0xe9, 0xb4,
	0x9c, 0x5e, 0x29, 0x38, 0xac, 0x5d, 0xfc, 0xd9, 0x33, 0x42, 0xa5, 0x96, 0x66, 0x8c, 0xf6, 0x17,
	0x78, 0xbf, 0xf2, 0xe2, 0x87, 0xbd, 0xfd, 0xae, 0x77, 0x63, 0x3b, 0x8f, 0xa0, 0x8e, 0xba, 0xca,
	0x7f, 0xf5, 0x6a, 0xfe, 0x9a, 0xd4, 0xa7, 0x73, 0xb8, 0x3b, 0xa2, 0x69, 0x9a, 0x61, 0x79, 0x7f,
	0x3d, 0x9a, 0xcf, 0x31, 0xe3, 0x84, 0xe6, 0x37, 0x36, 0x51, 0x7e, 0x57, 0x8e, 0xd4, 0x87, 0x2b,
	0xa0, 0xf2, 0x7f, 0xf2, 0x0a, 0xd6, 0xe5, 0x75, 0xa0, 0x7b, 0xf0, 0xce, 0xe0, 0x6d, 0xe0, 0x87,
	0x93, 0x71, 0xf0, 0x66, 0xe8, 0xf7, 0xfa, 0xc7, 0x7d, 0xff, 0xc8, 0x34, 0x90, 0x09, 0x6f, 0x2b,
	0xfa, 0xf5, 0xe0, 0x68, 0x7c, 0xe2, 0x9b, 0x00, 0x21, 0xd8, 0x50, 0x8c, 0xff, 0x6e, 0xe4, 0x87,
	0xc1, 0xc1, 0x89, 0xb9, 0xd2, 0xac, 0x7d, 0xfd, 0x69, 0x1b, 0x87, 0x2f, 0x2f, 0x16, 0x36, 0xb8,
	0x5c, 0xd8, 0xe0, 0xef, 0xc2, 0x06, 0xdf, 0x97, 0xb6, 0x71, 0xb9, 0xb4, 0x8d, 0xdf, 0x4b, 0xdb,
	0x78, 0xef, 0xa4, 0x44, 0x7c, 0x3c, 0x9b, 0x3a, 0x31, 0x3d, 0x75, 0xc7, 0x72, 0x3b, 0x02, 0x2c,
	0xce, 0x29, 0x9b, 0xb9, 0x7a, 0x99, 0x3f, 0x55, 0xaf, 0x42, 0x7c, 0x2e, 0x30, 0x9f, 0xde, 0x92,
	0xcb, 0xfc, 0xf4, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xc4, 0x47, 0xb4, 0x33, 0x03, 0x00,
	0x00,
}

func (this *TokenPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPair)
	if !ok {
		that2, ok := that.(TokenPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Erc721Address != that1.Erc721Address {
		return false
	}
	if this.ClassId != that1.ClassId {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ContractOwner != that1.ContractOwner {
		return false
	}
	return true
}
func (this *ToggleTokenConversionProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ToggleTokenConversionProposal)
	if !ok {
		that2, ok := that.(ToggleTokenConversionProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	return true
}
func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ContractOwner != 0 {
		i = encodeVarintErc721(dAtA, i, uint64(m.ContractOwner))
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc721Address) > 0 {
		i -= len(m.Erc721Address)
		copy(dAtA[i:], m.Erc721Address)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Erc721Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterNFTProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterNFTProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterNFTProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Class.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintErc721(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterERC721Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterERC721Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterERC721Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc721Address) > 0 {
		i -= len(m.Erc721Address)
		copy(dAtA[i:], m.Erc721Address)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Erc721Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ToggleTokenConversionProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToggleTokenConversionProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ToggleTokenConversionProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErc721(dAtA []byte, offset int, v uint64) int {
	offset -= sovErc721(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc721Address)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.ContractOwner != 0 {
		n += 1 + sovErc721(uint64(m.ContractOwner))
	}
	return n
}

func (m *RegisterNFTProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = m.Class.Size()
	n += 1 + l + sovErc721(uint64(l))
	return n
}

func (m *RegisterERC721Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.Erc721Address)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	return n
}

func (m *ToggleTokenConversionProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	return n
}

func sovErc721(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErc721(x uint64) (n int) {
	return sovErc721(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc721Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc721Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			m.ContractOwner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractOwner |= Owner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func (m *RegisterNFTProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: RegisterNFTProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterNFTProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Class.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func (m *RegisterERC721Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: RegisterERC721Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterERC721Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc721Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc721Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func (m *ToggleTokenConversionProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: ToggleTokenConversionProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToggleTokenConversionProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func skipErc721(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErc721
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
					return 0, ErrIntOverflowErc721
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
					return 0, ErrIntOverflowErc721
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
				return 0, ErrInvalidLengthErc721
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErc721
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErc721
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErc721        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErc721          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErc721 = fmt.Errorf("proto: unexpected end of group")
)