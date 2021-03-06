// Code generated by protoc-gen-go. DO NOT EDIT.
// comments/deprecated.proto is a deprecated file.

package comments

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

// Deprecated: Use DeprecatedEnum.Type.Values instead.
var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

// Deprecated: Use DeprecatedEnum.Type.Values instead.
var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (DeprecatedEnum) Type() protoreflect.EnumType {
	return xxx_File_comments_deprecated_proto_enumTypes[0]
}

func (x DeprecatedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeprecatedEnum.Type instead.
func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
type DeprecatedMessage struct {
	DeprecatedField      string   `protobuf:"bytes,1,opt,name=deprecated_field,json=deprecatedField,proto3" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *DeprecatedMessage) Reset() {
	*x = DeprecatedMessage{}
}

func (x *DeprecatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeprecatedMessage) ProtoMessage() {}

func (x *DeprecatedMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_deprecated_proto_messageTypes[0].MessageOf(x)
}

func (m *DeprecatedMessage) XXX_Methods() *protoiface.Methods {
	return xxx_File_comments_deprecated_proto_messageTypes[0].Methods()
}

// Deprecated: Use DeprecatedMessage.ProtoReflect.Type instead.
func (*DeprecatedMessage) Descriptor() ([]byte, []int) {
	return xxx_File_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *DeprecatedMessage) GetDeprecatedField() string {
	if x != nil {
		return x.DeprecatedField
	}
	return ""
}

var File_comments_deprecated_proto protoreflect.FileDescriptor

var xxx_File_comments_deprecated_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x2a, 0x28, 0x0a, 0x0e,
	0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12,
	0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x02,
	0x08, 0x01, 0x1a, 0x02, 0x18, 0x01, 0x42, 0x46, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0xb8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_comments_deprecated_proto_rawDesc_once sync.Once
	xxx_File_comments_deprecated_proto_rawDesc_data = xxx_File_comments_deprecated_proto_rawDesc
)

func xxx_File_comments_deprecated_proto_rawDescGZIP() []byte {
	xxx_File_comments_deprecated_proto_rawDesc_once.Do(func() {
		xxx_File_comments_deprecated_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_comments_deprecated_proto_rawDesc_data)
	})
	return xxx_File_comments_deprecated_proto_rawDesc_data
}

var xxx_File_comments_deprecated_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_comments_deprecated_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_comments_deprecated_proto_goTypes = []interface{}{
	(DeprecatedEnum)(0),       // 0: goproto.protoc.comments.DeprecatedEnum
	(*DeprecatedMessage)(nil), // 1: goproto.protoc.comments.DeprecatedMessage
}
var xxx_File_comments_deprecated_proto_depIdxs = []int32{}

func init() { xxx_File_comments_deprecated_proto_init() }
func xxx_File_comments_deprecated_proto_init() {
	if File_comments_deprecated_proto != nil {
		return
	}
	File_comments_deprecated_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_comments_deprecated_proto_rawDesc,
		GoTypes:            xxx_File_comments_deprecated_proto_goTypes,
		DependencyIndexes:  xxx_File_comments_deprecated_proto_depIdxs,
		EnumOutputTypes:    xxx_File_comments_deprecated_proto_enumTypes,
		MessageOutputTypes: xxx_File_comments_deprecated_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_comments_deprecated_proto_rawDesc = nil
	xxx_File_comments_deprecated_proto_goTypes = nil
	xxx_File_comments_deprecated_proto_depIdxs = nil
}
