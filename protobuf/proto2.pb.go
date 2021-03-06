// Code generated by protoc-gen-go.
// source: proto2.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	proto2.proto
	proto3.proto

It has these top-level messages:
	LargerP2Define
	SmallerP2Define
	LargerP2ToP3Define
	LargerP3Define
	SmallerP3Define
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LargerP2Define struct {
	Provider         *uint32 `protobuf:"varint,1,opt,name=provider" json:"provider,omitempty"`
	CompetitionId    *uint32 `protobuf:"varint,2,opt,name=competition_id,json=competitionId" json:"competition_id,omitempty"`
	EventId          *uint32 `protobuf:"varint,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	MarketId         *uint32 `protobuf:"varint,4,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LargerP2Define) Reset()                    { *m = LargerP2Define{} }
func (m *LargerP2Define) String() string            { return proto.CompactTextString(m) }
func (*LargerP2Define) ProtoMessage()               {}
func (*LargerP2Define) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LargerP2Define) GetProvider() uint32 {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return 0
}

func (m *LargerP2Define) GetCompetitionId() uint32 {
	if m != nil && m.CompetitionId != nil {
		return *m.CompetitionId
	}
	return 0
}

func (m *LargerP2Define) GetEventId() uint32 {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return 0
}

func (m *LargerP2Define) GetMarketId() uint32 {
	if m != nil && m.MarketId != nil {
		return *m.MarketId
	}
	return 0
}

type SmallerP2Define struct {
	CompetitionId    *uint32 `protobuf:"varint,2,opt,name=competition_id,json=competitionId" json:"competition_id,omitempty"`
	MarketId         *uint32 `protobuf:"varint,4,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmallerP2Define) Reset()                    { *m = SmallerP2Define{} }
func (m *SmallerP2Define) String() string            { return proto.CompactTextString(m) }
func (*SmallerP2Define) ProtoMessage()               {}
func (*SmallerP2Define) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SmallerP2Define) GetCompetitionId() uint32 {
	if m != nil && m.CompetitionId != nil {
		return *m.CompetitionId
	}
	return 0
}

func (m *SmallerP2Define) GetMarketId() uint32 {
	if m != nil && m.MarketId != nil {
		return *m.MarketId
	}
	return 0
}

type LargerP2ToP3Define struct {
	Provider         *uint32  `protobuf:"varint,1,opt,name=provider" json:"provider,omitempty"`
	CompetitionId    *uint32  `protobuf:"varint,2,opt,name=competition_id,json=competitionId" json:"competition_id,omitempty"`
	EventId          *uint32  `protobuf:"varint,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	MarketId         *uint32  `protobuf:"varint,4,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	Outcome          *string  `protobuf:"bytes,5,opt,name=outcome" json:"outcome,omitempty"`
	SpecialBetValue  *string  `protobuf:"bytes,6,opt,name=special_bet_value,json=specialBetValue" json:"special_bet_value,omitempty"`
	Odds             *float64 `protobuf:"fixed64,7,opt,name=odds" json:"odds,omitempty"`
	IsLive           *bool    `protobuf:"varint,8,opt,name=is_live,json=isLive" json:"is_live,omitempty"`
	CurrencyPair     *string  `protobuf:"bytes,9,opt,name=currency_pair,json=currencyPair" json:"currency_pair,omitempty"`
	StakeFactor      *float64 `protobuf:"fixed64,10,opt,name=stake_factor,json=stakeFactor" json:"stake_factor,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *LargerP2ToP3Define) Reset()                    { *m = LargerP2ToP3Define{} }
func (m *LargerP2ToP3Define) String() string            { return proto.CompactTextString(m) }
func (*LargerP2ToP3Define) ProtoMessage()               {}
func (*LargerP2ToP3Define) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LargerP2ToP3Define) GetProvider() uint32 {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return 0
}

func (m *LargerP2ToP3Define) GetCompetitionId() uint32 {
	if m != nil && m.CompetitionId != nil {
		return *m.CompetitionId
	}
	return 0
}

func (m *LargerP2ToP3Define) GetEventId() uint32 {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return 0
}

func (m *LargerP2ToP3Define) GetMarketId() uint32 {
	if m != nil && m.MarketId != nil {
		return *m.MarketId
	}
	return 0
}

func (m *LargerP2ToP3Define) GetOutcome() string {
	if m != nil && m.Outcome != nil {
		return *m.Outcome
	}
	return ""
}

func (m *LargerP2ToP3Define) GetSpecialBetValue() string {
	if m != nil && m.SpecialBetValue != nil {
		return *m.SpecialBetValue
	}
	return ""
}

func (m *LargerP2ToP3Define) GetOdds() float64 {
	if m != nil && m.Odds != nil {
		return *m.Odds
	}
	return 0
}

func (m *LargerP2ToP3Define) GetIsLive() bool {
	if m != nil && m.IsLive != nil {
		return *m.IsLive
	}
	return false
}

func (m *LargerP2ToP3Define) GetCurrencyPair() string {
	if m != nil && m.CurrencyPair != nil {
		return *m.CurrencyPair
	}
	return ""
}

func (m *LargerP2ToP3Define) GetStakeFactor() float64 {
	if m != nil && m.StakeFactor != nil {
		return *m.StakeFactor
	}
	return 0
}

func init() {
	proto.RegisterType((*LargerP2Define)(nil), "protobuf.LargerP2Define")
	proto.RegisterType((*SmallerP2Define)(nil), "protobuf.SmallerP2Define")
	proto.RegisterType((*LargerP2ToP3Define)(nil), "protobuf.LargerP2ToP3Define")
}

func init() { proto.RegisterFile("proto2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x8f, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0x95, 0xfe, 0xfd, 0x9b, 0x74, 0xe8, 0x45, 0xcc, 0x06, 0x03, 0x1b, 0x28, 0x42, 0x42,
	0x2c, 0x58, 0x94, 0x37, 0x40, 0x08, 0xa9, 0x52, 0x17, 0x55, 0xb8, 0x6c, 0x23, 0x37, 0x9e, 0x22,
	0xab, 0x49, 0x1c, 0x39, 0x4e, 0x24, 0x9e, 0x81, 0x37, 0xe3, 0xa9, 0x70, 0xa6, 0x04, 0x75, 0x85,
	0xd8, 0xb1, 0x1a, 0x9f, 0xef, 0x1c, 0x1f, 0x7b, 0x60, 0x54, 0x5a, 0xe3, 0xcc, 0xfc, 0x86, 0x07,
	0x46, 0x3c, 0xd6, 0xf5, 0x66, 0xf6, 0x1e, 0xc0, 0x64, 0x29, 0xed, 0x2b, 0xd9, 0xd5, 0xfc, 0x9e,
	0x36, 0xba, 0x20, 0x3c, 0x81, 0xd6, 0x6e, 0xb4, 0x22, 0x2b, 0x82, 0xb3, 0xe0, 0x6a, 0x1c, 0x7f,
	0x6b, 0xbc, 0x84, 0x49, 0x6a, 0xf2, 0x92, 0x9c, 0x76, 0xda, 0x14, 0x89, 0x56, 0xa2, 0xc7, 0x89,
	0xf1, 0x1e, 0x5d, 0x28, 0x3c, 0x86, 0x88, 0x1a, 0x2a, 0x5c, 0x1b, 0xf8, 0xc7, 0x81, 0x90, 0xb5,
	0xb7, 0x4e, 0x61, 0x98, 0x4b, 0xbb, 0x25, 0xf6, 0xfa, 0xbb, 0xfa, 0x1d, 0x58, 0xa8, 0xd9, 0x33,
	0x4c, 0x1f, 0x73, 0x99, 0x65, 0x7b, 0xbf, 0xf9, 0xe5, 0x8b, 0x3f, 0xd6, 0x7e, 0xf4, 0x00, 0xbb,
	0x25, 0x9f, 0xcc, 0xea, 0xf6, 0xcf, 0x17, 0x45, 0x01, 0xa1, 0xa9, 0x9d, 0xef, 0x22, 0xf1, 0xdf,
	0x5b, 0xc3, 0xb8, 0x93, 0x78, 0x0d, 0x87, 0x55, 0x49, 0xa9, 0x96, 0x59, 0xb2, 0xf6, 0x77, 0x1b,
	0x99, 0xd5, 0x24, 0x06, 0x9c, 0x99, 0x7e, 0x19, 0x77, 0xe4, 0x5e, 0x5a, 0x8c, 0x08, 0x7d, 0xa3,
	0x54, 0x25, 0x42, 0x6f, 0x07, 0x31, 0x9f, 0xf1, 0x08, 0x42, 0x5d, 0x25, 0x99, 0x6e, 0x48, 0x44,
	0x1e, 0x47, 0xf1, 0x40, 0x57, 0x4b, 0xaf, 0xf0, 0x02, 0xc6, 0x69, 0x6d, 0x2d, 0x15, 0xe9, 0x5b,
	0x52, 0x4a, 0x6d, 0xc5, 0x90, 0x4b, 0x47, 0x1d, 0x5c, 0x79, 0x86, 0xe7, 0x30, 0xaa, 0x9c, 0xdc,
	0x52, 0xb2, 0x91, 0xa9, 0x33, 0x56, 0x00, 0x37, 0x1f, 0x30, 0x7b, 0x60, 0xf4, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x69, 0xad, 0x2c, 0x92, 0x4a, 0x02, 0x00, 0x00,
}
