// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: labs/v1/labs.proto

package v1

import (
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
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

type LabRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LabName       string                 `protobuf:"bytes,1,opt,name=LabName,proto3" json:"LabName,omitempty"`
	MakeFile      *v1.FileUpload         `protobuf:"bytes,2,opt,name=makeFile,proto3" json:"makeFile,omitempty"`
	GraderFile    *v1.FileUpload         `protobuf:"bytes,3,opt,name=graderFile,proto3" json:"graderFile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabRequest) Reset() {
	*x = LabRequest{}
	mi := &file_labs_v1_labs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabRequest) ProtoMessage() {}

func (x *LabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabRequest.ProtoReflect.Descriptor instead.
func (*LabRequest) Descriptor() ([]byte, []int) {
	return file_labs_v1_labs_proto_rawDescGZIP(), []int{0}
}

func (x *LabRequest) GetLabName() string {
	if x != nil {
		return x.LabName
	}
	return ""
}

func (x *LabRequest) GetMakeFile() *v1.FileUpload {
	if x != nil {
		return x.MakeFile
	}
	return nil
}

func (x *LabRequest) GetGraderFile() *v1.FileUpload {
	if x != nil {
		return x.GraderFile
	}
	return nil
}

type NewLabResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewLabResponse) Reset() {
	*x = NewLabResponse{}
	mi := &file_labs_v1_labs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewLabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLabResponse) ProtoMessage() {}

func (x *NewLabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLabResponse.ProtoReflect.Descriptor instead.
func (*NewLabResponse) Descriptor() ([]byte, []int) {
	return file_labs_v1_labs_proto_rawDescGZIP(), []int{1}
}

type EditLabResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditLabResponse) Reset() {
	*x = EditLabResponse{}
	mi := &file_labs_v1_labs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditLabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditLabResponse) ProtoMessage() {}

func (x *EditLabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditLabResponse.ProtoReflect.Descriptor instead.
func (*EditLabResponse) Descriptor() ([]byte, []int) {
	return file_labs_v1_labs_proto_rawDescGZIP(), []int{2}
}

type DeleteLabRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LabName       string                 `protobuf:"bytes,1,opt,name=LabName,proto3" json:"LabName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteLabRequest) Reset() {
	*x = DeleteLabRequest{}
	mi := &file_labs_v1_labs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabRequest) ProtoMessage() {}

func (x *DeleteLabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLabRequest.ProtoReflect.Descriptor instead.
func (*DeleteLabRequest) Descriptor() ([]byte, []int) {
	return file_labs_v1_labs_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteLabRequest) GetLabName() string {
	if x != nil {
		return x.LabName
	}
	return ""
}

type DeleteLabResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteLabResponse) Reset() {
	*x = DeleteLabResponse{}
	mi := &file_labs_v1_labs_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabResponse) ProtoMessage() {}

func (x *DeleteLabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLabResponse.ProtoReflect.Descriptor instead.
func (*DeleteLabResponse) Descriptor() ([]byte, []int) {
	return file_labs_v1_labs_proto_rawDescGZIP(), []int{4}
}

var File_labs_v1_labs_proto protoreflect.FileDescriptor

var file_labs_v1_labs_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08,
	0x6d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x6d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34,
	0x0a, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x4c, 0x61, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x4c, 0x61,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc8, 0x01, 0x0a,
	0x0a, 0x4c, 0x61, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x4e,
	0x65, 0x77, 0x4c, 0x61, 0x62, 0x12, 0x13, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x61, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x4c, 0x61, 0x62,
	0x12, 0x13, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x12, 0x19,
	0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x61, 0x62, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8c, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c, 0x61, 0x62, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6c, 0x65, 0x76, 0x69, 0x61, 0x74, 0x68, 0x61, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4c, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x4c, 0x61, 0x62, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4c, 0x61,
	0x62, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4c, 0x61, 0x62, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4c, 0x61,
	0x62, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_labs_v1_labs_proto_rawDescOnce sync.Once
	file_labs_v1_labs_proto_rawDescData []byte
)

func file_labs_v1_labs_proto_rawDescGZIP() []byte {
	file_labs_v1_labs_proto_rawDescOnce.Do(func() {
		file_labs_v1_labs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_labs_v1_labs_proto_rawDesc), len(file_labs_v1_labs_proto_rawDesc)))
	})
	return file_labs_v1_labs_proto_rawDescData
}

var file_labs_v1_labs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_labs_v1_labs_proto_goTypes = []any{
	(*LabRequest)(nil),        // 0: labs.v1.LabRequest
	(*NewLabResponse)(nil),    // 1: labs.v1.NewLabResponse
	(*EditLabResponse)(nil),   // 2: labs.v1.EditLabResponse
	(*DeleteLabRequest)(nil),  // 3: labs.v1.DeleteLabRequest
	(*DeleteLabResponse)(nil), // 4: labs.v1.DeleteLabResponse
	(*v1.FileUpload)(nil),     // 5: types.v1.FileUpload
}
var file_labs_v1_labs_proto_depIdxs = []int32{
	5, // 0: labs.v1.LabRequest.makeFile:type_name -> types.v1.FileUpload
	5, // 1: labs.v1.LabRequest.graderFile:type_name -> types.v1.FileUpload
	0, // 2: labs.v1.LabService.NewLab:input_type -> labs.v1.LabRequest
	0, // 3: labs.v1.LabService.EditLab:input_type -> labs.v1.LabRequest
	3, // 4: labs.v1.LabService.DeleteLab:input_type -> labs.v1.DeleteLabRequest
	1, // 5: labs.v1.LabService.NewLab:output_type -> labs.v1.NewLabResponse
	2, // 6: labs.v1.LabService.EditLab:output_type -> labs.v1.EditLabResponse
	4, // 7: labs.v1.LabService.DeleteLab:output_type -> labs.v1.DeleteLabResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_labs_v1_labs_proto_init() }
func file_labs_v1_labs_proto_init() {
	if File_labs_v1_labs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_labs_v1_labs_proto_rawDesc), len(file_labs_v1_labs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_labs_v1_labs_proto_goTypes,
		DependencyIndexes: file_labs_v1_labs_proto_depIdxs,
		MessageInfos:      file_labs_v1_labs_proto_msgTypes,
	}.Build()
	File_labs_v1_labs_proto = out.File
	file_labs_v1_labs_proto_goTypes = nil
	file_labs_v1_labs_proto_depIdxs = nil
}
