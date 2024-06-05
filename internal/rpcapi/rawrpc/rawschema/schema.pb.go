// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.15.6
// source: terraform/schema.proto

package rawschema

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

type Schema_NestedBlock_NestingMode int32

const (
	Schema_NestedBlock_INVALID Schema_NestedBlock_NestingMode = 0
	Schema_NestedBlock_SINGLE  Schema_NestedBlock_NestingMode = 1
	Schema_NestedBlock_LIST    Schema_NestedBlock_NestingMode = 2
	Schema_NestedBlock_SET     Schema_NestedBlock_NestingMode = 3
	Schema_NestedBlock_MAP     Schema_NestedBlock_NestingMode = 4
	Schema_NestedBlock_GROUP   Schema_NestedBlock_NestingMode = 5
)

// Enum value maps for Schema_NestedBlock_NestingMode.
var (
	Schema_NestedBlock_NestingMode_name = map[int32]string{
		0: "INVALID",
		1: "SINGLE",
		2: "LIST",
		3: "SET",
		4: "MAP",
		5: "GROUP",
	}
	Schema_NestedBlock_NestingMode_value = map[string]int32{
		"INVALID": 0,
		"SINGLE":  1,
		"LIST":    2,
		"SET":     3,
		"MAP":     4,
		"GROUP":   5,
	}
)

func (x Schema_NestedBlock_NestingMode) Enum() *Schema_NestedBlock_NestingMode {
	p := new(Schema_NestedBlock_NestingMode)
	*p = x
	return p
}

func (x Schema_NestedBlock_NestingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schema_NestedBlock_NestingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_terraform_schema_proto_enumTypes[0].Descriptor()
}

func (Schema_NestedBlock_NestingMode) Type() protoreflect.EnumType {
	return &file_terraform_schema_proto_enumTypes[0]
}

func (x Schema_NestedBlock_NestingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schema_NestedBlock_NestingMode.Descriptor instead.
func (Schema_NestedBlock_NestingMode) EnumDescriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 2, 0}
}

type Schema_Object_NestingMode int32

const (
	Schema_Object_INVALID Schema_Object_NestingMode = 0
	Schema_Object_SINGLE  Schema_Object_NestingMode = 1
	Schema_Object_LIST    Schema_Object_NestingMode = 2
	Schema_Object_SET     Schema_Object_NestingMode = 3
	Schema_Object_MAP     Schema_Object_NestingMode = 4
)

// Enum value maps for Schema_Object_NestingMode.
var (
	Schema_Object_NestingMode_name = map[int32]string{
		0: "INVALID",
		1: "SINGLE",
		2: "LIST",
		3: "SET",
		4: "MAP",
	}
	Schema_Object_NestingMode_value = map[string]int32{
		"INVALID": 0,
		"SINGLE":  1,
		"LIST":    2,
		"SET":     3,
		"MAP":     4,
	}
)

func (x Schema_Object_NestingMode) Enum() *Schema_Object_NestingMode {
	p := new(Schema_Object_NestingMode)
	*p = x
	return p
}

func (x Schema_Object_NestingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schema_Object_NestingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_terraform_schema_proto_enumTypes[1].Descriptor()
}

func (Schema_Object_NestingMode) Type() protoreflect.EnumType {
	return &file_terraform_schema_proto_enumTypes[1]
}

func (x Schema_Object_NestingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schema_Object_NestingMode.Descriptor instead.
func (Schema_Object_NestingMode) EnumDescriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 3, 0}
}

type Schema_DocString_Format int32

const (
	Schema_DocString_PLAIN    Schema_DocString_Format = 0
	Schema_DocString_MARKDOWN Schema_DocString_Format = 1
)

// Enum value maps for Schema_DocString_Format.
var (
	Schema_DocString_Format_name = map[int32]string{
		0: "PLAIN",
		1: "MARKDOWN",
	}
	Schema_DocString_Format_value = map[string]int32{
		"PLAIN":    0,
		"MARKDOWN": 1,
	}
)

func (x Schema_DocString_Format) Enum() *Schema_DocString_Format {
	p := new(Schema_DocString_Format)
	*p = x
	return p
}

func (x Schema_DocString_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schema_DocString_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_terraform_schema_proto_enumTypes[2].Descriptor()
}

func (Schema_DocString_Format) Type() protoreflect.EnumType {
	return &file_terraform_schema_proto_enumTypes[2]
}

func (x Schema_DocString_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schema_DocString_Format.Descriptor instead.
func (Schema_DocString_Format) EnumDescriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 4, 0}
}

// Schema describes a schema for an instance of a particular object, such as
// a resource type or a provider's overall configuration.
type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Block is the top level configuration block for this schema.
	Block *Schema_Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetBlock() *Schema_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type Schema_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes  []*Schema_Attribute   `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	BlockTypes  []*Schema_NestedBlock `protobuf:"bytes,2,rep,name=block_types,json=blockTypes,proto3" json:"block_types,omitempty"`
	Description *Schema_DocString     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deprecated  bool                  `protobuf:"varint,4,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *Schema_Block) Reset() {
	*x = Schema_Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_Block) ProtoMessage() {}

func (x *Schema_Block) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_Block.ProtoReflect.Descriptor instead.
func (*Schema_Block) Descriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Schema_Block) GetAttributes() []*Schema_Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Schema_Block) GetBlockTypes() []*Schema_NestedBlock {
	if x != nil {
		return x.BlockTypes
	}
	return nil
}

func (x *Schema_Block) GetDescription() *Schema_DocString {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Schema_Block) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

type Schema_Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type        []byte            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	NestedType  *Schema_Object    `protobuf:"bytes,10,opt,name=nested_type,json=nestedType,proto3" json:"nested_type,omitempty"`
	Description *Schema_DocString `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Required    bool              `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Optional    bool              `protobuf:"varint,5,opt,name=optional,proto3" json:"optional,omitempty"`
	Computed    bool              `protobuf:"varint,6,opt,name=computed,proto3" json:"computed,omitempty"`
	Sensitive   bool              `protobuf:"varint,7,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
	Deprecated  bool              `protobuf:"varint,8,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *Schema_Attribute) Reset() {
	*x = Schema_Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_Attribute) ProtoMessage() {}

func (x *Schema_Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_Attribute.ProtoReflect.Descriptor instead.
func (*Schema_Attribute) Descriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Schema_Attribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Schema_Attribute) GetType() []byte {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Schema_Attribute) GetNestedType() *Schema_Object {
	if x != nil {
		return x.NestedType
	}
	return nil
}

func (x *Schema_Attribute) GetDescription() *Schema_DocString {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Schema_Attribute) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Schema_Attribute) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *Schema_Attribute) GetComputed() bool {
	if x != nil {
		return x.Computed
	}
	return false
}

func (x *Schema_Attribute) GetSensitive() bool {
	if x != nil {
		return x.Sensitive
	}
	return false
}

func (x *Schema_Attribute) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

type Schema_NestedBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeName string                         `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Block    *Schema_Block                  `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Nesting  Schema_NestedBlock_NestingMode `protobuf:"varint,3,opt,name=nesting,proto3,enum=terraform.schema.Schema_NestedBlock_NestingMode" json:"nesting,omitempty"`
}

func (x *Schema_NestedBlock) Reset() {
	*x = Schema_NestedBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_NestedBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_NestedBlock) ProtoMessage() {}

func (x *Schema_NestedBlock) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_NestedBlock.ProtoReflect.Descriptor instead.
func (*Schema_NestedBlock) Descriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Schema_NestedBlock) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Schema_NestedBlock) GetBlock() *Schema_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Schema_NestedBlock) GetNesting() Schema_NestedBlock_NestingMode {
	if x != nil {
		return x.Nesting
	}
	return Schema_NestedBlock_INVALID
}

type Schema_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*Schema_Attribute       `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Nesting    Schema_Object_NestingMode `protobuf:"varint,3,opt,name=nesting,proto3,enum=terraform.schema.Schema_Object_NestingMode" json:"nesting,omitempty"`
}

func (x *Schema_Object) Reset() {
	*x = Schema_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_Object) ProtoMessage() {}

func (x *Schema_Object) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_Object.ProtoReflect.Descriptor instead.
func (*Schema_Object) Descriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Schema_Object) GetAttributes() []*Schema_Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Schema_Object) GetNesting() Schema_Object_NestingMode {
	if x != nil {
		return x.Nesting
	}
	return Schema_Object_INVALID
}

type Schema_DocString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string                  `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Format      Schema_DocString_Format `protobuf:"varint,2,opt,name=format,proto3,enum=terraform.schema.Schema_DocString_Format" json:"format,omitempty"`
}

func (x *Schema_DocString) Reset() {
	*x = Schema_DocString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_DocString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_DocString) ProtoMessage() {}

func (x *Schema_DocString) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_DocString.ProtoReflect.Descriptor instead.
func (*Schema_DocString) Descriptor() ([]byte, []int) {
	return file_terraform_schema_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Schema_DocString) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schema_DocString) GetFormat() Schema_DocString_Format {
	if x != nil {
		return x.Format
	}
	return Schema_DocString_PLAIN
}

var File_terraform_schema_proto protoreflect.FileDescriptor

var file_terraform_schema_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xf7, 0x09, 0x0a, 0x06, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x34, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0xf8, 0x01, 0x0a, 0x05,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x44, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x44, 0x6f, 0x63, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x1a, 0xcd, 0x02, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x0a, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x44,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x6f,
	0x63, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x1a, 0xfb, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x4a, 0x0a, 0x07, 0x6e, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x6e, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4d, 0x0a, 0x0b, 0x4e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x03,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x50, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x05, 0x1a, 0xd7, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x42, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x6e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x07, 0x6e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x42, 0x0a, 0x0b, 0x4e, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x45, 0x54, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x50, 0x10, 0x04, 0x1a, 0x93,
	0x01, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41,
	0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x22, 0x21, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x4c, 0x41, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41, 0x52, 0x4b, 0x44, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_terraform_schema_proto_rawDescOnce sync.Once
	file_terraform_schema_proto_rawDescData = file_terraform_schema_proto_rawDesc
)

func file_terraform_schema_proto_rawDescGZIP() []byte {
	file_terraform_schema_proto_rawDescOnce.Do(func() {
		file_terraform_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_terraform_schema_proto_rawDescData)
	})
	return file_terraform_schema_proto_rawDescData
}

var file_terraform_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_terraform_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_terraform_schema_proto_goTypes = []interface{}{
	(Schema_NestedBlock_NestingMode)(0), // 0: terraform.schema.Schema.NestedBlock.NestingMode
	(Schema_Object_NestingMode)(0),      // 1: terraform.schema.Schema.Object.NestingMode
	(Schema_DocString_Format)(0),        // 2: terraform.schema.Schema.DocString.Format
	(*Schema)(nil),                      // 3: terraform.schema.Schema
	(*Schema_Block)(nil),                // 4: terraform.schema.Schema.Block
	(*Schema_Attribute)(nil),            // 5: terraform.schema.Schema.Attribute
	(*Schema_NestedBlock)(nil),          // 6: terraform.schema.Schema.NestedBlock
	(*Schema_Object)(nil),               // 7: terraform.schema.Schema.Object
	(*Schema_DocString)(nil),            // 8: terraform.schema.Schema.DocString
}
var file_terraform_schema_proto_depIdxs = []int32{
	4,  // 0: terraform.schema.Schema.block:type_name -> terraform.schema.Schema.Block
	5,  // 1: terraform.schema.Schema.Block.attributes:type_name -> terraform.schema.Schema.Attribute
	6,  // 2: terraform.schema.Schema.Block.block_types:type_name -> terraform.schema.Schema.NestedBlock
	8,  // 3: terraform.schema.Schema.Block.description:type_name -> terraform.schema.Schema.DocString
	7,  // 4: terraform.schema.Schema.Attribute.nested_type:type_name -> terraform.schema.Schema.Object
	8,  // 5: terraform.schema.Schema.Attribute.description:type_name -> terraform.schema.Schema.DocString
	4,  // 6: terraform.schema.Schema.NestedBlock.block:type_name -> terraform.schema.Schema.Block
	0,  // 7: terraform.schema.Schema.NestedBlock.nesting:type_name -> terraform.schema.Schema.NestedBlock.NestingMode
	5,  // 8: terraform.schema.Schema.Object.attributes:type_name -> terraform.schema.Schema.Attribute
	1,  // 9: terraform.schema.Schema.Object.nesting:type_name -> terraform.schema.Schema.Object.NestingMode
	2,  // 10: terraform.schema.Schema.DocString.format:type_name -> terraform.schema.Schema.DocString.Format
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_terraform_schema_proto_init() }
func file_terraform_schema_proto_init() {
	if File_terraform_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terraform_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
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
		file_terraform_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_Block); i {
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
		file_terraform_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_Attribute); i {
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
		file_terraform_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_NestedBlock); i {
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
		file_terraform_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_Object); i {
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
		file_terraform_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_DocString); i {
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
			RawDescriptor: file_terraform_schema_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_terraform_schema_proto_goTypes,
		DependencyIndexes: file_terraform_schema_proto_depIdxs,
		EnumInfos:         file_terraform_schema_proto_enumTypes,
		MessageInfos:      file_terraform_schema_proto_msgTypes,
	}.Build()
	File_terraform_schema_proto = out.File
	file_terraform_schema_proto_rawDesc = nil
	file_terraform_schema_proto_goTypes = nil
	file_terraform_schema_proto_depIdxs = nil
}