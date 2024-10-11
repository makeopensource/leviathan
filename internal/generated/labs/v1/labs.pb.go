// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: labs/v1/labs.proto

package v1

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

type LabRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabName    string      `protobuf:"bytes,1,opt,name=LabName,proto3" json:"LabName,omitempty"`
	MakeFile   *FileUpload `protobuf:"bytes,2,opt,name=makeFile,proto3" json:"makeFile,omitempty"`
	GraderFile *FileUpload `protobuf:"bytes,3,opt,name=graderFile,proto3" json:"graderFile,omitempty"`
}

func (x *LabRequest) Reset() {
	*x = LabRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labs_v1_labs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabRequest) ProtoMessage() {}

func (x *LabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *LabRequest) GetMakeFile() *FileUpload {
	if x != nil {
		return x.MakeFile
	}
	return nil
}

func (x *LabRequest) GetGraderFile() *FileUpload {
	if x != nil {
		return x.GraderFile
	}
	return nil
}

type NewLabResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewLabResponse) Reset() {
	*x = NewLabResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labs_v1_labs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLabResponse) ProtoMessage() {}

func (x *NewLabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditLabResponse) Reset() {
	*x = EditLabResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labs_v1_labs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditLabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditLabResponse) ProtoMessage() {}

func (x *EditLabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabName string `protobuf:"bytes,1,opt,name=LabName,proto3" json:"LabName,omitempty"`
}

func (x *DeleteLabRequest) Reset() {
	*x = DeleteLabRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labs_v1_labs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabRequest) ProtoMessage() {}

func (x *DeleteLabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLabResponse) Reset() {
	*x = DeleteLabResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labs_v1_labs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabResponse) ProtoMessage() {}

func (x *DeleteLabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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

type FileUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileUpload) Reset() {
	*x = FileUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labs_v1_labs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUpload) ProtoMessage() {}

func (x *FileUpload) ProtoReflect() protoreflect.Message {
	mi := &file_labs_v1_labs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUpload.ProtoReflect.Descriptor instead.
func (*FileUpload) Descriptor() ([]byte, []int) {
	return file_labs_v1_labs_proto_rawDescGZIP(), []int{5}
}

func (x *FileUpload) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileUpload) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_labs_v1_labs_proto protoreflect.FileDescriptor

var file_labs_v1_labs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x8c, 0x01,
	0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c,
	0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x61, 0x6b, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x6d,
	0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x61,
	0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x10, 0x0a, 0x0e,
	0x4e, 0x65, 0x77, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11,
	0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xc8, 0x01, 0x0a, 0x0a, 0x4c, 0x61, 0x62,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x4c, 0x61,
	0x62, 0x12, 0x13, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x12, 0x13, 0x2e, 0x6c,
	0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x95, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x61, 0x62, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c, 0x61, 0x62, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x6b,
	0x65, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x65, 0x76, 0x69,
	0x61, 0x74, 0x68, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4c, 0x61, 0x62, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x07, 0x4c, 0x61, 0x62, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4c, 0x61, 0x62,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x08, 0x4c, 0x61, 0x62, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_labs_v1_labs_proto_rawDescOnce sync.Once
	file_labs_v1_labs_proto_rawDescData = file_labs_v1_labs_proto_rawDesc
)

func file_labs_v1_labs_proto_rawDescGZIP() []byte {
	file_labs_v1_labs_proto_rawDescOnce.Do(func() {
		file_labs_v1_labs_proto_rawDescData = protoimpl.X.CompressGZIP(file_labs_v1_labs_proto_rawDescData)
	})
	return file_labs_v1_labs_proto_rawDescData
}

var file_labs_v1_labs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_labs_v1_labs_proto_goTypes = []any{
	(*LabRequest)(nil),        // 0: labs.v1.LabRequest
	(*NewLabResponse)(nil),    // 1: labs.v1.NewLabResponse
	(*EditLabResponse)(nil),   // 2: labs.v1.EditLabResponse
	(*DeleteLabRequest)(nil),  // 3: labs.v1.DeleteLabRequest
	(*DeleteLabResponse)(nil), // 4: labs.v1.DeleteLabResponse
	(*FileUpload)(nil),        // 5: labs.v1.FileUpload
}
var file_labs_v1_labs_proto_depIdxs = []int32{
	5, // 0: labs.v1.LabRequest.makeFile:type_name -> labs.v1.FileUpload
	5, // 1: labs.v1.LabRequest.graderFile:type_name -> labs.v1.FileUpload
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
	if !protoimpl.UnsafeEnabled {
		file_labs_v1_labs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LabRequest); i {
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
		file_labs_v1_labs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NewLabResponse); i {
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
		file_labs_v1_labs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EditLabResponse); i {
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
		file_labs_v1_labs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteLabRequest); i {
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
		file_labs_v1_labs_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteLabResponse); i {
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
		file_labs_v1_labs_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FileUpload); i {
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
			RawDescriptor: file_labs_v1_labs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_labs_v1_labs_proto_goTypes,
		DependencyIndexes: file_labs_v1_labs_proto_depIdxs,
		MessageInfos:      file_labs_v1_labs_proto_msgTypes,
	}.Build()
	File_labs_v1_labs_proto = out.File
	file_labs_v1_labs_proto_rawDesc = nil
	file_labs_v1_labs_proto_goTypes = nil
	file_labs_v1_labs_proto_depIdxs = nil
}
