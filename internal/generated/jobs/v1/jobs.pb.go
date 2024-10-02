// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: jobs/v1/jobs.proto

package jobs

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

// todo figure out request/response
type NewJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewJobRequest) Reset() {
	*x = NewJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobRequest) ProtoMessage() {}

func (x *NewJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewJobRequest.ProtoReflect.Descriptor instead.
func (*NewJobRequest) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{0}
}

type NewJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewJobResponse) Reset() {
	*x = NewJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobResponse) ProtoMessage() {}

func (x *NewJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewJobResponse.ProtoReflect.Descriptor instead.
func (*NewJobResponse) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{1}
}

type JobStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobStatusRequest) Reset() {
	*x = JobStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusRequest) ProtoMessage() {}

func (x *JobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusRequest.ProtoReflect.Descriptor instead.
func (*JobStatusRequest) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{2}
}

type JobStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobStatusResponse) Reset() {
	*x = JobStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusResponse) ProtoMessage() {}

func (x *JobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusResponse.ProtoReflect.Descriptor instead.
func (*JobStatusResponse) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{3}
}

type CancelJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelJobRequest) Reset() {
	*x = CancelJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobRequest) ProtoMessage() {}

func (x *CancelJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJobRequest.ProtoReflect.Descriptor instead.
func (*CancelJobRequest) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{4}
}

type CancelJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelJobResponse) Reset() {
	*x = CancelJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobResponse) ProtoMessage() {}

func (x *CancelJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJobResponse.ProtoReflect.Descriptor instead.
func (*CancelJobResponse) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{5}
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{6}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageResponse string `protobuf:"bytes,1,opt,name=message_response,json=messageResponse,proto3" json:"message_response,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_v1_jobs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{7}
}

func (x *EchoResponse) GetMessageResponse() string {
	if x != nil {
		return x.MessageResponse
	}
	return ""
}

var File_jobs_v1_jobs_proto protoreflect.FileDescriptor

var file_jobs_v1_jobs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x0f, 0x0a,
	0x0d, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a,
	0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x0c, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8c, 0x02, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x12,
	0x16, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x19, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6a, 0x6f, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x12, 0x19, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jobs_v1_jobs_proto_rawDescOnce sync.Once
	file_jobs_v1_jobs_proto_rawDescData = file_jobs_v1_jobs_proto_rawDesc
)

func file_jobs_v1_jobs_proto_rawDescGZIP() []byte {
	file_jobs_v1_jobs_proto_rawDescOnce.Do(func() {
		file_jobs_v1_jobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_jobs_v1_jobs_proto_rawDescData)
	})
	return file_jobs_v1_jobs_proto_rawDescData
}

var file_jobs_v1_jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_jobs_v1_jobs_proto_goTypes = []any{
	(*NewJobRequest)(nil),     // 0: jobs.v1.NewJobRequest
	(*NewJobResponse)(nil),    // 1: jobs.v1.NewJobResponse
	(*JobStatusRequest)(nil),  // 2: jobs.v1.JobStatusRequest
	(*JobStatusResponse)(nil), // 3: jobs.v1.JobStatusResponse
	(*CancelJobRequest)(nil),  // 4: jobs.v1.CancelJobRequest
	(*CancelJobResponse)(nil), // 5: jobs.v1.CancelJobResponse
	(*EchoRequest)(nil),       // 6: jobs.v1.EchoRequest
	(*EchoResponse)(nil),      // 7: jobs.v1.EchoResponse
}
var file_jobs_v1_jobs_proto_depIdxs = []int32{
	0, // 0: jobs.v1.JobService.NewJob:input_type -> jobs.v1.NewJobRequest
	2, // 1: jobs.v1.JobService.JobStatus:input_type -> jobs.v1.JobStatusRequest
	4, // 2: jobs.v1.JobService.CancelJob:input_type -> jobs.v1.CancelJobRequest
	6, // 3: jobs.v1.JobService.Echo:input_type -> jobs.v1.EchoRequest
	1, // 4: jobs.v1.JobService.NewJob:output_type -> jobs.v1.NewJobResponse
	3, // 5: jobs.v1.JobService.JobStatus:output_type -> jobs.v1.JobStatusResponse
	5, // 6: jobs.v1.JobService.CancelJob:output_type -> jobs.v1.CancelJobResponse
	7, // 7: jobs.v1.JobService.Echo:output_type -> jobs.v1.EchoResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jobs_v1_jobs_proto_init() }
func file_jobs_v1_jobs_proto_init() {
	if File_jobs_v1_jobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jobs_v1_jobs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NewJobRequest); i {
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
		file_jobs_v1_jobs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NewJobResponse); i {
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
		file_jobs_v1_jobs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*JobStatusRequest); i {
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
		file_jobs_v1_jobs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*JobStatusResponse); i {
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
		file_jobs_v1_jobs_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CancelJobRequest); i {
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
		file_jobs_v1_jobs_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CancelJobResponse); i {
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
		file_jobs_v1_jobs_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EchoRequest); i {
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
		file_jobs_v1_jobs_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_jobs_v1_jobs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jobs_v1_jobs_proto_goTypes,
		DependencyIndexes: file_jobs_v1_jobs_proto_depIdxs,
		MessageInfos:      file_jobs_v1_jobs_proto_msgTypes,
	}.Build()
	File_jobs_v1_jobs_proto = out.File
	file_jobs_v1_jobs_proto_rawDesc = nil
	file_jobs_v1_jobs_proto_goTypes = nil
	file_jobs_v1_jobs_proto_depIdxs = nil
}
