// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: rpc.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The ErrCode enum defines errors for gRPC API functions. These are converted
// from the Go error types returned by gnoclient.
type ErrCode int32

const (
	ErrCode_Undefined               ErrCode = 0 // default value, should never be set manually
	ErrCode_TODO                    ErrCode = 1 // indicates that you plan to create an error later
	ErrCode_ErrNotImplemented       ErrCode = 2 // indicates that a method is not implemented yet
	ErrCode_ErrInternal             ErrCode = 3 // indicates an unknown error (without Code), i.e. in gRPC
	ErrCode_ErrInvalidInput         ErrCode = 100
	ErrCode_ErrBridgeInterrupted    ErrCode = 101
	ErrCode_ErrMissingInput         ErrCode = 102
	ErrCode_ErrSerialization        ErrCode = 103
	ErrCode_ErrDeserialization      ErrCode = 104
	ErrCode_ErrCryptoKeyTypeUnknown ErrCode = 105
	ErrCode_ErrCryptoKeyNotFound    ErrCode = 106
	ErrCode_ErrNoActiveAccount      ErrCode = 107
	ErrCode_ErrRunGRPCServer        ErrCode = 108
	ErrCode_ErrDecryptionFailed     ErrCode = 109
	ErrCode_ErrUnknownAddress       ErrCode = 110 // indicates that the address is unknown on the blockchain
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:   "Undefined",
		1:   "TODO",
		2:   "ErrNotImplemented",
		3:   "ErrInternal",
		100: "ErrInvalidInput",
		101: "ErrBridgeInterrupted",
		102: "ErrMissingInput",
		103: "ErrSerialization",
		104: "ErrDeserialization",
		105: "ErrCryptoKeyTypeUnknown",
		106: "ErrCryptoKeyNotFound",
		107: "ErrNoActiveAccount",
		108: "ErrRunGRPCServer",
		109: "ErrDecryptionFailed",
		110: "ErrUnknownAddress",
	}
	ErrCode_value = map[string]int32{
		"Undefined":               0,
		"TODO":                    1,
		"ErrNotImplemented":       2,
		"ErrInternal":             3,
		"ErrInvalidInput":         100,
		"ErrBridgeInterrupted":    101,
		"ErrMissingInput":         102,
		"ErrSerialization":        103,
		"ErrDeserialization":      104,
		"ErrCryptoKeyTypeUnknown": 105,
		"ErrCryptoKeyNotFound":    106,
		"ErrNoActiveAccount":      107,
		"ErrRunGRPCServer":        108,
		"ErrDecryptionFailed":     109,
		"ErrUnknownAddress":       110,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

type ErrDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []ErrCode `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=land.gno.gnomobile.v1.ErrCode" json:"codes,omitempty"`
}

func (x *ErrDetails) Reset() {
	*x = ErrDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrDetails) ProtoMessage() {}

func (x *ErrDetails) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrDetails.ProtoReflect.Descriptor instead.
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ErrDetails) GetCodes() []ErrCode {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x14, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f,
	0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2a, 0xd1, 0x02, 0x0a,
	0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x44, 0x4f, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x72, 0x72,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x10, 0x64, 0x12, 0x18,
	0x0a, 0x14, 0x45, 0x72, 0x72, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x72, 0x75, 0x70, 0x74, 0x65, 0x64, 0x10, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x10, 0x66, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x72, 0x72, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x67, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x68, 0x12, 0x1b, 0x0a, 0x17, 0x45,
	0x72, 0x72, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x69, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x6a, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x6b, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x72,
	0x72, 0x52, 0x75, 0x6e, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x6c,
	0x12, 0x17, 0x0a, 0x13, 0x45, 0x72, 0x72, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x6d, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x72, 0x72,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x6e,
	0x32, 0xa5, 0x14, 0x0a, 0x10, 0x47, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e,
	0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x61,
	0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67,
	0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x12, 0x34, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67,
	0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x68, 0x72, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6c, 0x61, 0x6e, 0x64,
	0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x64, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x29, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x4b, 0x65, 0x79,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e,
	0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x61, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e,
	0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x4b, 0x65,
	0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x70, 0x0a, 0x0f, 0x48, 0x61, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2d, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e,
	0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x4b, 0x65,
	0x79, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x4b, 0x65, 0x79,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x48, 0x61, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x4f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x2e, 0x6c, 0x61,
	0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x4f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x4b, 0x65, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x31, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e,
	0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f,
	0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67,
	0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x4f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x6c, 0x61,
	0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e,
	0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67,
	0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e,
	0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x29, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c,
	0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x6c,
	0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6c,
	0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x2e,
	0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x61, 0x6e, 0x64,
	0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67,
	0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e,
	0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x52, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x6c, 0x61,
	0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x24, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e,
	0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x05, 0x51, 0x45, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e,
	0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x61,
	0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x6c, 0x61, 0x6e, 0x64,
	0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x70, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x42,
	0x65, 0x63, 0x68, 0x33, 0x32, 0x12, 0x2d, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f,
	0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x42, 0x65, 0x63, 0x68, 0x33, 0x32, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e,
	0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x42, 0x65, 0x63, 0x68, 0x33, 0x32, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46,
	0x72, 0x6f, 0x6d, 0x42, 0x65, 0x63, 0x68, 0x33, 0x32, 0x12, 0x2f, 0x2e, 0x6c, 0x61, 0x6e, 0x64,
	0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x65, 0x63,
	0x68, 0x33, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x65,
	0x63, 0x68, 0x33, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x23, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f,
	0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x66, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x29, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x2e, 0x67, 0x6e, 0x6f, 0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x30, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6e, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x03, 0x52, 0x54, 0x47, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_proto_goTypes = []interface{}{
	(ErrCode)(0),                              // 0: land.gno.gnomobile.v1.ErrCode
	(*ErrDetails)(nil),                        // 1: land.gno.gnomobile.v1.ErrDetails
	(*SetRemoteRequest)(nil),                  // 2: land.gno.gnomobile.v1.SetRemoteRequest
	(*SetChainIDRequest)(nil),                 // 3: land.gno.gnomobile.v1.SetChainIDRequest
	(*GenerateRecoveryPhraseRequest)(nil),     // 4: land.gno.gnomobile.v1.GenerateRecoveryPhraseRequest
	(*ListKeyInfoRequest)(nil),                // 5: land.gno.gnomobile.v1.ListKeyInfoRequest
	(*HasKeyByNameRequest)(nil),               // 6: land.gno.gnomobile.v1.HasKeyByNameRequest
	(*HasKeyByAddressRequest)(nil),            // 7: land.gno.gnomobile.v1.HasKeyByAddressRequest
	(*HasKeyByNameOrAddressRequest)(nil),      // 8: land.gno.gnomobile.v1.HasKeyByNameOrAddressRequest
	(*GetKeyInfoByNameRequest)(nil),           // 9: land.gno.gnomobile.v1.GetKeyInfoByNameRequest
	(*GetKeyInfoByAddressRequest)(nil),        // 10: land.gno.gnomobile.v1.GetKeyInfoByAddressRequest
	(*GetKeyInfoByNameOrAddressRequest)(nil),  // 11: land.gno.gnomobile.v1.GetKeyInfoByNameOrAddressRequest
	(*CreateAccountRequest)(nil),              // 12: land.gno.gnomobile.v1.CreateAccountRequest
	(*SelectAccountRequest)(nil),              // 13: land.gno.gnomobile.v1.SelectAccountRequest
	(*SetPasswordRequest)(nil),                // 14: land.gno.gnomobile.v1.SetPasswordRequest
	(*GetActiveAccountRequest)(nil),           // 15: land.gno.gnomobile.v1.GetActiveAccountRequest
	(*QueryAccountRequest)(nil),               // 16: land.gno.gnomobile.v1.QueryAccountRequest
	(*DeleteAccountRequest)(nil),              // 17: land.gno.gnomobile.v1.DeleteAccountRequest
	(*QueryRequest)(nil),                      // 18: land.gno.gnomobile.v1.QueryRequest
	(*RenderRequest)(nil),                     // 19: land.gno.gnomobile.v1.RenderRequest
	(*QEvalRequest)(nil),                      // 20: land.gno.gnomobile.v1.QEvalRequest
	(*CallRequest)(nil),                       // 21: land.gno.gnomobile.v1.CallRequest
	(*AddressToBech32Request)(nil),            // 22: land.gno.gnomobile.v1.AddressToBech32Request
	(*AddressFromBech32Request)(nil),          // 23: land.gno.gnomobile.v1.AddressFromBech32Request
	(*HelloRequest)(nil),                      // 24: land.gno.gnomobile.v1.HelloRequest
	(*HelloStreamRequest)(nil),                // 25: land.gno.gnomobile.v1.HelloStreamRequest
	(*SetRemoteResponse)(nil),                 // 26: land.gno.gnomobile.v1.SetRemoteResponse
	(*SetChainIDResponse)(nil),                // 27: land.gno.gnomobile.v1.SetChainIDResponse
	(*GenerateRecoveryPhraseResponse)(nil),    // 28: land.gno.gnomobile.v1.GenerateRecoveryPhraseResponse
	(*ListKeyInfoResponse)(nil),               // 29: land.gno.gnomobile.v1.ListKeyInfoResponse
	(*HasKeyByNameResponse)(nil),              // 30: land.gno.gnomobile.v1.HasKeyByNameResponse
	(*HasKeyByAddressResponse)(nil),           // 31: land.gno.gnomobile.v1.HasKeyByAddressResponse
	(*HasKeyByNameOrAddressResponse)(nil),     // 32: land.gno.gnomobile.v1.HasKeyByNameOrAddressResponse
	(*GetKeyInfoByNameResponse)(nil),          // 33: land.gno.gnomobile.v1.GetKeyInfoByNameResponse
	(*GetKeyInfoByAddressResponse)(nil),       // 34: land.gno.gnomobile.v1.GetKeyInfoByAddressResponse
	(*GetKeyInfoByNameOrAddressResponse)(nil), // 35: land.gno.gnomobile.v1.GetKeyInfoByNameOrAddressResponse
	(*CreateAccountResponse)(nil),             // 36: land.gno.gnomobile.v1.CreateAccountResponse
	(*SelectAccountResponse)(nil),             // 37: land.gno.gnomobile.v1.SelectAccountResponse
	(*SetPasswordResponse)(nil),               // 38: land.gno.gnomobile.v1.SetPasswordResponse
	(*GetActiveAccountResponse)(nil),          // 39: land.gno.gnomobile.v1.GetActiveAccountResponse
	(*QueryAccountResponse)(nil),              // 40: land.gno.gnomobile.v1.QueryAccountResponse
	(*DeleteAccountResponse)(nil),             // 41: land.gno.gnomobile.v1.DeleteAccountResponse
	(*QueryResponse)(nil),                     // 42: land.gno.gnomobile.v1.QueryResponse
	(*RenderResponse)(nil),                    // 43: land.gno.gnomobile.v1.RenderResponse
	(*QEvalResponse)(nil),                     // 44: land.gno.gnomobile.v1.QEvalResponse
	(*CallResponse)(nil),                      // 45: land.gno.gnomobile.v1.CallResponse
	(*AddressToBech32Response)(nil),           // 46: land.gno.gnomobile.v1.AddressToBech32Response
	(*AddressFromBech32Response)(nil),         // 47: land.gno.gnomobile.v1.AddressFromBech32Response
	(*HelloResponse)(nil),                     // 48: land.gno.gnomobile.v1.HelloResponse
	(*HelloStreamResponse)(nil),               // 49: land.gno.gnomobile.v1.HelloStreamResponse
}
var file_rpc_proto_depIdxs = []int32{
	0,  // 0: land.gno.gnomobile.v1.ErrDetails.codes:type_name -> land.gno.gnomobile.v1.ErrCode
	2,  // 1: land.gno.gnomobile.v1.GnomobileService.SetRemote:input_type -> land.gno.gnomobile.v1.SetRemoteRequest
	3,  // 2: land.gno.gnomobile.v1.GnomobileService.SetChainID:input_type -> land.gno.gnomobile.v1.SetChainIDRequest
	4,  // 3: land.gno.gnomobile.v1.GnomobileService.GenerateRecoveryPhrase:input_type -> land.gno.gnomobile.v1.GenerateRecoveryPhraseRequest
	5,  // 4: land.gno.gnomobile.v1.GnomobileService.ListKeyInfo:input_type -> land.gno.gnomobile.v1.ListKeyInfoRequest
	6,  // 5: land.gno.gnomobile.v1.GnomobileService.HasKeyByName:input_type -> land.gno.gnomobile.v1.HasKeyByNameRequest
	7,  // 6: land.gno.gnomobile.v1.GnomobileService.HasKeyByAddress:input_type -> land.gno.gnomobile.v1.HasKeyByAddressRequest
	8,  // 7: land.gno.gnomobile.v1.GnomobileService.HasKeyByNameOrAddress:input_type -> land.gno.gnomobile.v1.HasKeyByNameOrAddressRequest
	9,  // 8: land.gno.gnomobile.v1.GnomobileService.GetKeyInfoByName:input_type -> land.gno.gnomobile.v1.GetKeyInfoByNameRequest
	10, // 9: land.gno.gnomobile.v1.GnomobileService.GetKeyInfoByAddress:input_type -> land.gno.gnomobile.v1.GetKeyInfoByAddressRequest
	11, // 10: land.gno.gnomobile.v1.GnomobileService.GetKeyInfoByNameOrAddress:input_type -> land.gno.gnomobile.v1.GetKeyInfoByNameOrAddressRequest
	12, // 11: land.gno.gnomobile.v1.GnomobileService.CreateAccount:input_type -> land.gno.gnomobile.v1.CreateAccountRequest
	13, // 12: land.gno.gnomobile.v1.GnomobileService.SelectAccount:input_type -> land.gno.gnomobile.v1.SelectAccountRequest
	14, // 13: land.gno.gnomobile.v1.GnomobileService.SetPassword:input_type -> land.gno.gnomobile.v1.SetPasswordRequest
	15, // 14: land.gno.gnomobile.v1.GnomobileService.GetActiveAccount:input_type -> land.gno.gnomobile.v1.GetActiveAccountRequest
	16, // 15: land.gno.gnomobile.v1.GnomobileService.QueryAccount:input_type -> land.gno.gnomobile.v1.QueryAccountRequest
	17, // 16: land.gno.gnomobile.v1.GnomobileService.DeleteAccount:input_type -> land.gno.gnomobile.v1.DeleteAccountRequest
	18, // 17: land.gno.gnomobile.v1.GnomobileService.Query:input_type -> land.gno.gnomobile.v1.QueryRequest
	19, // 18: land.gno.gnomobile.v1.GnomobileService.Render:input_type -> land.gno.gnomobile.v1.RenderRequest
	20, // 19: land.gno.gnomobile.v1.GnomobileService.QEval:input_type -> land.gno.gnomobile.v1.QEvalRequest
	21, // 20: land.gno.gnomobile.v1.GnomobileService.Call:input_type -> land.gno.gnomobile.v1.CallRequest
	22, // 21: land.gno.gnomobile.v1.GnomobileService.AddressToBech32:input_type -> land.gno.gnomobile.v1.AddressToBech32Request
	23, // 22: land.gno.gnomobile.v1.GnomobileService.AddressFromBech32:input_type -> land.gno.gnomobile.v1.AddressFromBech32Request
	24, // 23: land.gno.gnomobile.v1.GnomobileService.Hello:input_type -> land.gno.gnomobile.v1.HelloRequest
	25, // 24: land.gno.gnomobile.v1.GnomobileService.HelloStream:input_type -> land.gno.gnomobile.v1.HelloStreamRequest
	26, // 25: land.gno.gnomobile.v1.GnomobileService.SetRemote:output_type -> land.gno.gnomobile.v1.SetRemoteResponse
	27, // 26: land.gno.gnomobile.v1.GnomobileService.SetChainID:output_type -> land.gno.gnomobile.v1.SetChainIDResponse
	28, // 27: land.gno.gnomobile.v1.GnomobileService.GenerateRecoveryPhrase:output_type -> land.gno.gnomobile.v1.GenerateRecoveryPhraseResponse
	29, // 28: land.gno.gnomobile.v1.GnomobileService.ListKeyInfo:output_type -> land.gno.gnomobile.v1.ListKeyInfoResponse
	30, // 29: land.gno.gnomobile.v1.GnomobileService.HasKeyByName:output_type -> land.gno.gnomobile.v1.HasKeyByNameResponse
	31, // 30: land.gno.gnomobile.v1.GnomobileService.HasKeyByAddress:output_type -> land.gno.gnomobile.v1.HasKeyByAddressResponse
	32, // 31: land.gno.gnomobile.v1.GnomobileService.HasKeyByNameOrAddress:output_type -> land.gno.gnomobile.v1.HasKeyByNameOrAddressResponse
	33, // 32: land.gno.gnomobile.v1.GnomobileService.GetKeyInfoByName:output_type -> land.gno.gnomobile.v1.GetKeyInfoByNameResponse
	34, // 33: land.gno.gnomobile.v1.GnomobileService.GetKeyInfoByAddress:output_type -> land.gno.gnomobile.v1.GetKeyInfoByAddressResponse
	35, // 34: land.gno.gnomobile.v1.GnomobileService.GetKeyInfoByNameOrAddress:output_type -> land.gno.gnomobile.v1.GetKeyInfoByNameOrAddressResponse
	36, // 35: land.gno.gnomobile.v1.GnomobileService.CreateAccount:output_type -> land.gno.gnomobile.v1.CreateAccountResponse
	37, // 36: land.gno.gnomobile.v1.GnomobileService.SelectAccount:output_type -> land.gno.gnomobile.v1.SelectAccountResponse
	38, // 37: land.gno.gnomobile.v1.GnomobileService.SetPassword:output_type -> land.gno.gnomobile.v1.SetPasswordResponse
	39, // 38: land.gno.gnomobile.v1.GnomobileService.GetActiveAccount:output_type -> land.gno.gnomobile.v1.GetActiveAccountResponse
	40, // 39: land.gno.gnomobile.v1.GnomobileService.QueryAccount:output_type -> land.gno.gnomobile.v1.QueryAccountResponse
	41, // 40: land.gno.gnomobile.v1.GnomobileService.DeleteAccount:output_type -> land.gno.gnomobile.v1.DeleteAccountResponse
	42, // 41: land.gno.gnomobile.v1.GnomobileService.Query:output_type -> land.gno.gnomobile.v1.QueryResponse
	43, // 42: land.gno.gnomobile.v1.GnomobileService.Render:output_type -> land.gno.gnomobile.v1.RenderResponse
	44, // 43: land.gno.gnomobile.v1.GnomobileService.QEval:output_type -> land.gno.gnomobile.v1.QEvalResponse
	45, // 44: land.gno.gnomobile.v1.GnomobileService.Call:output_type -> land.gno.gnomobile.v1.CallResponse
	46, // 45: land.gno.gnomobile.v1.GnomobileService.AddressToBech32:output_type -> land.gno.gnomobile.v1.AddressToBech32Response
	47, // 46: land.gno.gnomobile.v1.GnomobileService.AddressFromBech32:output_type -> land.gno.gnomobile.v1.AddressFromBech32Response
	48, // 47: land.gno.gnomobile.v1.GnomobileService.Hello:output_type -> land.gno.gnomobile.v1.HelloResponse
	49, // 48: land.gno.gnomobile.v1.GnomobileService.HelloStream:output_type -> land.gno.gnomobile.v1.HelloStreamResponse
	25, // [25:49] is the sub-list for method output_type
	1,  // [1:25] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_gnomobiletypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		EnumInfos:         file_rpc_proto_enumTypes,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
