// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	common.proto
	service.proto
	stream.proto

It has these top-level messages:
	Empty
	Transaction
	TxData
	Params
	Envelope
	StreamMessage
	Block
	ConnectionEstablish
	PeerTable
	Peer
	ConsensusMessage
	View
*/
package message

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

type Transaction_Status int32

const (
	Transaction_UNCONFIRMED Transaction_Status = 0
	Transaction_CONFIRMED   Transaction_Status = 1
	Transaction_UNKNOWN     Transaction_Status = 2
)

var Transaction_Status_name = map[int32]string{
	0: "UNCONFIRMED",
	1: "CONFIRMED",
	2: "UNKNOWN",
}
var Transaction_Status_value = map[string]int32{
	"UNCONFIRMED": 0,
	"CONFIRMED":   1,
	"UNKNOWN":     2,
}

func (x Transaction_Status) String() string {
	return proto.EnumName(Transaction_Status_name, int32(x))
}
func (Transaction_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type TxData_TxDataType int32

const (
	TxData_Invoke TxData_TxDataType = 0
	TxData_Query  TxData_TxDataType = 1
)

var TxData_TxDataType_name = map[int32]string{
	0: "Invoke",
	1: "Query",
}
var TxData_TxDataType_value = map[string]int32{
	"Invoke": 0,
	"Query":  1,
}

func (x TxData_TxDataType) String() string {
	return proto.EnumName(TxData_TxDataType_name, int32(x))
}
func (TxData_TxDataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Transaction struct {
	InvokePeerID      string             `protobuf:"bytes,1,opt,name=InvokePeerID" json:"InvokePeerID,omitempty"`
	TransactionID     string             `protobuf:"bytes,2,opt,name=TransactionID" json:"TransactionID,omitempty"`
	TransactionStatus Transaction_Status `protobuf:"varint,3,opt,name=TransactionStatus,enum=message.Transaction_Status" json:"TransactionStatus,omitempty"`
	TransactionHash   string             `protobuf:"bytes,4,opt,name=TransactionHash" json:"TransactionHash,omitempty"`
	TxData            *TxData            `protobuf:"bytes,5,opt,name=TxData" json:"TxData,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Transaction) GetInvokePeerID() string {
	if m != nil {
		return m.InvokePeerID
	}
	return ""
}

func (m *Transaction) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *Transaction) GetTransactionStatus() Transaction_Status {
	if m != nil {
		return m.TransactionStatus
	}
	return Transaction_UNCONFIRMED
}

func (m *Transaction) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *Transaction) GetTxData() *TxData {
	if m != nil {
		return m.TxData
	}
	return nil
}

type TxData struct {
	Jsonrpc    string            `protobuf:"bytes,1,opt,name=Jsonrpc" json:"Jsonrpc,omitempty"`
	Method     TxData_TxDataType `protobuf:"varint,2,opt,name=Method,enum=message.TxData_TxDataType" json:"Method,omitempty"`
	Params     *Params           `protobuf:"bytes,3,opt,name=Params" json:"Params,omitempty"`
	ContractID string            `protobuf:"bytes,4,opt,name=ContractID" json:"ContractID,omitempty"`
}

func (m *TxData) Reset()                    { *m = TxData{} }
func (m *TxData) String() string            { return proto.CompactTextString(m) }
func (*TxData) ProtoMessage()               {}
func (*TxData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TxData) GetJsonrpc() string {
	if m != nil {
		return m.Jsonrpc
	}
	return ""
}

func (m *TxData) GetMethod() TxData_TxDataType {
	if m != nil {
		return m.Method
	}
	return TxData_Invoke
}

func (m *TxData) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *TxData) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

type Params struct {
	ParamsType int32    `protobuf:"varint,1,opt,name=ParamsType" json:"ParamsType,omitempty"`
	Function   string   `protobuf:"bytes,2,opt,name=Function" json:"Function,omitempty"`
	Args       []string `protobuf:"bytes,3,rep,name=Args" json:"Args,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (m *Params) String() string            { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Params) GetParamsType() int32 {
	if m != nil {
		return m.ParamsType
	}
	return 0
}

func (m *Params) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Params) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "message.Empty")
	proto.RegisterType((*Transaction)(nil), "message.Transaction")
	proto.RegisterType((*TxData)(nil), "message.TxData")
	proto.RegisterType((*Params)(nil), "message.Params")
	proto.RegisterEnum("message.Transaction_Status", Transaction_Status_name, Transaction_Status_value)
	proto.RegisterEnum("message.TxData_TxDataType", TxData_TxDataType_name, TxData_TxDataType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x0e, 0xd2, 0x30,
	0x14, 0xc6, 0x19, 0xb0, 0xcd, 0x9d, 0xf1, 0x67, 0x9e, 0xab, 0x05, 0x13, 0x43, 0xa6, 0x89, 0xbb,
	0xda, 0xc5, 0x8c, 0x0f, 0x60, 0x18, 0xc4, 0x69, 0x18, 0x58, 0x21, 0x7a, 0x5b, 0x67, 0x03, 0xc6,
	0x6c, 0x5d, 0xba, 0x62, 0xe4, 0x9d, 0x7c, 0x0c, 0x1f, 0xcc, 0xd0, 0x15, 0x29, 0x78, 0xb5, 0x9e,
	0x5f, 0xbf, 0x9e, 0xf3, 0xed, 0x6b, 0x61, 0x54, 0xf2, 0xaa, 0xe2, 0x75, 0xd2, 0x08, 0x2e, 0x39,
	0xba, 0x15, 0x6b, 0x5b, 0x7a, 0x60, 0x91, 0x0b, 0xf6, 0xb2, 0x6a, 0xe4, 0x39, 0xfa, 0xdd, 0x07,
	0x7f, 0x27, 0x68, 0xdd, 0xd2, 0x52, 0x7e, 0xe7, 0x35, 0x46, 0x30, 0xca, 0xeb, 0x9f, 0xfc, 0x07,
	0xdb, 0x32, 0x26, 0xf2, 0x2c, 0xb4, 0xe6, 0x56, 0xec, 0x91, 0x3b, 0x86, 0x2f, 0x61, 0x6c, 0x1c,
	0xc9, 0xb3, 0xb0, 0xaf, 0x44, 0xf7, 0x10, 0x73, 0x78, 0x6a, 0x80, 0x4f, 0x92, 0xca, 0x53, 0x1b,
	0x0e, 0xe6, 0x56, 0x3c, 0x49, 0x9f, 0x25, 0xda, 0x47, 0x62, 0x28, 0x92, 0x4e, 0x42, 0xfe, 0x3f,
	0x85, 0x31, 0x4c, 0x0d, 0xf8, 0x8e, 0xb6, 0xc7, 0x70, 0xa8, 0x46, 0x3e, 0x62, 0x7c, 0x05, 0xce,
	0xee, 0x57, 0x46, 0x25, 0x0d, 0xed, 0xb9, 0x15, 0xfb, 0xe9, 0xf4, 0x36, 0x49, 0x61, 0xa2, 0xb7,
	0xa3, 0x37, 0xe0, 0xe8, 0xe6, 0x53, 0xf0, 0xf7, 0xc5, 0x62, 0x53, 0xac, 0x72, 0xb2, 0x5e, 0x66,
	0x41, 0x0f, 0xc7, 0xe0, 0xdd, 0x4a, 0x0b, 0x7d, 0x70, 0xf7, 0xc5, 0x87, 0x62, 0xf3, 0xb9, 0x08,
	0xfa, 0xd1, 0x1f, 0xeb, 0x3a, 0x00, 0x43, 0x70, 0xdf, 0xb7, 0xbc, 0x16, 0x4d, 0xa9, 0x43, 0xba,
	0x96, 0x98, 0x82, 0xb3, 0x66, 0xf2, 0xc8, 0xbf, 0xa9, 0x60, 0x26, 0xe9, 0xec, 0xc1, 0x84, 0xfe,
	0xec, 0xce, 0x0d, 0x23, 0x5a, 0x79, 0x31, 0xbe, 0xa5, 0x82, 0x56, 0x5d, 0x44, 0xa6, 0xf1, 0x0e,
	0x13, 0xbd, 0x8d, 0xcf, 0x01, 0x16, 0xbc, 0x96, 0x82, 0x96, 0x32, 0xcf, 0x74, 0x0c, 0x06, 0x89,
	0x5e, 0x00, 0xdc, 0xda, 0x23, 0x80, 0xd3, 0x5d, 0x5d, 0xd0, 0x43, 0x0f, 0xec, 0x8f, 0x27, 0x26,
	0xce, 0x81, 0x15, 0x7d, 0x01, 0xa3, 0x5d, 0xb7, 0xba, 0xc8, 0xd5, 0x8f, 0xd8, 0xc4, 0x20, 0x38,
	0x83, 0x27, 0xab, 0x53, 0xad, 0x02, 0xd6, 0xd7, 0xfc, 0xaf, 0x46, 0x84, 0xe1, 0x5b, 0x71, 0xb8,
	0x38, 0x1e, 0xc4, 0x1e, 0x51, 0xeb, 0xaf, 0x8e, 0x7a, 0x68, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x29, 0xfa, 0x11, 0x78, 0x02, 0x00, 0x00,
}