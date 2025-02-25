// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: jobs/v1/jobs.proto

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

type NewJobRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	MakeFile            *v1.FileUpload         `protobuf:"bytes,1,opt,name=makeFile,proto3" json:"makeFile,omitempty"`
	GraderFile          *v1.FileUpload         `protobuf:"bytes,2,opt,name=graderFile,proto3" json:"graderFile,omitempty"`
	StudentSubmission   *v1.FileUpload         `protobuf:"bytes,3,opt,name=studentSubmission,proto3" json:"studentSubmission,omitempty"`
	DockerFile          *v1.FileUpload         `protobuf:"bytes,4,opt,name=dockerFile,proto3" json:"dockerFile,omitempty"`
	ImageName           string                 `protobuf:"bytes,5,opt,name=imageName,proto3" json:"imageName,omitempty"`
	JobTimeoutInSeconds uint64                 `protobuf:"varint,6,opt,name=jobTimeoutInSeconds,proto3" json:"jobTimeoutInSeconds,omitempty"`
	EntryCmd            string                 `protobuf:"bytes,7,opt,name=entryCmd,proto3" json:"entryCmd,omitempty"`
	Limits              *MachineLimits         `protobuf:"bytes,8,opt,name=limits,proto3" json:"limits,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *NewJobRequest) Reset() {
	*x = NewJobRequest{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobRequest) ProtoMessage() {}

func (x *NewJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[0]
	if x != nil {
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

func (x *NewJobRequest) GetMakeFile() *v1.FileUpload {
	if x != nil {
		return x.MakeFile
	}
	return nil
}

func (x *NewJobRequest) GetGraderFile() *v1.FileUpload {
	if x != nil {
		return x.GraderFile
	}
	return nil
}

func (x *NewJobRequest) GetStudentSubmission() *v1.FileUpload {
	if x != nil {
		return x.StudentSubmission
	}
	return nil
}

func (x *NewJobRequest) GetDockerFile() *v1.FileUpload {
	if x != nil {
		return x.DockerFile
	}
	return nil
}

func (x *NewJobRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *NewJobRequest) GetJobTimeoutInSeconds() uint64 {
	if x != nil {
		return x.JobTimeoutInSeconds
	}
	return 0
}

func (x *NewJobRequest) GetEntryCmd() string {
	if x != nil {
		return x.EntryCmd
	}
	return ""
}

func (x *NewJobRequest) GetLimits() *MachineLimits {
	if x != nil {
		return x.Limits
	}
	return nil
}

type NewJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewJobResponse) Reset() {
	*x = NewJobResponse{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobResponse) ProtoMessage() {}

func (x *NewJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[1]
	if x != nil {
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

func (x *NewJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type CancelJobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelJobRequest) Reset() {
	*x = CancelJobRequest{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobRequest) ProtoMessage() {}

func (x *CancelJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[2]
	if x != nil {
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
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{2}
}

func (x *CancelJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type CancelJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelJobResponse) Reset() {
	*x = CancelJobResponse{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobResponse) ProtoMessage() {}

func (x *CancelJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[3]
	if x != nil {
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
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{3}
}

type JobLogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobLogRequest) Reset() {
	*x = JobLogRequest{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogRequest) ProtoMessage() {}

func (x *JobLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogRequest.ProtoReflect.Descriptor instead.
func (*JobLogRequest) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{4}
}

func (x *JobLogRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type JobLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobInfo       *JobStatus             `protobuf:"bytes,1,opt,name=jobInfo,proto3" json:"jobInfo,omitempty"`
	Logs          string                 `protobuf:"bytes,2,opt,name=logs,proto3" json:"logs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobLogsResponse) Reset() {
	*x = JobLogsResponse{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogsResponse) ProtoMessage() {}

func (x *JobLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogsResponse.ProtoReflect.Descriptor instead.
func (*JobLogsResponse) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{5}
}

func (x *JobLogsResponse) GetJobInfo() *JobStatus {
	if x != nil {
		return x.JobInfo
	}
	return nil
}

func (x *JobLogsResponse) GetLogs() string {
	if x != nil {
		return x.Logs
	}
	return ""
}

type JobStatus struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	JobId            string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	MachineId        string                 `protobuf:"bytes,2,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	ContainerId      string                 `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Status           string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	StatusMessage    string                 `protobuf:"bytes,5,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	OutputFilePath   string                 `protobuf:"bytes,6,opt,name=output_file_path,json=outputFilePath,proto3" json:"output_file_path,omitempty"`
	TmpJobFolderPath string                 `protobuf:"bytes,7,opt,name=tmp_job_folder_path,json=tmpJobFolderPath,proto3" json:"tmp_job_folder_path,omitempty"`
	JobTimeout       int64                  `protobuf:"varint,8,opt,name=job_timeout,json=jobTimeout,proto3" json:"job_timeout,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{6}
}

func (x *JobStatus) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobStatus) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *JobStatus) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *JobStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobStatus) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *JobStatus) GetOutputFilePath() string {
	if x != nil {
		return x.OutputFilePath
	}
	return ""
}

func (x *JobStatus) GetTmpJobFolderPath() string {
	if x != nil {
		return x.TmpJobFolderPath
	}
	return ""
}

func (x *JobStatus) GetJobTimeout() int64 {
	if x != nil {
		return x.JobTimeout
	}
	return 0
}

type MachineLimits struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CPUCores      float32                `protobuf:"fixed32,1,opt,name=CPUCores,proto3" json:"CPUCores,omitempty"`
	MemoryInMb    int64                  `protobuf:"varint,2,opt,name=memoryInMb,proto3" json:"memoryInMb,omitempty"`
	PidLimit      int64                  `protobuf:"varint,3,opt,name=PidLimit,proto3" json:"PidLimit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineLimits) Reset() {
	*x = MachineLimits{}
	mi := &file_jobs_v1_jobs_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineLimits) ProtoMessage() {}

func (x *MachineLimits) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_v1_jobs_proto_msgTypes[7]
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
	return file_jobs_v1_jobs_proto_rawDescGZIP(), []int{7}
}

func (x *MachineLimits) GetCPUCores() float32 {
	if x != nil {
		return x.CPUCores
	}
	return 0
}

func (x *MachineLimits) GetMemoryInMb() int64 {
	if x != nil {
		return x.MemoryInMb
	}
	return 0
}

func (x *MachineLimits) GetPidLimit() int64 {
	if x != nil {
		return x.PidLimit
	}
	return 0
}

var File_jobs_v1_jobs_proto protoreflect.FileDescriptor

var file_jobs_v1_jobs_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x6d,
	0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x42, 0x0a,
	0x11, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x11,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x13, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6d, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6d, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x06, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x22, 0x26, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x4a, 0x6f,
	0x62, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x64, 0x22, 0x53, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x9d, 0x02, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x13, 0x74, 0x6d, 0x70, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x6d, 0x70, 0x4a, 0x6f, 0x62, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x67, 0x0a, 0x0d, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x50, 0x55, 0x43, 0x6f,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x43, 0x50, 0x55, 0x43, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x4d,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x4d, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x69, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x69, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x32,
	0xd5, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x44, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x12, 0x19,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8c, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4a, 0x6f, 0x62, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6c, 0x65, 0x76, 0x69, 0x61, 0x74, 0x68, 0x61, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4a, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x4a, 0x6f, 0x62, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4a, 0x6f,
	0x62, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4a, 0x6f, 0x62, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4a, 0x6f,
	0x62, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_jobs_v1_jobs_proto_rawDescOnce sync.Once
	file_jobs_v1_jobs_proto_rawDescData []byte
)

func file_jobs_v1_jobs_proto_rawDescGZIP() []byte {
	file_jobs_v1_jobs_proto_rawDescOnce.Do(func() {
		file_jobs_v1_jobs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_jobs_v1_jobs_proto_rawDesc), len(file_jobs_v1_jobs_proto_rawDesc)))
	})
	return file_jobs_v1_jobs_proto_rawDescData
}

var file_jobs_v1_jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_jobs_v1_jobs_proto_goTypes = []any{
	(*NewJobRequest)(nil),     // 0: jobs.v1.NewJobRequest
	(*NewJobResponse)(nil),    // 1: jobs.v1.NewJobResponse
	(*CancelJobRequest)(nil),  // 2: jobs.v1.CancelJobRequest
	(*CancelJobResponse)(nil), // 3: jobs.v1.CancelJobResponse
	(*JobLogRequest)(nil),     // 4: jobs.v1.JobLogRequest
	(*JobLogsResponse)(nil),   // 5: jobs.v1.JobLogsResponse
	(*JobStatus)(nil),         // 6: jobs.v1.JobStatus
	(*MachineLimits)(nil),     // 7: jobs.v1.MachineLimits
	(*v1.FileUpload)(nil),     // 8: types.v1.FileUpload
}
var file_jobs_v1_jobs_proto_depIdxs = []int32{
	8, // 0: jobs.v1.NewJobRequest.makeFile:type_name -> types.v1.FileUpload
	8, // 1: jobs.v1.NewJobRequest.graderFile:type_name -> types.v1.FileUpload
	8, // 2: jobs.v1.NewJobRequest.studentSubmission:type_name -> types.v1.FileUpload
	8, // 3: jobs.v1.NewJobRequest.dockerFile:type_name -> types.v1.FileUpload
	7, // 4: jobs.v1.NewJobRequest.limits:type_name -> jobs.v1.MachineLimits
	6, // 5: jobs.v1.JobLogsResponse.jobInfo:type_name -> jobs.v1.JobStatus
	0, // 6: jobs.v1.JobService.NewJob:input_type -> jobs.v1.NewJobRequest
	4, // 7: jobs.v1.JobService.StreamStatus:input_type -> jobs.v1.JobLogRequest
	2, // 8: jobs.v1.JobService.CancelJob:input_type -> jobs.v1.CancelJobRequest
	1, // 9: jobs.v1.JobService.NewJob:output_type -> jobs.v1.NewJobResponse
	5, // 10: jobs.v1.JobService.StreamStatus:output_type -> jobs.v1.JobLogsResponse
	3, // 11: jobs.v1.JobService.CancelJob:output_type -> jobs.v1.CancelJobResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_jobs_v1_jobs_proto_init() }
func file_jobs_v1_jobs_proto_init() {
	if File_jobs_v1_jobs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_jobs_v1_jobs_proto_rawDesc), len(file_jobs_v1_jobs_proto_rawDesc)),
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
	file_jobs_v1_jobs_proto_goTypes = nil
	file_jobs_v1_jobs_proto_depIdxs = nil
}
