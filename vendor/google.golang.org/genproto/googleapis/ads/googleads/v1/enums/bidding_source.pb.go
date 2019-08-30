// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/bidding_source.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Indicates where a bid or target is defined. For example, an ad group
// criterion may define a cpc bid directly, or it can inherit its cpc bid from
// the ad group.
type BiddingSourceEnum_BiddingSource int32

const (
	// Not specified.
	BiddingSourceEnum_UNSPECIFIED BiddingSourceEnum_BiddingSource = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingSourceEnum_UNKNOWN BiddingSourceEnum_BiddingSource = 1
	// Effective bid or target is inherited from campaign bidding strategy.
	BiddingSourceEnum_CAMPAIGN_BIDDING_STRATEGY BiddingSourceEnum_BiddingSource = 5
	// The bid or target is defined on the ad group.
	BiddingSourceEnum_AD_GROUP BiddingSourceEnum_BiddingSource = 6
	// The bid or target is defined on the ad group criterion.
	BiddingSourceEnum_AD_GROUP_CRITERION BiddingSourceEnum_BiddingSource = 7
)

var BiddingSourceEnum_BiddingSource_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "CAMPAIGN_BIDDING_STRATEGY",
	6: "AD_GROUP",
	7: "AD_GROUP_CRITERION",
}
var BiddingSourceEnum_BiddingSource_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"CAMPAIGN_BIDDING_STRATEGY": 5,
	"AD_GROUP":                  6,
	"AD_GROUP_CRITERION":        7,
}

func (x BiddingSourceEnum_BiddingSource) String() string {
	return proto.EnumName(BiddingSourceEnum_BiddingSource_name, int32(x))
}
func (BiddingSourceEnum_BiddingSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bidding_source_cb63232b9a57dcf9, []int{0, 0}
}

// Container for enum describing possible bidding sources.
type BiddingSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingSourceEnum) Reset()         { *m = BiddingSourceEnum{} }
func (m *BiddingSourceEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingSourceEnum) ProtoMessage()    {}
func (*BiddingSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_source_cb63232b9a57dcf9, []int{0}
}
func (m *BiddingSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingSourceEnum.Unmarshal(m, b)
}
func (m *BiddingSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingSourceEnum.Marshal(b, m, deterministic)
}
func (dst *BiddingSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingSourceEnum.Merge(dst, src)
}
func (m *BiddingSourceEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingSourceEnum.Size(m)
}
func (m *BiddingSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BiddingSourceEnum)(nil), "google.ads.googleads.v1.enums.BiddingSourceEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.BiddingSourceEnum_BiddingSource", BiddingSourceEnum_BiddingSource_name, BiddingSourceEnum_BiddingSource_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/bidding_source.proto", fileDescriptor_bidding_source_cb63232b9a57dcf9)
}

var fileDescriptor_bidding_source_cb63232b9a57dcf9 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xfb, 0x30,
	0x18, 0xfd, 0x75, 0x3f, 0xdc, 0x24, 0x53, 0xac, 0xb9, 0x10, 0x14, 0x77, 0xb1, 0x3d, 0x40, 0x4a,
	0xf5, 0x2e, 0x5e, 0xa5, 0x6b, 0x2d, 0x45, 0xcc, 0x4a, 0xf7, 0x47, 0x94, 0x42, 0xe9, 0x96, 0x12,
	0x0a, 0x5b, 0x32, 0x9a, 0x6d, 0xaf, 0xe0, 0x7b, 0x78, 0xe9, 0xa3, 0xf8, 0x22, 0x82, 0x4f, 0x21,
	0x4d, 0xec, 0x60, 0x17, 0x7a, 0x13, 0x4e, 0xbe, 0xef, 0x9c, 0xc3, 0x77, 0x0e, 0xb8, 0xe1, 0x52,
	0xf2, 0x65, 0xe1, 0xe4, 0x4c, 0x39, 0x06, 0xd6, 0x68, 0xe7, 0x3a, 0x85, 0xd8, 0xae, 0x94, 0x33,
	0x2f, 0x19, 0x2b, 0x05, 0xcf, 0x94, 0xdc, 0x56, 0x8b, 0x02, 0xad, 0x2b, 0xb9, 0x91, 0xb0, 0x67,
	0x88, 0x28, 0x67, 0x0a, 0xed, 0x35, 0x68, 0xe7, 0x22, 0xad, 0xb9, 0xba, 0x6e, 0x2c, 0xd7, 0xa5,
	0x93, 0x0b, 0x21, 0x37, 0xf9, 0xa6, 0x94, 0x42, 0x19, 0xf1, 0xe0, 0xd5, 0x02, 0xe7, 0x9e, 0x71,
	0x1d, 0x6b, 0xd3, 0x40, 0x6c, 0x57, 0x83, 0x0a, 0x9c, 0x1e, 0x0c, 0xe1, 0x19, 0xe8, 0x4e, 0xe9,
	0x38, 0x0e, 0x86, 0xd1, 0x7d, 0x14, 0xf8, 0xf6, 0x3f, 0xd8, 0x05, 0x9d, 0x29, 0x7d, 0xa0, 0xa3,
	0x27, 0x6a, 0x5b, 0xb0, 0x07, 0x2e, 0x87, 0xe4, 0x31, 0x26, 0x51, 0x48, 0x33, 0x2f, 0xf2, 0xfd,
	0x88, 0x86, 0xd9, 0x78, 0x92, 0x90, 0x49, 0x10, 0x3e, 0xdb, 0x47, 0xf0, 0x04, 0x1c, 0x13, 0x3f,
	0x0b, 0x93, 0xd1, 0x34, 0xb6, 0xdb, 0xf0, 0x02, 0xc0, 0xe6, 0x97, 0x0d, 0x93, 0x68, 0x12, 0x24,
	0xd1, 0x88, 0xda, 0x1d, 0xef, 0xd3, 0x02, 0xfd, 0x85, 0x5c, 0xa1, 0x3f, 0xd3, 0x78, 0xf0, 0xe0,
	0xae, 0xb8, 0xce, 0x10, 0x5b, 0x2f, 0xde, 0x8f, 0x88, 0xcb, 0x65, 0x2e, 0x38, 0x92, 0x15, 0x77,
	0x78, 0x21, 0x74, 0xc2, 0xa6, 0xc6, 0x75, 0xa9, 0x7e, 0x69, 0xf5, 0x4e, 0xbf, 0x6f, 0xad, 0xff,
	0x21, 0x21, 0xef, 0xad, 0x5e, 0x68, 0xac, 0x08, 0x53, 0xc8, 0xc0, 0x1a, 0xcd, 0x5c, 0x54, 0x17,
	0xa3, 0x3e, 0x9a, 0x7d, 0x4a, 0x98, 0x4a, 0xf7, 0xfb, 0x74, 0xe6, 0xa6, 0x7a, 0xff, 0xd5, 0xea,
	0x9b, 0x21, 0xc6, 0x84, 0x29, 0x8c, 0xf7, 0x0c, 0x8c, 0x67, 0x2e, 0xc6, 0x9a, 0x33, 0x6f, 0xeb,
	0xc3, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xc9, 0xa8, 0x91, 0xed, 0x01, 0x00, 0x00,
}
