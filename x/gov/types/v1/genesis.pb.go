// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gov/v1/genesis.proto

package v1

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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// deposits defines all the deposits present at genesis.
	Deposits []*Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits,omitempty"`
	// votes defines all the votes present at genesis.
	Votes []*Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
	// proposals defines all the proposals present at genesis.
	Proposals []*Proposal `protobuf:"bytes,4,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// params defines all the paramaters of related to deposit.
	DepositParams *DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params,omitempty"`
	// params defines all the paramaters of related to voting.
	VotingParams *VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params,omitempty"`
	// params defines all the paramaters of related to tally.
	TallyParams *TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef7cfd15e3ded621, []int{0}
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

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetDeposits() []*Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() []*Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() *DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return nil
}

func (m *GenesisState) GetVotingParams() *VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return nil
}

func (m *GenesisState) GetTallyParams() *TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.gov.v1.GenesisState")
}

func init() { proto.RegisterFile("cosmos/gov/v1/genesis.proto", fileDescriptor_ef7cfd15e3ded621) }

var fileDescriptor_ef7cfd15e3ded621 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0xa9, 0xfc, 0xa8, 0x03, 0xb8, 0x18, 0x8d, 0x34, 0x60, 0x1a, 0xe2, 0x0a, 0x63, 0x6c,
	0x05, 0xe3, 0xd2, 0xc4, 0xf8, 0x13, 0xe3, 0x8e, 0x54, 0xe3, 0xc2, 0x0d, 0x29, 0x74, 0x52, 0x1b,
	0x81, 0xdb, 0xf4, 0x5e, 0x27, 0xf2, 0x16, 0x3e, 0x80, 0x0f, 0xe4, 0x92, 0xa5, 0x4b, 0x03, 0x2f,
	0x62, 0x98, 0x69, 0x05, 0xab, 0xab, 0xc9, 0xcd, 0xf9, 0xce, 0x99, 0x93, 0x7b, 0x59, 0x63, 0x00,
	0x38, 0x02, 0x74, 0x02, 0x90, 0x8e, 0x6c, 0x3b, 0x81, 0x18, 0x0b, 0x0c, 0xd1, 0x8e, 0x62, 0x20,
	0xe0, 0x55, 0x2d, 0xda, 0x01, 0x48, 0x5b, 0xb6, 0xeb, 0xb5, 0x0c, 0x0b, 0x52, 0x73, 0xfb, 0xef,
	0x79, 0x56, 0xb9, 0xd1, 0xce, 0x3b, 0xf2, 0x48, 0xf0, 0x63, 0xb6, 0x83, 0xe4, 0xc5, 0x14, 0x8e,
	0x83, 0x5e, 0x14, 0x43, 0x04, 0xe8, 0x0d, 0x7b, 0xa1, 0x6f, 0x1a, 0x4d, 0xa3, 0x55, 0x70, 0x79,
	0xaa, 0x75, 0x13, 0xe9, 0xd6, 0xe7, 0x1d, 0xb6, 0xe1, 0x8b, 0x08, 0x30, 0x24, 0x34, 0xd7, 0x9a,
	0xf9, 0x56, 0xb9, 0xb3, 0x6b, 0xff, 0xfa, 0xdd, 0xbe, 0xd2, 0xb2, 0xfb, 0xc3, 0xf1, 0x03, 0x56,
	0x94, 0x40, 0x02, 0xcd, 0xbc, 0x32, 0x6c, 0x67, 0x0c, 0x0f, 0x40, 0xc2, 0xd5, 0x04, 0x3f, 0x65,
	0x9b, 0x69, 0x0f, 0x34, 0x0b, 0x0a, 0xaf, 0x65, 0xf0, 0xb4, 0x8c, 0xbb, 0x24, 0xf9, 0x25, 0xdb,
	0x4a, 0x7e, 0xeb, 0x45, 0x5e, 0xec, 0x8d, 0xd0, 0x2c, 0x36, 0x8d, 0x56, 0xb9, 0xb3, 0xf7, 0x7f,
	0xb7, 0xae, 0x62, 0xdc, 0xaa, 0xbf, 0x3a, 0xf2, 0x73, 0x56, 0x95, 0xa0, 0x57, 0xa1, 0x33, 0x4a,
	0x2a, 0xa3, 0xf1, 0xb7, 0xee, 0x62, 0x25, 0x3a, 0xa2, 0x22, 0x57, 0x26, 0x7e, 0xc6, 0x2a, 0xe4,
	0x0d, 0x87, 0x93, 0x34, 0x60, 0x5d, 0x05, 0xd4, 0x33, 0x01, 0xf7, 0x0b, 0x24, 0xf1, 0x97, 0x69,
	0x39, 0x5c, 0x5c, 0x7f, 0xcc, 0x2c, 0x63, 0x3a, 0xb3, 0x8c, 0xaf, 0x99, 0x65, 0xbc, 0xcd, 0xad,
	0xdc, 0x74, 0x6e, 0xe5, 0x3e, 0xe7, 0x56, 0xee, 0xf1, 0x30, 0x08, 0xe9, 0xe9, 0xa5, 0x6f, 0x0f,
	0x60, 0xe4, 0x24, 0xc7, 0xd5, 0xcf, 0x11, 0xfa, 0xcf, 0xce, 0xab, 0xba, 0x34, 0x4d, 0x22, 0x81,
	0x8e, 0x6c, 0xf7, 0x4b, 0xea, 0xd8, 0x27, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xec, 0x74,
	0x03, 0x33, 0x02, 0x00, 0x00,
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
	if m.TallyParams != nil {
		{
			size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.VotingParams != nil {
		{
			size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DepositParams != nil {
		{
			size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
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
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DepositParams != nil {
		l = m.DepositParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.VotingParams != nil {
		l = m.VotingParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TallyParams != nil {
		l = m.TallyParams.Size()
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, &Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, &Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
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
			if m.DepositParams == nil {
				m.DepositParams = &DepositParams{}
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
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
			if m.VotingParams == nil {
				m.VotingParams = &VotingParams{}
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
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
			if m.TallyParams == nil {
				m.TallyParams = &TallyParams{}
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
