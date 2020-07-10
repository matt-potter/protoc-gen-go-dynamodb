// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: dynamopb/dynamopb.proto

package dynamopb

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type KeyDefinitionAttributeType int32

const (
	KeyDefinition_BINARY KeyDefinitionAttributeType = 0
	KeyDefinition_STRING KeyDefinitionAttributeType = 1
	KeyDefinition_NUMBER KeyDefinitionAttributeType = 2
)

// Enum value maps for KeyDefinitionAttributeType.
var (
	KeyDefinitionAttributeType_name = map[int32]string{
		0: "BINARY",
		1: "STRING",
		2: "NUMBER",
	}
	KeyDefinitionAttributeType_value = map[string]int32{
		"BINARY": 0,
		"STRING": 1,
		"NUMBER": 2,
	}
)

func (x KeyDefinitionAttributeType) Enum() *KeyDefinitionAttributeType {
	p := new(KeyDefinitionAttributeType)
	*p = x
	return p
}

func (x KeyDefinitionAttributeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyDefinitionAttributeType) Descriptor() protoreflect.EnumDescriptor {
	return file_dynamopb_dynamopb_proto_enumTypes[0].Descriptor()
}

func (KeyDefinitionAttributeType) Type() protoreflect.EnumType {
	return &file_dynamopb_dynamopb_proto_enumTypes[0]
}

func (x KeyDefinitionAttributeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyDefinitionAttributeType.Descriptor instead.
func (KeyDefinitionAttributeType) EnumDescriptor() ([]byte, []int) {
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{1, 0}
}

type Gsi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PartitionKey *KeyDefinition `protobuf:"bytes,2,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	SortKey      *KeyDefinition `protobuf:"bytes,3,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
}

func (x *Gsi) Reset() {
	*x = Gsi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamopb_dynamopb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gsi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gsi) ProtoMessage() {}

func (x *Gsi) ProtoReflect() protoreflect.Message {
	mi := &file_dynamopb_dynamopb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gsi.ProtoReflect.Descriptor instead.
func (*Gsi) Descriptor() ([]byte, []int) {
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{0}
}

func (x *Gsi) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Gsi) GetPartitionKey() *KeyDefinition {
	if x != nil {
		return x.PartitionKey
	}
	return nil
}

func (x *Gsi) GetSortKey() *KeyDefinition {
	if x != nil {
		return x.SortKey
	}
	return nil
}

type KeyDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type KeyDefinitionAttributeType `protobuf:"varint,2,opt,name=type,proto3,enum=dynamopb.KeyDefinitionAttributeType" json:"type,omitempty"`
}

func (x *KeyDefinition) Reset() {
	*x = KeyDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamopb_dynamopb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyDefinition) ProtoMessage() {}

func (x *KeyDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_dynamopb_dynamopb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyDefinition.ProtoReflect.Descriptor instead.
func (*KeyDefinition) Descriptor() ([]byte, []int) {
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{1}
}

func (x *KeyDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KeyDefinition) GetType() KeyDefinitionAttributeType {
	if x != nil {
		return x.Type
	}
	return KeyDefinition_BINARY
}

var file_dynamopb_dynamopb_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50001,
		Name:          "dynamopb.storable",
		Tag:           "varint,50001,opt,name=storable",
		Filename:      "dynamopb/dynamopb.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*KeyDefinition)(nil),
		Field:         50002,
		Name:          "dynamopb.primary_index",
		Tag:           "bytes,50002,opt,name=primary_index",
		Filename:      "dynamopb/dynamopb.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: ([]*Gsi)(nil),
		Field:         50003,
		Name:          "dynamopb.global_secondary_indexes",
		Tag:           "bytes,50003,rep,name=global_secondary_indexes",
		Filename:      "dynamopb/dynamopb.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional bool storable = 50001;
	E_Storable = &file_dynamopb_dynamopb_proto_extTypes[0]
	// optional dynamopb.key_definition primary_index = 50002;
	E_PrimaryIndex = &file_dynamopb_dynamopb_proto_extTypes[1]
	// repeated dynamopb.gsi global_secondary_indexes = 50003;
	E_GlobalSecondaryIndexes = &file_dynamopb_dynamopb_proto_extTypes[2]
)

var File_dynamopb_dynamopb_proto protoreflect.FileDescriptor

var file_dynamopb_dynamopb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x70, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x03, 0x67, 0x73, 0x69, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x70, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x33, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x6b, 0x65,
	0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x97, 0x01, 0x0a, 0x0e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x42,
	0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x3a,
	0x3d, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x60,
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x70, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x3a, 0x6a, 0x0a, 0x18, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x86,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62,
	0x2e, 0x67, 0x73, 0x69, 0x52, 0x16, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x74, 0x2d,
	0x70, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2f, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dynamopb_dynamopb_proto_rawDescOnce sync.Once
	file_dynamopb_dynamopb_proto_rawDescData = file_dynamopb_dynamopb_proto_rawDesc
)

func file_dynamopb_dynamopb_proto_rawDescGZIP() []byte {
	file_dynamopb_dynamopb_proto_rawDescOnce.Do(func() {
		file_dynamopb_dynamopb_proto_rawDescData = protoimpl.X.CompressGZIP(file_dynamopb_dynamopb_proto_rawDescData)
	})
	return file_dynamopb_dynamopb_proto_rawDescData
}

var file_dynamopb_dynamopb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dynamopb_dynamopb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dynamopb_dynamopb_proto_goTypes = []interface{}{
	(KeyDefinitionAttributeType)(0),   // 0: dynamopb.key_definition.attribute_type
	(*Gsi)(nil),                       // 1: dynamopb.gsi
	(*KeyDefinition)(nil),             // 2: dynamopb.key_definition
	(*descriptor.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
}
var file_dynamopb_dynamopb_proto_depIdxs = []int32{
	2, // 0: dynamopb.gsi.partition_key:type_name -> dynamopb.key_definition
	2, // 1: dynamopb.gsi.sort_key:type_name -> dynamopb.key_definition
	0, // 2: dynamopb.key_definition.type:type_name -> dynamopb.key_definition.attribute_type
	3, // 3: dynamopb.storable:extendee -> google.protobuf.MessageOptions
	3, // 4: dynamopb.primary_index:extendee -> google.protobuf.MessageOptions
	3, // 5: dynamopb.global_secondary_indexes:extendee -> google.protobuf.MessageOptions
	2, // 6: dynamopb.primary_index:type_name -> dynamopb.key_definition
	1, // 7: dynamopb.global_secondary_indexes:type_name -> dynamopb.gsi
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	3, // [3:6] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dynamopb_dynamopb_proto_init() }
func file_dynamopb_dynamopb_proto_init() {
	if File_dynamopb_dynamopb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dynamopb_dynamopb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gsi); i {
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
		file_dynamopb_dynamopb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyDefinition); i {
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
			RawDescriptor: file_dynamopb_dynamopb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_dynamopb_dynamopb_proto_goTypes,
		DependencyIndexes: file_dynamopb_dynamopb_proto_depIdxs,
		EnumInfos:         file_dynamopb_dynamopb_proto_enumTypes,
		MessageInfos:      file_dynamopb_dynamopb_proto_msgTypes,
		ExtensionInfos:    file_dynamopb_dynamopb_proto_extTypes,
	}.Build()
	File_dynamopb_dynamopb_proto = out.File
	file_dynamopb_dynamopb_proto_rawDesc = nil
	file_dynamopb_dynamopb_proto_goTypes = nil
	file_dynamopb_dynamopb_proto_depIdxs = nil
}
