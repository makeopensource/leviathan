// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: types/v1/types.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LabData struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	Labname                  string                 `protobuf:"bytes,1,opt,name=labname,proto3" json:"labname,omitempty"`
	EntryCmd                 string                 `protobuf:"bytes,2,opt,name=entryCmd,proto3" json:"entryCmd,omitempty"`
	JobTimeoutInSeconds      uint64                 `protobuf:"varint,3,opt,name=jobTimeoutInSeconds,proto3" json:"jobTimeoutInSeconds,omitempty"`
	AutolabCompatibilityMode bool                   `protobuf:"varint,4,opt,name=autolabCompatibilityMode,proto3" json:"autolabCompatibilityMode,omitempty"`
	Limits                   *MachineLimits         `protobuf:"bytes,5,opt,name=limits,proto3" json:"limits,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *LabData) Reset() {
	*x = LabData{}
	mi := &file_types_v1_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabData) ProtoMessage() {}

func (x *LabData) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabData.ProtoReflect.Descriptor instead.
func (*LabData) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *LabData) GetLabname() string {
	if x != nil {
		return x.Labname
	}
	return ""
}

func (x *LabData) GetEntryCmd() string {
	if x != nil {
		return x.EntryCmd
	}
	return ""
}

func (x *LabData) GetJobTimeoutInSeconds() uint64 {
	if x != nil {
		return x.JobTimeoutInSeconds
	}
	return 0
}

func (x *LabData) GetAutolabCompatibilityMode() bool {
	if x != nil {
		return x.AutolabCompatibilityMode
	}
	return false
}

func (x *LabData) GetLimits() *MachineLimits {
	if x != nil {
		return x.Limits
	}
	return nil
}

type MachineLimits struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CPUCores      int32                  `protobuf:"varint,1,opt,name=CPUCores,proto3" json:"CPUCores,omitempty"`
	MemoryInMb    int32                  `protobuf:"varint,2,opt,name=memoryInMb,proto3" json:"memoryInMb,omitempty"`
	PidLimit      int32                  `protobuf:"varint,3,opt,name=PidLimit,proto3" json:"PidLimit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineLimits) Reset() {
	*x = MachineLimits{}
	mi := &file_types_v1_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineLimits) ProtoMessage() {}

func (x *MachineLimits) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineLimits.ProtoReflect.Descriptor instead.
func (*MachineLimits) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *MachineLimits) GetCPUCores() int32 {
	if x != nil {
		return x.CPUCores
	}
	return 0
}

func (x *MachineLimits) GetMemoryInMb() int32 {
	if x != nil {
		return x.MemoryInMb
	}
	return 0
}

func (x *MachineLimits) GetPidLimit() int32 {
	if x != nil {
		return x.PidLimit
	}
	return 0
}

var File_types_v1_types_proto protoreflect.FileDescriptor

const file_types_v1_types_proto_rawDesc = "" +
	"\n" +
	"\x14types/v1/types.proto\x12\btypes.v1\"\xde\x01\n" +
	"\aLabData\x12\x18\n" +
	"\alabname\x18\x01 \x01(\tR\alabname\x12\x1a\n" +
	"\bentryCmd\x18\x02 \x01(\tR\bentryCmd\x120\n" +
	"\x13jobTimeoutInSeconds\x18\x03 \x01(\x04R\x13jobTimeoutInSeconds\x12:\n" +
	"\x18autolabCompatibilityMode\x18\x04 \x01(\bR\x18autolabCompatibilityMode\x12/\n" +
	"\x06limits\x18\x05 \x01(\v2\x17.types.v1.MachineLimitsR\x06limits\"g\n" +
	"\rMachineLimits\x12\x1a\n" +
	"\bCPUCores\x18\x01 \x01(\x05R\bCPUCores\x12\x1e\n" +
	"\n" +
	"memoryInMb\x18\x02 \x01(\x05R\n" +
	"memoryInMb\x12\x1a\n" +
	"\bPidLimit\x18\x03 \x01(\x05R\bPidLimitB\x93\x01\n" +
	"\fcom.types.v1B\n" +
	"TypesProtoP\x01Z6github.com/makeopensource/leviathan/generated/types/v1\xa2\x02\x03TXX\xaa\x02\bTypes.V1\xca\x02\bTypes\\V1\xe2\x02\x14Types\\V1\\GPBMetadata\xea\x02\tTypes::V1b\x06proto3"

var (
	file_types_v1_types_proto_rawDescOnce sync.Once
	file_types_v1_types_proto_rawDescData []byte
)

func file_types_v1_types_proto_rawDescGZIP() []byte {
	file_types_v1_types_proto_rawDescOnce.Do(func() {
		file_types_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_v1_types_proto_rawDesc), len(file_types_v1_types_proto_rawDesc)))
	})
	return file_types_v1_types_proto_rawDescData
}

var file_types_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_v1_types_proto_goTypes = []any{
	(*LabData)(nil),       // 0: types.v1.LabData
	(*MachineLimits)(nil), // 1: types.v1.MachineLimits
}
var file_types_v1_types_proto_depIdxs = []int32{
	1, // 0: types.v1.LabData.limits:type_name -> types.v1.MachineLimits
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_v1_types_proto_init() }
func file_types_v1_types_proto_init() {
	if File_types_v1_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_v1_types_proto_rawDesc), len(file_types_v1_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_types_proto_goTypes,
		DependencyIndexes: file_types_v1_types_proto_depIdxs,
		MessageInfos:      file_types_v1_types_proto_msgTypes,
	}.Build()
	File_types_v1_types_proto = out.File
	file_types_v1_types_proto_goTypes = nil
	file_types_v1_types_proto_depIdxs = nil
}
