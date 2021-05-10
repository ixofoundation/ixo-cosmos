// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: project/genesis.proto

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

// GenesisState defines the did module's genesis state.
type GenesisState struct {
	ProjectDocs      []ProjectDoc             `protobuf:"bytes,1,rep,name=project_docs,json=projectDocs,proto3" json:"project_docs"`
	AccountMaps      []GenesisAccountMap      `protobuf:"bytes,2,rep,name=account_maps,json=accountMaps,proto3" json:"account_maps"`
	WithdrawalsInfos []WithdrawalInfoDocsList `protobuf:"bytes,3,rep,name=withdrawals_infos,json=withdrawalsInfos,proto3" json:"withdrawals_infos"`
	Claims           []ClaimsList             `protobuf:"bytes,4,rep,name=claims,proto3" json:"claims"`
	Params           Params                   `protobuf:"bytes,5,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4845f246be6bdeb0, []int{0}
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

func (m *GenesisState) GetProjectDocs() []ProjectDoc {
	if m != nil {
		return m.ProjectDocs
	}
	return nil
}

func (m *GenesisState) GetAccountMaps() []GenesisAccountMap {
	if m != nil {
		return m.AccountMaps
	}
	return nil
}

func (m *GenesisState) GetWithdrawalsInfos() []WithdrawalInfoDocsList {
	if m != nil {
		return m.WithdrawalsInfos
	}
	return nil
}

func (m *GenesisState) GetClaims() []ClaimsList {
	if m != nil {
		return m.Claims
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
	proto.RegisterType((*GenesisState)(nil), "project.GenesisState")
}

func init() { proto.RegisterFile("project/genesis.proto", fileDescriptor_4845f246be6bdeb0) }

var fileDescriptor_4845f246be6bdeb0 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3f, 0x4f, 0x32, 0x41,
	0x10, 0xc6, 0xef, 0x80, 0x97, 0x37, 0x59, 0x48, 0xd4, 0x53, 0x93, 0x0b, 0xc5, 0x41, 0xac, 0x68,
	0xe0, 0x22, 0x16, 0x36, 0x36, 0x82, 0x89, 0x31, 0xd1, 0x44, 0xb1, 0x30, 0xb1, 0x21, 0xc3, 0x72,
	0x1c, 0xab, 0xdc, 0xce, 0xe6, 0x66, 0x09, 0xf8, 0x2d, 0xfc, 0x56, 0x52, 0x52, 0x5a, 0x19, 0xc3,
	0x7d, 0x11, 0x73, 0x7f, 0x25, 0x56, 0xb3, 0xfb, 0x9b, 0xe7, 0x79, 0x26, 0xbb, 0xc3, 0x8e, 0x55,
	0x88, 0x2f, 0x1e, 0xd7, 0xae, 0xef, 0x49, 0x8f, 0x04, 0x75, 0x55, 0x88, 0x1a, 0xad, 0xff, 0x19,
	0x6e, 0x1c, 0xf9, 0xe8, 0x63, 0xc2, 0xdc, 0xf8, 0x94, 0xb6, 0x1b, 0x85, 0x2b, 0xab, 0x29, 0x3e,
	0xf9, 0x28, 0xb1, 0xfa, 0x75, 0x9a, 0xf3, 0xa8, 0x41, 0x7b, 0xd6, 0x05, 0xab, 0x67, 0x8a, 0xd1,
	0x04, 0x39, 0xd9, 0x66, 0xab, 0xdc, 0xae, 0xf5, 0x0e, 0xbb, 0xb9, 0xed, 0x3e, 0xad, 0x57, 0xc8,
	0xfb, 0x95, 0xf5, 0x57, 0xd3, 0x18, 0xd6, 0x54, 0x41, 0xc8, 0x1a, 0xb0, 0x3a, 0x70, 0x8e, 0x0b,
	0xa9, 0x47, 0x01, 0x28, 0xb2, 0x4b, 0x89, 0xbb, 0x51, 0xb8, 0xb3, 0x51, 0x97, 0xa9, 0xe6, 0x0e,
	0x54, 0x1e, 0x02, 0x05, 0x21, 0x6b, 0xc8, 0x0e, 0x96, 0x42, 0xcf, 0x26, 0x21, 0x2c, 0x61, 0x4e,
	0x23, 0x21, 0xa7, 0x48, 0x76, 0x39, 0x49, 0x6a, 0x16, 0x49, 0x4f, 0x85, 0xe2, 0x46, 0x4e, 0x31,
	0x1e, 0x7e, 0x2b, 0x48, 0x67, 0x71, 0xfb, 0x3b, 0xfe, 0xb8, 0x4d, 0xd6, 0x29, 0xab, 0xf2, 0x39,
	0x88, 0x80, 0xec, 0xca, 0x9f, 0x07, 0x0d, 0x12, 0xbc, 0x63, 0xce, 0x84, 0x56, 0x87, 0x55, 0x15,
	0x84, 0x10, 0x90, 0xfd, 0xaf, 0x65, 0xb6, 0x6b, 0xbd, 0xbd, 0xdf, 0x3f, 0x48, 0x70, 0x2e, 0x4f,
	0x45, 0xfd, 0x87, 0xf5, 0xd6, 0x31, 0x37, 0x5b, 0xc7, 0xfc, 0xde, 0x3a, 0xe6, 0x7b, 0xe4, 0x18,
	0x9b, 0xc8, 0x31, 0x3e, 0x23, 0xc7, 0x78, 0x3e, 0xf7, 0x85, 0x9e, 0x2d, 0xc6, 0x5d, 0x8e, 0x81,
	0x2b, 0x56, 0x38, 0xc5, 0x85, 0x9c, 0x80, 0x16, 0x28, 0xe3, 0x5b, 0x67, 0x3c, 0x47, 0xfe, 0xca,
	0x67, 0x20, 0xa4, 0xbb, 0xca, 0x97, 0xe3, 0xea, 0x37, 0xe5, 0xd1, 0xb8, 0x9a, 0xec, 0xe8, 0xec,
	0x27, 0x00, 0x00, 0xff, 0xff, 0xea, 0x50, 0xd4, 0x29, 0xf2, 0x01, 0x00, 0x00,
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
	dAtA[i] = 0x2a
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.WithdrawalsInfos) > 0 {
		for iNdEx := len(m.WithdrawalsInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawalsInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.AccountMaps) > 0 {
		for iNdEx := len(m.AccountMaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountMaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ProjectDocs) > 0 {
		for iNdEx := len(m.ProjectDocs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectDocs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ProjectDocs) > 0 {
		for _, e := range m.ProjectDocs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AccountMaps) > 0 {
		for _, e := range m.AccountMaps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WithdrawalsInfos) > 0 {
		for _, e := range m.WithdrawalsInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
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
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectDocs", wireType)
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
			m.ProjectDocs = append(m.ProjectDocs, ProjectDoc{})
			if err := m.ProjectDocs[len(m.ProjectDocs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountMaps", wireType)
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
			m.AccountMaps = append(m.AccountMaps, GenesisAccountMap{})
			if err := m.AccountMaps[len(m.AccountMaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalsInfos", wireType)
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
			m.WithdrawalsInfos = append(m.WithdrawalsInfos, WithdrawalInfoDocsList{})
			if err := m.WithdrawalsInfos[len(m.WithdrawalsInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
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
			m.Claims = append(m.Claims, ClaimsList{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
