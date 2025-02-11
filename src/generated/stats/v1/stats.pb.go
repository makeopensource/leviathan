// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: stats/v1/stats.proto

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

type EchoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	mi := &file_stats_v1_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	MessageResponse string                 `protobuf:"bytes,1,opt,name=message_response,json=messageResponse,proto3" json:"message_response,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	mi := &file_stats_v1_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetMessageResponse() string {
	if x != nil {
		return x.MessageResponse
	}
	return ""
}

var File_stats_v1_stats_proto protoreflect.FileDescriptor

var file_stats_v1_stats_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x27, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x0c, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x47, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x15, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x93, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x6f, 0x70, 0x65,
	0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x65, 0x76, 0x69, 0x61, 0x74, 0x68, 0x61,
	0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x53, 0x74, 0x61, 0x74, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x14, 0x53, 0x74, 0x61, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_stats_v1_stats_proto_rawDescOnce sync.Once
	file_stats_v1_stats_proto_rawDescData []byte
)

func file_stats_v1_stats_proto_rawDescGZIP() []byte {
	file_stats_v1_stats_proto_rawDescOnce.Do(func() {
		file_stats_v1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stats_v1_stats_proto_rawDesc), len(file_stats_v1_stats_proto_rawDesc)))
	})
	return file_stats_v1_stats_proto_rawDescData
}

var file_stats_v1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stats_v1_stats_proto_goTypes = []any{
	(*EchoRequest)(nil),  // 0: stats.v1.EchoRequest
	(*EchoResponse)(nil), // 1: stats.v1.EchoResponse
}
var file_stats_v1_stats_proto_depIdxs = []int32{
	0, // 0: stats.v1.StatsService.Echo:input_type -> stats.v1.EchoRequest
	1, // 1: stats.v1.StatsService.Echo:output_type -> stats.v1.EchoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stats_v1_stats_proto_init() }
func file_stats_v1_stats_proto_init() {
	if File_stats_v1_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stats_v1_stats_proto_rawDesc), len(file_stats_v1_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stats_v1_stats_proto_goTypes,
		DependencyIndexes: file_stats_v1_stats_proto_depIdxs,
		MessageInfos:      file_stats_v1_stats_proto_msgTypes,
	}.Build()
	File_stats_v1_stats_proto = out.File
	file_stats_v1_stats_proto_goTypes = nil
	file_stats_v1_stats_proto_depIdxs = nil
}
