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
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{2, 0}
}

type Cfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName              string   `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	PrimaryIndex           *Index   `protobuf:"bytes,2,opt,name=primary_index,json=primaryIndex,proto3" json:"primary_index,omitempty"`
	GlobalSecondaryIndexes []*Index `protobuf:"bytes,3,rep,name=global_secondary_indexes,json=globalSecondaryIndexes,proto3" json:"global_secondary_indexes,omitempty"`
}

func (x *Cfg) Reset() {
	*x = Cfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamopb_dynamopb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cfg) ProtoMessage() {}

func (x *Cfg) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Cfg.ProtoReflect.Descriptor instead.
func (*Cfg) Descriptor() ([]byte, []int) {
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{0}
}

func (x *Cfg) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Cfg) GetPrimaryIndex() *Index {
	if x != nil {
		return x.PrimaryIndex
	}
	return nil
}

func (x *Cfg) GetGlobalSecondaryIndexes() []*Index {
	if x != nil {
		return x.GlobalSecondaryIndexes
	}
	return nil
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IndexName    string         `protobuf:"bytes,1,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	PartitionKey *KeyDefinition `protobuf:"bytes,2,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	SortKey      *KeyDefinition `protobuf:"bytes,3,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamopb_dynamopb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{1}
}

func (x *Index) GetIndexName() string {
	if x != nil {
		return x.IndexName
	}
	return ""
}

func (x *Index) GetPartitionKey() *KeyDefinition {
	if x != nil {
		return x.PartitionKey
	}
	return nil
}

func (x *Index) GetSortKey() *KeyDefinition {
	if x != nil {
		return x.SortKey
	}
	return nil
}

type KeyDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrName string                     `protobuf:"bytes,1,opt,name=attr_name,json=attrName,proto3" json:"attr_name,omitempty"`
	AttrType KeyDefinitionAttributeType `protobuf:"varint,2,opt,name=attr_type,json=attrType,proto3,enum=dynamopb.KeyDefinitionAttributeType" json:"attr_type,omitempty"`
}

func (x *KeyDefinition) Reset() {
	*x = KeyDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamopb_dynamopb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyDefinition) ProtoMessage() {}

func (x *KeyDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_dynamopb_dynamopb_proto_msgTypes[2]
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
	return file_dynamopb_dynamopb_proto_rawDescGZIP(), []int{2}
}

func (x *KeyDefinition) GetAttrName() string {
	if x != nil {
		return x.AttrName
	}
	return ""
}

func (x *KeyDefinition) GetAttrType() KeyDefinitionAttributeType {
	if x != nil {
		return x.AttrType
	}
	return KeyDefinition_BINARY
}

var file_dynamopb_dynamopb_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*Cfg)(nil),
		Field:         50001,
		Name:          "dynamopb.config",
		Tag:           "bytes,50001,opt,name=config",
		Filename:      "dynamopb/dynamopb.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional dynamopb.cfg config = 50001;
	E_Config = &file_dynamopb_dynamopb_proto_extTypes[0]
)

var File_dynamopb_dynamopb_proto protoreflect.FileDescriptor

var file_dynamopb_dynamopb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x70, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x03, 0x63, 0x66, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0d,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x49, 0x0a, 0x18, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x16, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x22, 0x9a, 0x01,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f,
	0x70, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x6b,
	0x65, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x61, 0x74,
	0x74, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x34, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x3a, 0x48, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x63, 0x66, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x74, 0x74, 0x2d, 0x70, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64,
	0x62, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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
var file_dynamopb_dynamopb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dynamopb_dynamopb_proto_goTypes = []interface{}{
	(KeyDefinitionAttributeType)(0),   // 0: dynamopb.key_definition.attribute_type
	(*Cfg)(nil),                       // 1: dynamopb.cfg
	(*Index)(nil),                     // 2: dynamopb.index
	(*KeyDefinition)(nil),             // 3: dynamopb.key_definition
	(*descriptor.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
}
var file_dynamopb_dynamopb_proto_depIdxs = []int32{
	2, // 0: dynamopb.cfg.primary_index:type_name -> dynamopb.index
	2, // 1: dynamopb.cfg.global_secondary_indexes:type_name -> dynamopb.index
	3, // 2: dynamopb.index.partition_key:type_name -> dynamopb.key_definition
	3, // 3: dynamopb.index.sort_key:type_name -> dynamopb.key_definition
	0, // 4: dynamopb.key_definition.attr_type:type_name -> dynamopb.key_definition.attribute_type
	4, // 5: dynamopb.config:extendee -> google.protobuf.MessageOptions
	1, // 6: dynamopb.config:type_name -> dynamopb.cfg
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	6, // [6:7] is the sub-list for extension type_name
	5, // [5:6] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_dynamopb_dynamopb_proto_init() }
func file_dynamopb_dynamopb_proto_init() {
	if File_dynamopb_dynamopb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dynamopb_dynamopb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cfg); i {
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
			switch v := v.(*Index); i {
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
		file_dynamopb_dynamopb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   3,
			NumExtensions: 1,
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
