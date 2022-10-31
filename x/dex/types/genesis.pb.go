// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params                                Params                              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TickMapList                           []TickMap                           `protobuf:"bytes,2,rep,name=tickMapList,proto3" json:"tickMapList"`
	PairMapList                           []PairMap                           `protobuf:"bytes,3,rep,name=pairMapList,proto3" json:"pairMapList"`
	TokensList                            []Tokens                            `protobuf:"bytes,4,rep,name=tokensList,proto3" json:"tokensList"`
	TokensCount                           uint64                              `protobuf:"varint,5,opt,name=tokensCount,proto3" json:"tokensCount,omitempty"`
	TokenMapList                          []TokenMap                          `protobuf:"bytes,6,rep,name=tokenMapList,proto3" json:"tokenMapList"`
	SharesList                            []Shares                            `protobuf:"bytes,7,rep,name=sharesList,proto3" json:"sharesList"`
	FeeListList                           []FeeList                           `protobuf:"bytes,8,rep,name=feeListList,proto3" json:"feeListList"`
	FeeListCount                          uint64                              `protobuf:"varint,9,opt,name=feeListCount,proto3" json:"feeListCount,omitempty"`
	EdgeRowList                           []EdgeRow                           `protobuf:"bytes,10,rep,name=edgeRowList,proto3" json:"edgeRowList"`
	EdgeRowCount                          uint64                              `protobuf:"varint,11,opt,name=edgeRowCount,proto3" json:"edgeRowCount,omitempty"`
	AdjanceyMatrixList                    []AdjanceyMatrix                    `protobuf:"bytes,12,rep,name=adjanceyMatrixList,proto3" json:"adjanceyMatrixList"`
	AdjanceyMatrixCount                   uint64                              `protobuf:"varint,13,opt,name=adjanceyMatrixCount,proto3" json:"adjanceyMatrixCount,omitempty"`
	LimitOrderPoolUserShareMapList        []LimitOrderPoolUserShareMap        `protobuf:"bytes,14,rep,name=limitOrderPoolUserShareMapList,proto3" json:"limitOrderPoolUserShareMapList"`
	LimitOrderPoolUserSharesWithdrawnList []LimitOrderPoolUserSharesWithdrawn `protobuf:"bytes,15,rep,name=limitOrderPoolUserSharesWithdrawnList,proto3" json:"limitOrderPoolUserSharesWithdrawnList"`
	LimitOrderPoolTotalSharesMapList      []LimitOrderPoolTotalSharesMap      `protobuf:"bytes,16,rep,name=limitOrderPoolTotalSharesMapList,proto3" json:"limitOrderPoolTotalSharesMapList"`
	LimitOrderPoolReserveMapList          []LimitOrderPoolReserveMap          `protobuf:"bytes,17,rep,name=limitOrderPoolReserveMapList,proto3" json:"limitOrderPoolReserveMapList"`
	LimitOrderPoolFillMapList             []LimitOrderPoolFillMap             `protobuf:"bytes,18,rep,name=limitOrderPoolFillMapList,proto3" json:"limitOrderPoolFillMapList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTickMapList() []TickMap {
	if m != nil {
		return m.TickMapList
	}
	return nil
}

func (m *GenesisState) GetPairMapList() []PairMap {
	if m != nil {
		return m.PairMapList
	}
	return nil
}

func (m *GenesisState) GetTokensList() []Tokens {
	if m != nil {
		return m.TokensList
	}
	return nil
}

func (m *GenesisState) GetTokensCount() uint64 {
	if m != nil {
		return m.TokensCount
	}
	return 0
}

func (m *GenesisState) GetTokenMapList() []TokenMap {
	if m != nil {
		return m.TokenMapList
	}
	return nil
}

func (m *GenesisState) GetSharesList() []Shares {
	if m != nil {
		return m.SharesList
	}
	return nil
}

func (m *GenesisState) GetFeeListList() []FeeList {
	if m != nil {
		return m.FeeListList
	}
	return nil
}

func (m *GenesisState) GetFeeListCount() uint64 {
	if m != nil {
		return m.FeeListCount
	}
	return 0
}

func (m *GenesisState) GetEdgeRowList() []EdgeRow {
	if m != nil {
		return m.EdgeRowList
	}
	return nil
}

func (m *GenesisState) GetEdgeRowCount() uint64 {
	if m != nil {
		return m.EdgeRowCount
	}
	return 0
}

func (m *GenesisState) GetAdjanceyMatrixList() []AdjanceyMatrix {
	if m != nil {
		return m.AdjanceyMatrixList
	}
	return nil
}

func (m *GenesisState) GetAdjanceyMatrixCount() uint64 {
	if m != nil {
		return m.AdjanceyMatrixCount
	}
	return 0
}

func (m *GenesisState) GetLimitOrderPoolUserShareMapList() []LimitOrderPoolUserShareMap {
	if m != nil {
		return m.LimitOrderPoolUserShareMapList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolUserSharesWithdrawnList() []LimitOrderPoolUserSharesWithdrawn {
	if m != nil {
		return m.LimitOrderPoolUserSharesWithdrawnList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolTotalSharesMapList() []LimitOrderPoolTotalSharesMap {
	if m != nil {
		return m.LimitOrderPoolTotalSharesMapList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolReserveMapList() []LimitOrderPoolReserveMap {
	if m != nil {
		return m.LimitOrderPoolReserveMapList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderPoolFillMapList() []LimitOrderPoolFillMap {
	if m != nil {
		return m.LimitOrderPoolFillMapList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nicholasdotsol.duality.dex.GenesisState")
}

func init() { proto.RegisterFile("dex/genesis.proto", fileDescriptor_a803aaabd08db59d) }

var fileDescriptor_a803aaabd08db59d = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x4f, 0x13, 0x4b,
	0x1c, 0xef, 0x3e, 0x78, 0xbc, 0xf7, 0xa6, 0x7d, 0x0a, 0x83, 0x07, 0x68, 0xcc, 0xda, 0x80, 0x1a,
	0xd4, 0xd8, 0x22, 0x1a, 0xe3, 0xc5, 0x44, 0x51, 0xc1, 0x44, 0x40, 0x52, 0x30, 0x26, 0x5e, 0xd6,
	0xa1, 0x3b, 0xb4, 0x23, 0xd3, 0xce, 0x66, 0x66, 0x6a, 0xcb, 0xc5, 0x93, 0x27, 0x13, 0x13, 0xaf,
	0xfe, 0x47, 0x1c, 0x39, 0x7a, 0x32, 0x06, 0xfe, 0x11, 0x33, 0xdf, 0x99, 0xb1, 0xbb, 0xa6, 0xed,
	0x36, 0xde, 0x3a, 0x9f, 0xf9, 0xfc, 0xea, 0xfc, 0x5a, 0x34, 0x17, 0xd3, 0x7e, 0xad, 0x49, 0x3b,
	0x54, 0x31, 0x55, 0x4d, 0xa4, 0xd0, 0x02, 0x97, 0x3b, 0xac, 0xd1, 0x12, 0x9c, 0xa8, 0x58, 0x68,
	0x25, 0x78, 0x35, 0xee, 0x12, 0xce, 0xf4, 0x71, 0x35, 0xa6, 0xfd, 0xf2, 0xa5, 0xa6, 0x68, 0x0a,
	0xa0, 0xd5, 0xcc, 0x2f, 0xab, 0x28, 0xcf, 0x1a, 0x93, 0x84, 0x48, 0xd2, 0x76, 0x1e, 0x65, 0x6c,
	0x10, 0xcd, 0x1a, 0x47, 0x51, 0x9b, 0x24, 0x69, 0x2c, 0x21, 0x4c, 0xa6, 0x30, 0x50, 0x6a, 0x71,
	0x44, 0x3b, 0x5e, 0x39, 0xff, 0x0b, 0xf9, 0x9d, 0xa6, 0x5a, 0x44, 0xd2, 0x4c, 0xc0, 0x21, 0xa5,
	0x11, 0x67, 0x4a, 0xa7, 0x31, 0x1a, 0x37, 0x69, 0x24, 0x45, 0xcf, 0x61, 0x8b, 0x06, 0x23, 0xf1,
	0x3b, 0xd2, 0x69, 0xd0, 0xe3, 0xa8, 0x4d, 0xb4, 0x64, 0x7d, 0x37, 0x75, 0xc3, 0x4c, 0x71, 0xd6,
	0x66, 0x3a, 0x12, 0x32, 0xa6, 0x32, 0x4a, 0x84, 0xe0, 0x51, 0x57, 0x51, 0x19, 0x41, 0x54, 0x2a,
	0x7f, 0x35, 0x87, 0xaa, 0xa2, 0x1e, 0xd3, 0xad, 0x58, 0x92, 0x5e, 0xc7, 0x29, 0x6e, 0x0d, 0x55,
	0x68, 0xa1, 0x09, 0xf7, 0x92, 0x81, 0xfd, 0xf5, 0xa1, 0x64, 0x49, 0x15, 0x95, 0xef, 0xd3, 0x35,
	0x96, 0x87, 0xf2, 0x0e, 0x19, 0xe7, 0x03, 0xd2, 0xd2, 0xe7, 0x12, 0x2a, 0x6d, 0xda, 0x0d, 0xdd,
	0xd3, 0x44, 0x53, 0xfc, 0x08, 0xcd, 0xd8, 0xbd, 0x59, 0x08, 0x2a, 0xc1, 0x4a, 0x71, 0x6d, 0xa9,
	0x3a, 0x7a, 0x83, 0xab, 0xbb, 0xc0, 0x5c, 0x9f, 0x3e, 0xf9, 0x7e, 0xa5, 0x50, 0x77, 0x3a, 0xfc,
	0x02, 0x15, 0xcd, 0x5e, 0x6e, 0x93, 0x64, 0x8b, 0x29, 0xbd, 0xf0, 0x57, 0x65, 0x6a, 0xa5, 0xb8,
	0xb6, 0x3c, 0xce, 0x66, 0xdf, 0xd2, 0x9d, 0x4f, 0x5a, 0x6d, 0xcc, 0xcc, 0x21, 0xf0, 0x66, 0x53,
	0xf9, 0x66, 0xbb, 0x96, 0xee, 0xcd, 0x52, 0x6a, 0xfc, 0x1c, 0x21, 0x7b, 0x7a, 0xc0, 0x6b, 0x1a,
	0xbc, 0xc6, 0xfe, 0xbf, 0x7d, 0x60, 0x3b, 0xab, 0x94, 0x16, 0x57, 0x50, 0xd1, 0x8e, 0x9e, 0x88,
	0x6e, 0x47, 0x2f, 0xfc, 0x5d, 0x09, 0x56, 0xa6, 0xeb, 0x69, 0x08, 0xef, 0xa0, 0x12, 0x0c, 0x7d,
	0xf3, 0x19, 0x48, 0xbb, 0x9a, 0x9b, 0x36, 0xa8, 0x9e, 0xd1, 0x9b, 0xee, 0xf6, 0x24, 0x80, 0xdb,
	0x3f, 0xf9, 0xdd, 0xf7, 0x80, 0xed, 0xbb, 0x0f, 0xb4, 0x66, 0x49, 0x0f, 0x29, 0x35, 0x3f, 0xc1,
	0xea, 0xdf, 0xfc, 0x25, 0xdd, 0xb0, 0x74, 0xbf, 0xa4, 0x29, 0x35, 0x5e, 0x42, 0x25, 0x37, 0xb4,
	0x2b, 0xf1, 0x1f, 0xac, 0x44, 0x06, 0x33, 0x81, 0xe6, 0x9e, 0xd5, 0x45, 0x0f, 0x02, 0x51, 0x7e,
	0xe0, 0x33, 0x4b, 0xf7, 0x81, 0x29, 0xb5, 0x09, 0x74, 0x43, 0x1b, 0x58, 0xb4, 0x81, 0x69, 0x0c,
	0xbf, 0x45, 0xd8, 0x5f, 0xe2, 0x6d, 0xb8, 0xc3, 0x90, 0x5b, 0x82, 0xdc, 0x9b, 0xe3, 0x72, 0x1f,
	0x67, 0x54, 0x2e, 0x7e, 0x88, 0x17, 0x5e, 0x45, 0xf3, 0x59, 0xd4, 0x96, 0xf9, 0x1f, 0xca, 0x0c,
	0x9b, 0xc2, 0x1f, 0x03, 0x14, 0xc2, 0x65, 0x7c, 0x69, 0xee, 0xe2, 0xae, 0x10, 0xfc, 0x95, 0xa2,
	0x12, 0x36, 0xc9, 0x1f, 0x91, 0x0b, 0x50, 0xf0, 0xfe, 0xb8, 0x82, 0x5b, 0x23, 0x1d, 0x5c, 0xd9,
	0x9c, 0x0c, 0xfc, 0x35, 0x40, 0xd7, 0x46, 0x50, 0xd4, 0x6b, 0xff, 0x2a, 0x41, 0x9b, 0x8b, 0xd0,
	0xe6, 0xe1, 0x1f, 0xb4, 0x19, 0x18, 0xb9, 0x52, 0x93, 0x25, 0xe2, 0x4f, 0x01, 0xaa, 0x64, 0x99,
	0xfb, 0xe6, 0x05, 0xb4, 0x54, 0xbf, 0x48, 0xb3, 0x50, 0xeb, 0xc1, 0xe4, 0xb5, 0xb2, 0x1e, 0xae,
	0x51, 0x6e, 0x0e, 0xfe, 0x80, 0x2e, 0x67, 0x39, 0x75, 0xfb, 0xc0, 0xfa, 0x1e, 0x73, 0xd0, 0xe3,
	0xde, 0xe4, 0x3d, 0x06, 0x7a, 0xd7, 0x61, 0xac, 0x3f, 0xee, 0xa2, 0xc5, 0xec, 0xfc, 0x06, 0xe3,
	0xdc, 0x87, 0x63, 0x08, 0xbf, 0x33, 0x79, 0xb8, 0x13, 0xbb, 0xe4, 0xd1, 0xce, 0xeb, 0x9b, 0x27,
	0x67, 0x61, 0x70, 0x7a, 0x16, 0x06, 0x3f, 0xce, 0xc2, 0xe0, 0xcb, 0x79, 0x58, 0x38, 0x3d, 0x0f,
	0x0b, 0xdf, 0xce, 0xc3, 0xc2, 0x9b, 0xdb, 0x4d, 0xa6, 0x5b, 0xdd, 0x83, 0x6a, 0x43, 0xb4, 0x6b,
	0x3b, 0x2e, 0xf7, 0xa9, 0xd0, 0x7b, 0x82, 0xd7, 0x5c, 0x6e, 0xad, 0x5f, 0x83, 0xcf, 0xf1, 0x71,
	0x42, 0xd5, 0xc1, 0x0c, 0x7c, 0x5f, 0xee, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xb9, 0xc4,
	0x7c, 0x33, 0x08, 0x00, 0x00,
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
	if len(m.LimitOrderPoolFillMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolFillMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolFillMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x92
		}
	}
	if len(m.LimitOrderPoolReserveMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolReserveMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolReserveMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
	}
	if len(m.LimitOrderPoolTotalSharesMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolTotalSharesMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolTotalSharesMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.LimitOrderPoolUserSharesWithdrawnList) > 0 {
		for iNdEx := len(m.LimitOrderPoolUserSharesWithdrawnList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolUserSharesWithdrawnList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.LimitOrderPoolUserShareMapList) > 0 {
		for iNdEx := len(m.LimitOrderPoolUserShareMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderPoolUserShareMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if m.AdjanceyMatrixCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AdjanceyMatrixCount))
		i--
		dAtA[i] = 0x68
	}
	if len(m.AdjanceyMatrixList) > 0 {
		for iNdEx := len(m.AdjanceyMatrixList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdjanceyMatrixList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.EdgeRowCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EdgeRowCount))
		i--
		dAtA[i] = 0x58
	}
	if len(m.EdgeRowList) > 0 {
		for iNdEx := len(m.EdgeRowList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EdgeRowList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.FeeListCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.FeeListCount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.FeeListList) > 0 {
		for iNdEx := len(m.FeeListList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeListList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.SharesList) > 0 {
		for iNdEx := len(m.SharesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SharesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.TokenMapList) > 0 {
		for iNdEx := len(m.TokenMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.TokensCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TokensCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TokensList) > 0 {
		for iNdEx := len(m.TokensList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokensList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PairMapList) > 0 {
		for iNdEx := len(m.PairMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TickMapList) > 0 {
		for iNdEx := len(m.TickMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TickMapList) > 0 {
		for _, e := range m.TickMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PairMapList) > 0 {
		for _, e := range m.PairMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokensList) > 0 {
		for _, e := range m.TokensList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TokensCount != 0 {
		n += 1 + sovGenesis(uint64(m.TokensCount))
	}
	if len(m.TokenMapList) > 0 {
		for _, e := range m.TokenMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SharesList) > 0 {
		for _, e := range m.SharesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FeeListList) > 0 {
		for _, e := range m.FeeListList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.FeeListCount != 0 {
		n += 1 + sovGenesis(uint64(m.FeeListCount))
	}
	if len(m.EdgeRowList) > 0 {
		for _, e := range m.EdgeRowList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EdgeRowCount != 0 {
		n += 1 + sovGenesis(uint64(m.EdgeRowCount))
	}
	if len(m.AdjanceyMatrixList) > 0 {
		for _, e := range m.AdjanceyMatrixList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.AdjanceyMatrixCount != 0 {
		n += 1 + sovGenesis(uint64(m.AdjanceyMatrixCount))
	}
	if len(m.LimitOrderPoolUserShareMapList) > 0 {
		for _, e := range m.LimitOrderPoolUserShareMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolUserSharesWithdrawnList) > 0 {
		for _, e := range m.LimitOrderPoolUserSharesWithdrawnList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolTotalSharesMapList) > 0 {
		for _, e := range m.LimitOrderPoolTotalSharesMapList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolReserveMapList) > 0 {
		for _, e := range m.LimitOrderPoolReserveMapList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderPoolFillMapList) > 0 {
		for _, e := range m.LimitOrderPoolFillMapList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickMapList", wireType)
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
			m.TickMapList = append(m.TickMapList, TickMap{})
			if err := m.TickMapList[len(m.TickMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairMapList", wireType)
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
			m.PairMapList = append(m.PairMapList, PairMap{})
			if err := m.PairMapList[len(m.PairMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensList", wireType)
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
			m.TokensList = append(m.TokensList, Tokens{})
			if err := m.TokensList[len(m.TokensList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensCount", wireType)
			}
			m.TokensCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokensCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMapList", wireType)
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
			m.TokenMapList = append(m.TokenMapList, TokenMap{})
			if err := m.TokenMapList[len(m.TokenMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesList", wireType)
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
			m.SharesList = append(m.SharesList, Shares{})
			if err := m.SharesList[len(m.SharesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeListList", wireType)
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
			m.FeeListList = append(m.FeeListList, FeeList{})
			if err := m.FeeListList[len(m.FeeListList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeListCount", wireType)
			}
			m.FeeListCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeListCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdgeRowList", wireType)
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
			m.EdgeRowList = append(m.EdgeRowList, EdgeRow{})
			if err := m.EdgeRowList[len(m.EdgeRowList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdgeRowCount", wireType)
			}
			m.EdgeRowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EdgeRowCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjanceyMatrixList", wireType)
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
			m.AdjanceyMatrixList = append(m.AdjanceyMatrixList, AdjanceyMatrix{})
			if err := m.AdjanceyMatrixList[len(m.AdjanceyMatrixList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjanceyMatrixCount", wireType)
			}
			m.AdjanceyMatrixCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdjanceyMatrixCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolUserShareMapList", wireType)
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
			m.LimitOrderPoolUserShareMapList = append(m.LimitOrderPoolUserShareMapList, LimitOrderPoolUserShareMap{})
			if err := m.LimitOrderPoolUserShareMapList[len(m.LimitOrderPoolUserShareMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolUserSharesWithdrawnList", wireType)
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
			m.LimitOrderPoolUserSharesWithdrawnList = append(m.LimitOrderPoolUserSharesWithdrawnList, LimitOrderPoolUserSharesWithdrawn{})
			if err := m.LimitOrderPoolUserSharesWithdrawnList[len(m.LimitOrderPoolUserSharesWithdrawnList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolTotalSharesMapList", wireType)
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
			m.LimitOrderPoolTotalSharesMapList = append(m.LimitOrderPoolTotalSharesMapList, LimitOrderPoolTotalSharesMap{})
			if err := m.LimitOrderPoolTotalSharesMapList[len(m.LimitOrderPoolTotalSharesMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolReserveMapList", wireType)
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
			m.LimitOrderPoolReserveMapList = append(m.LimitOrderPoolReserveMapList, LimitOrderPoolReserveMap{})
			if err := m.LimitOrderPoolReserveMapList[len(m.LimitOrderPoolReserveMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderPoolFillMapList", wireType)
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
			m.LimitOrderPoolFillMapList = append(m.LimitOrderPoolFillMapList, LimitOrderPoolFillMap{})
			if err := m.LimitOrderPoolFillMapList[len(m.LimitOrderPoolFillMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
