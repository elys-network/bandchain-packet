// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/packet.proto

package packet

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// ResolveStatus encodes the status of an oracle request.
type ResolveStatus int32

const (
	// Open - the request is not yet resolved.
	RESOLVE_STATUS_OPEN ResolveStatus = 0
	// Success - the request has been resolved successfully with no errors.
	RESOLVE_STATUS_SUCCESS ResolveStatus = 1
	// Failure - an error occured during the request's resolve call.
	RESOLVE_STATUS_FAILURE ResolveStatus = 2
	// Expired - the request does not get enough reports from validator within the
	// timeframe.
	RESOLVE_STATUS_EXPIRED ResolveStatus = 3
)

var ResolveStatus_name = map[int32]string{
	0: "RESOLVE_STATUS_OPEN_UNSPECIFIED",
	1: "RESOLVE_STATUS_SUCCESS",
	2: "RESOLVE_STATUS_FAILURE",
	3: "RESOLVE_STATUS_EXPIRED",
}

var ResolveStatus_value = map[string]int32{
	"RESOLVE_STATUS_OPEN_UNSPECIFIED": 0,
	"RESOLVE_STATUS_SUCCESS":          1,
	"RESOLVE_STATUS_FAILURE":          2,
	"RESOLVE_STATUS_EXPIRED":          3,
}

func (x ResolveStatus) String() string {
	return proto.EnumName(ResolveStatus_name, int32(x))
}

func (ResolveStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9094f190a55b1bcf, []int{0}
}

// OracleRequestPacketData encodes an oracle request sent from other blockchains
// to BandChain.
type OracleRequestPacketData struct {
	// ClientID is the unique identifier of this oracle request, as specified by
	// the client. This same unique ID will be sent back to the requester with the
	// oracle response.
	ClientID string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// OracleScriptID is the unique identifier of the oracle script to be
	// executed.
	OracleScriptID uint64 `protobuf:"varint,2,opt,name=oracle_script_id,json=oracleScriptId,proto3" json:"oracle_script_id,omitempty"`
	// Calldata is the OBI-encoded calldata bytes available for oracle executor to
	// read.
	Calldata []byte `protobuf:"bytes,3,opt,name=calldata,proto3" json:"calldata,omitempty"`
	// AskCount is the number of validators that are requested to respond to this
	// oracle request. Higher value means more security, at a higher gas cost.
	AskCount uint64 `protobuf:"varint,4,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty"`
	// MinCount is the minimum number of validators necessary for the request to
	// proceed to the execution phase. Higher value means more security, at the
	// cost of liveness.
	MinCount uint64 `protobuf:"varint,5,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	// FeeLimit is the maximum tokens that will be paid to all data source
	// providers.
	FeeLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=fee_limit,json=feeLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_limit"`
	// PrepareGas is amount of gas to pay to prepare raw requests
	PrepareGas uint64 `protobuf:"varint,7,opt,name=prepare_gas,json=prepareGas,proto3" json:"prepare_gas,omitempty"`
	// ExecuteGas is amount of gas to reserve for executing
	ExecuteGas uint64 `protobuf:"varint,8,opt,name=execute_gas,json=executeGas,proto3" json:"execute_gas,omitempty"`
}

func (m *OracleRequestPacketData) Reset()         { *m = OracleRequestPacketData{} }
func (m *OracleRequestPacketData) String() string { return proto.CompactTextString(m) }
func (*OracleRequestPacketData) ProtoMessage()    {}
func (*OracleRequestPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9094f190a55b1bcf, []int{0}
}
func (m *OracleRequestPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleRequestPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleRequestPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleRequestPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleRequestPacketData.Merge(m, src)
}
func (m *OracleRequestPacketData) XXX_Size() int {
	return m.Size()
}
func (m *OracleRequestPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleRequestPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_OracleRequestPacketData proto.InternalMessageInfo

func (m *OracleRequestPacketData) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OracleRequestPacketData) GetOracleScriptID() uint64 {
	if m != nil {
		return m.OracleScriptID
	}
	return 0
}

func (m *OracleRequestPacketData) GetCalldata() []byte {
	if m != nil {
		return m.Calldata
	}
	return nil
}

func (m *OracleRequestPacketData) GetAskCount() uint64 {
	if m != nil {
		return m.AskCount
	}
	return 0
}

func (m *OracleRequestPacketData) GetMinCount() uint64 {
	if m != nil {
		return m.MinCount
	}
	return 0
}

func (m *OracleRequestPacketData) GetFeeLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeLimit
	}
	return nil
}

func (m *OracleRequestPacketData) GetPrepareGas() uint64 {
	if m != nil {
		return m.PrepareGas
	}
	return 0
}

func (m *OracleRequestPacketData) GetExecuteGas() uint64 {
	if m != nil {
		return m.ExecuteGas
	}
	return 0
}

// OracleRequestPacketAcknowledgement encodes an oracle request acknowledgement
// send back to requester chain.
type OracleRequestPacketAcknowledgement struct {
	// RequestID is BandChain's unique identifier for this oracle request.
	RequestID uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *OracleRequestPacketAcknowledgement) Reset()         { *m = OracleRequestPacketAcknowledgement{} }
func (m *OracleRequestPacketAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*OracleRequestPacketAcknowledgement) ProtoMessage()    {}
func (*OracleRequestPacketAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_9094f190a55b1bcf, []int{1}
}
func (m *OracleRequestPacketAcknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleRequestPacketAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleRequestPacketAcknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleRequestPacketAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleRequestPacketAcknowledgement.Merge(m, src)
}
func (m *OracleRequestPacketAcknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *OracleRequestPacketAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleRequestPacketAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_OracleRequestPacketAcknowledgement proto.InternalMessageInfo

func (m *OracleRequestPacketAcknowledgement) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

// OracleResponsePacketData encodes an oracle response from BandChain to the
// requester.
type OracleResponsePacketData struct {
	// ClientID is the unique identifier matched with that of the oracle request
	// packet.
	ClientID string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// RequestID is BandChain's unique identifier for this oracle request.
	RequestID uint64 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// AnsCount is the number of validators among to the asked validators that
	// actually responded to this oracle request prior to this oracle request
	// being resolved.
	AnsCount uint64 `protobuf:"varint,3,opt,name=ans_count,json=ansCount,proto3" json:"ans_count,omitempty"`
	// RequestTime is the UNIX epoch time at which the request was sent to
	// BandChain.
	RequestTime int64 `protobuf:"varint,4,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// ResolveTime is the UNIX epoch time at which the request was resolved to the
	// final result.
	ResolveTime int64 `protobuf:"varint,5,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
	// ResolveStatus is the status of this oracle request, which can be OK,
	// FAILURE, or EXPIRED.
	ResolveStatus ResolveStatus `protobuf:"varint,6,opt,name=resolve_status,json=resolveStatus,proto3,enum=oracle.v1.ResolveStatus" json:"resolve_status,omitempty"`
	// Result is the final aggregated value encoded in OBI format. Only available
	// if status if OK.
	Result []byte `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *OracleResponsePacketData) Reset()         { *m = OracleResponsePacketData{} }
func (m *OracleResponsePacketData) String() string { return proto.CompactTextString(m) }
func (*OracleResponsePacketData) ProtoMessage()    {}
func (*OracleResponsePacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9094f190a55b1bcf, []int{2}
}
func (m *OracleResponsePacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleResponsePacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleResponsePacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleResponsePacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleResponsePacketData.Merge(m, src)
}
func (m *OracleResponsePacketData) XXX_Size() int {
	return m.Size()
}
func (m *OracleResponsePacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleResponsePacketData.DiscardUnknown(m)
}

var xxx_messageInfo_OracleResponsePacketData proto.InternalMessageInfo

func (m *OracleResponsePacketData) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OracleResponsePacketData) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *OracleResponsePacketData) GetAnsCount() uint64 {
	if m != nil {
		return m.AnsCount
	}
	return 0
}

func (m *OracleResponsePacketData) GetRequestTime() int64 {
	if m != nil {
		return m.RequestTime
	}
	return 0
}

func (m *OracleResponsePacketData) GetResolveTime() int64 {
	if m != nil {
		return m.ResolveTime
	}
	return 0
}

func (m *OracleResponsePacketData) GetResolveStatus() ResolveStatus {
	if m != nil {
		return m.ResolveStatus
	}
	return RESOLVE_STATUS_OPEN
}

func (m *OracleResponsePacketData) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterEnum("oracle.v1.ResolveStatus", ResolveStatus_name, ResolveStatus_value)
	proto.RegisterType((*OracleRequestPacketData)(nil), "oracle.v1.OracleRequestPacketData")
	proto.RegisterType((*OracleRequestPacketAcknowledgement)(nil), "oracle.v1.OracleRequestPacketAcknowledgement")
	proto.RegisterType((*OracleResponsePacketData)(nil), "oracle.v1.OracleResponsePacketData")
}

func init() { proto.RegisterFile("oracle/v1/packet.proto", fileDescriptor_9094f190a55b1bcf) }

var fileDescriptor_9094f190a55b1bcf = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x3f, 0x4f, 0xdb, 0x4c,
	0x18, 0x8f, 0x93, 0x90, 0x37, 0x39, 0x42, 0x14, 0xf9, 0xad, 0xc0, 0xa4, 0x92, 0x9d, 0x32, 0xa5,
	0x55, 0xb1, 0x0b, 0x95, 0x3a, 0x54, 0x48, 0x15, 0x49, 0x0c, 0xb2, 0x84, 0x20, 0xb2, 0x49, 0x85,
	0xba, 0x58, 0x17, 0xfb, 0x08, 0x56, 0x6c, 0x9f, 0xeb, 0xbb, 0xd0, 0xf2, 0x0d, 0x2a, 0xa6, 0x7e,
	0x01, 0xa4, 0x4a, 0xdd, 0xba, 0xf7, 0x3b, 0x30, 0x32, 0x76, 0x4a, 0xab, 0xb0, 0xf4, 0x3b, 0x74,
	0xa9, 0x7c, 0x77, 0x41, 0x80, 0xd2, 0x0e, 0x9d, 0x92, 0xe7, 0xf7, 0xe7, 0x79, 0x2e, 0xcf, 0x2f,
	0x77, 0x60, 0x19, 0xa7, 0xd0, 0x0b, 0x91, 0x71, 0xba, 0x61, 0x24, 0xd0, 0x1b, 0x21, 0xaa, 0x27,
	0x29, 0xa6, 0x58, 0xae, 0x70, 0x5c, 0x3f, 0xdd, 0x68, 0x3c, 0x18, 0xe2, 0x21, 0x66, 0xa8, 0x91,
	0x7d, 0xe3, 0x82, 0x86, 0x36, 0xc4, 0x78, 0x18, 0x22, 0x83, 0x55, 0x83, 0xf1, 0xb1, 0x41, 0x83,
	0x08, 0x11, 0x0a, 0xa3, 0x44, 0x08, 0x56, 0xef, 0x0b, 0x60, 0x7c, 0x26, 0x28, 0xd5, 0xc3, 0x24,
	0xc2, 0xc4, 0x18, 0x40, 0x92, 0x4d, 0x1e, 0x20, 0x0a, 0x37, 0x0c, 0x0f, 0x07, 0x31, 0xe7, 0xd7,
	0xce, 0x0b, 0x60, 0xe5, 0x80, 0xcd, 0xb7, 0xd1, 0xdb, 0x31, 0x22, 0xb4, 0xc7, 0x8e, 0xd6, 0x85,
	0x14, 0xca, 0x8f, 0x41, 0xc5, 0x0b, 0x03, 0x14, 0x53, 0x37, 0xf0, 0x15, 0xa9, 0x29, 0xb5, 0x2a,
	0xed, 0xea, 0x74, 0xa2, 0x95, 0x3b, 0x0c, 0xb4, 0xba, 0x76, 0x99, 0xd3, 0x96, 0x2f, 0x6f, 0x81,
	0x3a, 0xff, 0x15, 0x2e, 0xf1, 0xd2, 0x20, 0x61, 0x8e, 0x7c, 0x53, 0x6a, 0x15, 0xdb, 0xf2, 0x74,
	0xa2, 0xd5, 0xf8, 0x04, 0x87, 0x51, 0x56, 0xd7, 0xae, 0xe1, 0xdb, 0xb5, 0x2f, 0x37, 0x40, 0xd9,
	0x83, 0x61, 0xe8, 0x43, 0x0a, 0x95, 0x42, 0x53, 0x6a, 0x55, 0xed, 0x9b, 0x5a, 0x7e, 0x08, 0x2a,
	0x90, 0x8c, 0x5c, 0x0f, 0x8f, 0x63, 0xaa, 0x14, 0xb3, 0x96, 0x76, 0x19, 0x92, 0x51, 0x27, 0xab,
	0x33, 0x32, 0x0a, 0x62, 0x41, 0x2e, 0x70, 0x32, 0x0a, 0x62, 0x4e, 0x9e, 0x80, 0xca, 0x31, 0x42,
	0x6e, 0x18, 0x44, 0x01, 0x55, 0x4a, 0xcd, 0x42, 0x6b, 0x71, 0x73, 0x55, 0xe7, 0xeb, 0xd0, 0xb3,
	0x75, 0xe8, 0x62, 0x1d, 0x7a, 0x07, 0x07, 0x71, 0xfb, 0xd9, 0xe5, 0x44, 0xcb, 0x7d, 0xf9, 0xae,
	0xb5, 0x86, 0x01, 0x3d, 0x19, 0x0f, 0x74, 0x0f, 0x47, 0x86, 0xd8, 0x1d, 0xff, 0x58, 0x27, 0xfe,
	0xc8, 0xa0, 0x67, 0x09, 0x22, 0xcc, 0x40, 0xec, 0xf2, 0x31, 0x42, 0x7b, 0x59, 0x73, 0x59, 0x03,
	0x8b, 0x49, 0x8a, 0x12, 0x98, 0x22, 0x77, 0x08, 0x89, 0xf2, 0x1f, 0x3b, 0x08, 0x10, 0xd0, 0x2e,
	0x24, 0x99, 0x00, 0xbd, 0x47, 0xde, 0x98, 0x72, 0x41, 0x99, 0x0b, 0x04, 0xb4, 0x0b, 0xc9, 0xcb,
	0xe2, 0xcf, 0x4f, 0x9a, 0xb4, 0x76, 0x04, 0xd6, 0xe6, 0x64, 0xb1, 0xed, 0x8d, 0x62, 0xfc, 0x2e,
	0x44, 0xfe, 0x10, 0x45, 0x28, 0xa6, 0xf2, 0x53, 0x00, 0x52, 0xce, 0xcf, 0x72, 0x29, 0xb6, 0x97,
	0xa6, 0x13, 0xad, 0x22, 0x5c, 0x56, 0xd7, 0xae, 0x08, 0x81, 0xe5, 0x8b, 0xce, 0x5f, 0xf3, 0x40,
	0x99, 0xb5, 0x26, 0x09, 0x8e, 0x09, 0xfa, 0xb7, 0x9c, 0xef, 0xce, 0xce, 0xff, 0x7d, 0x36, 0xcb,
	0x2e, 0x26, 0x22, 0x9e, 0x82, 0xc8, 0x2e, 0x26, 0x3c, 0x9e, 0x47, 0xa0, 0x3a, 0x6b, 0x95, 0xfd,
	0x9f, 0x59, 0xb6, 0x05, 0x7b, 0x51, 0x60, 0x87, 0x41, 0x84, 0xb8, 0x84, 0xe0, 0xf0, 0x14, 0x71,
	0xc9, 0xc2, 0x4c, 0xc2, 0x30, 0x26, 0x79, 0x05, 0x6a, 0x33, 0x09, 0xa1, 0x90, 0x8e, 0x89, 0x52,
	0x6a, 0x4a, 0xad, 0xda, 0xa6, 0xa2, 0xdf, 0xdc, 0x2a, 0xdd, 0xe6, 0x02, 0x87, 0xf1, 0xf6, 0x52,
	0x7a, 0xbb, 0x94, 0x97, 0x41, 0x29, 0x45, 0x64, 0x1c, 0x52, 0x16, 0x5b, 0xd5, 0x16, 0x15, 0xdf,
	0xdb, 0x93, 0x5f, 0x12, 0x58, 0xba, 0x63, 0x97, 0xb7, 0x80, 0x66, 0x9b, 0xce, 0xc1, 0xde, 0x6b,
	0xd3, 0x75, 0x0e, 0xb7, 0x0f, 0xfb, 0x8e, 0x7b, 0xd0, 0x33, 0xf7, 0xdd, 0xfe, 0xbe, 0xd3, 0x33,
	0x3b, 0xd6, 0x8e, 0x65, 0x76, 0xeb, 0xb9, 0xc6, 0xca, 0xf9, 0x45, 0xf3, 0xff, 0x39, 0x32, 0xf9,
	0x05, 0x58, 0xbe, 0x07, 0x3b, 0xfd, 0x4e, 0xc7, 0x74, 0x9c, 0xba, 0xd4, 0x68, 0x9c, 0x5f, 0x34,
	0xff, 0xc0, 0xce, 0xf1, 0xed, 0x6c, 0x5b, 0x7b, 0x7d, 0xdb, 0xac, 0xe7, 0xe7, 0xfa, 0x04, 0x3b,
	0xc7, 0x67, 0x1e, 0xf5, 0x2c, 0xdb, 0xec, 0xd6, 0x0b, 0x73, 0x7d, 0x82, 0x6d, 0x14, 0x3f, 0x7c,
	0x56, 0x73, 0x6d, 0xeb, 0x72, 0xaa, 0x4a, 0x57, 0x53, 0x55, 0xfa, 0x31, 0x55, 0xa5, 0x8f, 0xd7,
	0x6a, 0xee, 0xea, 0x5a, 0xcd, 0x7d, 0xbb, 0x56, 0x73, 0x6f, 0x8c, 0x5b, 0xb7, 0x64, 0x00, 0x63,
	0x9f, 0x3d, 0x26, 0x1e, 0x0e, 0x59, 0xe1, 0x9d, 0xc0, 0x20, 0x5e, 0xe7, 0x6f, 0x9c, 0x78, 0xea,
	0x06, 0x25, 0xa6, 0x78, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x76, 0xa9, 0x41, 0x39, 0x05, 0x05,
	0x00, 0x00,
}

func (this *OracleRequestPacketData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OracleRequestPacketData)
	if !ok {
		that2, ok := that.(OracleRequestPacketData)
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
	if this.ClientID != that1.ClientID {
		return false
	}
	if this.OracleScriptID != that1.OracleScriptID {
		return false
	}
	if !bytes.Equal(this.Calldata, that1.Calldata) {
		return false
	}
	if this.AskCount != that1.AskCount {
		return false
	}
	if this.MinCount != that1.MinCount {
		return false
	}
	if len(this.FeeLimit) != len(that1.FeeLimit) {
		return false
	}
	for i := range this.FeeLimit {
		if !this.FeeLimit[i].Equal(&that1.FeeLimit[i]) {
			return false
		}
	}
	if this.PrepareGas != that1.PrepareGas {
		return false
	}
	if this.ExecuteGas != that1.ExecuteGas {
		return false
	}
	return true
}
func (this *OracleRequestPacketAcknowledgement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OracleRequestPacketAcknowledgement)
	if !ok {
		that2, ok := that.(OracleRequestPacketAcknowledgement)
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
	if this.RequestID != that1.RequestID {
		return false
	}
	return true
}
func (this *OracleResponsePacketData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OracleResponsePacketData)
	if !ok {
		that2, ok := that.(OracleResponsePacketData)
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
	if this.ClientID != that1.ClientID {
		return false
	}
	if this.RequestID != that1.RequestID {
		return false
	}
	if this.AnsCount != that1.AnsCount {
		return false
	}
	if this.RequestTime != that1.RequestTime {
		return false
	}
	if this.ResolveTime != that1.ResolveTime {
		return false
	}
	if this.ResolveStatus != that1.ResolveStatus {
		return false
	}
	if !bytes.Equal(this.Result, that1.Result) {
		return false
	}
	return true
}
func (m *OracleRequestPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleRequestPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleRequestPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExecuteGas != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.ExecuteGas))
		i--
		dAtA[i] = 0x40
	}
	if m.PrepareGas != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.PrepareGas))
		i--
		dAtA[i] = 0x38
	}
	if len(m.FeeLimit) > 0 {
		for iNdEx := len(m.FeeLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MinCount != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.MinCount))
		i--
		dAtA[i] = 0x28
	}
	if m.AskCount != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.AskCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Calldata) > 0 {
		i -= len(m.Calldata)
		copy(dAtA[i:], m.Calldata)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Calldata)))
		i--
		dAtA[i] = 0x1a
	}
	if m.OracleScriptID != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.OracleScriptID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OracleRequestPacketAcknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleRequestPacketAcknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleRequestPacketAcknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestID != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OracleResponsePacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleResponsePacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleResponsePacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ResolveStatus != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.ResolveStatus))
		i--
		dAtA[i] = 0x30
	}
	if m.ResolveTime != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.ResolveTime))
		i--
		dAtA[i] = 0x28
	}
	if m.RequestTime != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.RequestTime))
		i--
		dAtA[i] = 0x20
	}
	if m.AnsCount != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.AnsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.RequestID != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleRequestPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.OracleScriptID != 0 {
		n += 1 + sovPacket(uint64(m.OracleScriptID))
	}
	l = len(m.Calldata)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.AskCount != 0 {
		n += 1 + sovPacket(uint64(m.AskCount))
	}
	if m.MinCount != 0 {
		n += 1 + sovPacket(uint64(m.MinCount))
	}
	if len(m.FeeLimit) > 0 {
		for _, e := range m.FeeLimit {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	if m.PrepareGas != 0 {
		n += 1 + sovPacket(uint64(m.PrepareGas))
	}
	if m.ExecuteGas != 0 {
		n += 1 + sovPacket(uint64(m.ExecuteGas))
	}
	return n
}

func (m *OracleRequestPacketAcknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestID != 0 {
		n += 1 + sovPacket(uint64(m.RequestID))
	}
	return n
}

func (m *OracleResponsePacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.RequestID != 0 {
		n += 1 + sovPacket(uint64(m.RequestID))
	}
	if m.AnsCount != 0 {
		n += 1 + sovPacket(uint64(m.AnsCount))
	}
	if m.RequestTime != 0 {
		n += 1 + sovPacket(uint64(m.RequestTime))
	}
	if m.ResolveTime != 0 {
		n += 1 + sovPacket(uint64(m.ResolveTime))
	}
	if m.ResolveStatus != 0 {
		n += 1 + sovPacket(uint64(m.ResolveStatus))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleRequestPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: OracleRequestPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleRequestPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScriptID", wireType)
			}
			m.OracleScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Calldata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Calldata = append(m.Calldata[:0], dAtA[iNdEx:postIndex]...)
			if m.Calldata == nil {
				m.Calldata = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
			}
			m.AskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
			}
			m.MinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeLimit = append(m.FeeLimit, types.Coin{})
			if err := m.FeeLimit[len(m.FeeLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGas", wireType)
			}
			m.PrepareGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGas", wireType)
			}
			m.ExecuteGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *OracleRequestPacketAcknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: OracleRequestPacketAcknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleRequestPacketAcknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *OracleResponsePacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: OracleResponsePacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleResponsePacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnsCount", wireType)
			}
			m.AnsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTime", wireType)
			}
			m.RequestTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveTime", wireType)
			}
			m.ResolveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveStatus", wireType)
			}
			m.ResolveStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveStatus |= ResolveStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
