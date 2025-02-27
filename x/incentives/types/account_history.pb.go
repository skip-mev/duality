// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dualitylabs/duality/incentives/account_history.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Describes the total distributions to an account over time
type AccountHistory struct {
	// the address of this account
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// coins describes the total amount of coins that have been distributed to this user over time
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *AccountHistory) Reset()         { *m = AccountHistory{} }
func (m *AccountHistory) String() string { return proto.CompactTextString(m) }
func (*AccountHistory) ProtoMessage()    {}
func (*AccountHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e34734062a8723e, []int{0}
}
func (m *AccountHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountHistory.Merge(m, src)
}
func (m *AccountHistory) XXX_Size() int {
	return m.Size()
}
func (m *AccountHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountHistory.DiscardUnknown(m)
}

var xxx_messageInfo_AccountHistory proto.InternalMessageInfo

func (m *AccountHistory) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AccountHistory) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountHistory)(nil), "dualitylabs.duality.incentives.AccountHistory")
}

func init() {
	proto.RegisterFile("dualitylabs/duality/incentives/account_history.proto", fileDescriptor_3e34734062a8723e)
}

var fileDescriptor_3e34734062a8723e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0x29, 0x4d, 0xcc,
	0xc9, 0x2c, 0xa9, 0xcc, 0x49, 0x4c, 0x2a, 0xd6, 0x87, 0xb2, 0xf5, 0x33, 0xf3, 0x92, 0x53, 0xf3,
	0x4a, 0x32, 0xcb, 0x52, 0x8b, 0xf5, 0x13, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0xe2, 0x33, 0x32,
	0x8b, 0x4b, 0xf2, 0x8b, 0x2a, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x90, 0x74, 0xe9,
	0x41, 0xd9, 0x7a, 0x08, 0x5d, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa5, 0xfa, 0x20, 0x16,
	0x44, 0x97, 0x94, 0x5c, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0xaa, 0x7e,
	0x99, 0x61, 0x52, 0x6a, 0x49, 0xa2, 0xa1, 0x7e, 0x72, 0x7e, 0x66, 0x1e, 0x44, 0x5e, 0xa9, 0x97,
	0x91, 0x8b, 0xcf, 0x11, 0x62, 0x9f, 0x07, 0xc4, 0x3a, 0x21, 0x09, 0x2e, 0x76, 0xa8, 0x0b, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xa1, 0x44, 0x2e, 0x56, 0x90, 0xd6, 0x62, 0x09,
	0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x49, 0x3d, 0x88, 0xe1, 0x7a, 0x20, 0xc3, 0xf5, 0xa0, 0x86,
	0xeb, 0x39, 0xe7, 0x67, 0xe6, 0x39, 0x19, 0x9c, 0xb8, 0x27, 0xcf, 0xb0, 0xea, 0xbe, 0xbc, 0x46,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4, 0x25, 0x10, 0x4a, 0xb7,
	0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0xac, 0xa1, 0x38, 0x08, 0x62, 0xb2, 0x93,
	0xcf, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0x21, 0x19, 0x05, 0xf5,
	0xbe, 0x2e, 0x4a, 0x08, 0x56, 0x20, 0x87, 0x21, 0xd8, 0xe8, 0x24, 0x36, 0xb0, 0x27, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0x9a, 0x15, 0x2b, 0x72, 0x01, 0x00, 0x00,
}

func (m *AccountHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountHistory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountHistory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountHistory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintAccountHistory(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountHistory(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountHistory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountHistory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovAccountHistory(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovAccountHistory(uint64(l))
		}
	}
	return n
}

func sovAccountHistory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountHistory(x uint64) (n int) {
	return sovAccountHistory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountHistory
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
			return fmt.Errorf("proto: AccountHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountHistory
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
				return ErrInvalidLengthAccountHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountHistory
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
				return ErrInvalidLengthAccountHistory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountHistory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountHistory
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
func skipAccountHistory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountHistory
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
					return 0, ErrIntOverflowAccountHistory
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
					return 0, ErrIntOverflowAccountHistory
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
				return 0, ErrInvalidLengthAccountHistory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountHistory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountHistory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountHistory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountHistory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountHistory = fmt.Errorf("proto: unexpected end of group")
)
