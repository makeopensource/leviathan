// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file jobs/v1/jobs.proto (package jobs.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { FileUpload } from "../../types/v1/types_pb";
import { file_types_v1_types } from "../../types/v1/types_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file jobs/v1/jobs.proto.
 */
export const file_jobs_v1_jobs: GenFile = /*@__PURE__*/
  fileDesc("ChJqb2JzL3YxL2pvYnMucHJvdG8SB2pvYnMudjEipgIKDU5ld0pvYlJlcXVlc3QSJgoIbWFrZUZpbGUYASABKAsyFC50eXBlcy52MS5GaWxlVXBsb2FkEigKCmdyYWRlckZpbGUYAiABKAsyFC50eXBlcy52MS5GaWxlVXBsb2FkEi8KEXN0dWRlbnRTdWJtaXNzaW9uGAMgASgLMhQudHlwZXMudjEuRmlsZVVwbG9hZBIoCgpkb2NrZXJGaWxlGAQgASgLMhQudHlwZXMudjEuRmlsZVVwbG9hZBIRCglpbWFnZU5hbWUYBSABKAkSGwoTam9iVGltZW91dEluU2Vjb25kcxgGIAEoBBIQCghlbnRyeUNtZBgHIAEoCRImCgZsaW1pdHMYCCABKAsyFi5qb2JzLnYxLk1hY2hpbmVMaW1pdHMiHwoOTmV3Sm9iUmVzcG9uc2USDQoFam9iSWQYASABKAkiIQoQQ2FuY2VsSm9iUmVxdWVzdBINCgVqb2JJZBgBIAEoCSITChFDYW5jZWxKb2JSZXNwb25zZSIeCg1Kb2JMb2dSZXF1ZXN0Eg0KBWpvYklkGAEgASgJIkQKD0pvYkxvZ3NSZXNwb25zZRIjCgdqb2JJbmZvGAEgASgLMhIuam9icy52MS5Kb2JTdGF0dXMSDAoEbG9ncxgCIAEoCSK5AQoJSm9iU3RhdHVzEg4KBmpvYl9pZBgBIAEoCRISCgptYWNoaW5lX2lkGAIgASgJEhQKDGNvbnRhaW5lcl9pZBgDIAEoCRIOCgZzdGF0dXMYBCABKAkSFgoOc3RhdHVzX21lc3NhZ2UYBSABKAkSGAoQb3V0cHV0X2ZpbGVfcGF0aBgGIAEoCRIbChN0bXBfam9iX2ZvbGRlcl9wYXRoGAcgASgJEhMKC2pvYl90aW1lb3V0GAggASgDIkcKDU1hY2hpbmVMaW1pdHMSEAoIQ1BVQ29yZXMYASABKAISEgoKbWVtb3J5SW5NYhgCIAEoAxIQCghQaWRMaW1pdBgDIAEoAzLVAQoKSm9iU2VydmljZRI7CgZOZXdKb2ISFi5qb2JzLnYxLk5ld0pvYlJlcXVlc3QaFy5qb2JzLnYxLk5ld0pvYlJlc3BvbnNlIgASRAoMU3RyZWFtU3RhdHVzEhYuam9icy52MS5Kb2JMb2dSZXF1ZXN0Ghguam9icy52MS5Kb2JMb2dzUmVzcG9uc2UiADABEkQKCUNhbmNlbEpvYhIZLmpvYnMudjEuQ2FuY2VsSm9iUmVxdWVzdBoaLmpvYnMudjEuQ2FuY2VsSm9iUmVzcG9uc2UiAEKMAQoLY29tLmpvYnMudjFCCUpvYnNQcm90b1ABWjVnaXRodWIuY29tL21ha2VvcGVuc291cmNlL2xldmlhdGhhbi9nZW5lcmF0ZWQvam9icy92MaICA0pYWKoCB0pvYnMuVjHKAgdKb2JzXFYx4gITSm9ic1xWMVxHUEJNZXRhZGF0YeoCCEpvYnM6OlYxYgZwcm90bzM", [file_types_v1_types]);

/**
 * @generated from message jobs.v1.NewJobRequest
 */
export type NewJobRequest = Message<"jobs.v1.NewJobRequest"> & {
  /**
   * @generated from field: types.v1.FileUpload makeFile = 1;
   */
  makeFile?: FileUpload;

  /**
   * @generated from field: types.v1.FileUpload graderFile = 2;
   */
  graderFile?: FileUpload;

  /**
   * @generated from field: types.v1.FileUpload studentSubmission = 3;
   */
  studentSubmission?: FileUpload;

  /**
   * @generated from field: types.v1.FileUpload dockerFile = 4;
   */
  dockerFile?: FileUpload;

  /**
   * @generated from field: string imageName = 5;
   */
  imageName: string;

  /**
   * @generated from field: uint64 jobTimeoutInSeconds = 6;
   */
  jobTimeoutInSeconds: bigint;

  /**
   * @generated from field: string entryCmd = 7;
   */
  entryCmd: string;

  /**
   * @generated from field: jobs.v1.MachineLimits limits = 8;
   */
  limits?: MachineLimits;
};

/**
 * Describes the message jobs.v1.NewJobRequest.
 * Use `create(NewJobRequestSchema)` to create a new message.
 */
export const NewJobRequestSchema: GenMessage<NewJobRequest> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 0);

/**
 * @generated from message jobs.v1.NewJobResponse
 */
export type NewJobResponse = Message<"jobs.v1.NewJobResponse"> & {
  /**
   * @generated from field: string jobId = 1;
   */
  jobId: string;
};

/**
 * Describes the message jobs.v1.NewJobResponse.
 * Use `create(NewJobResponseSchema)` to create a new message.
 */
export const NewJobResponseSchema: GenMessage<NewJobResponse> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 1);

/**
 * @generated from message jobs.v1.CancelJobRequest
 */
export type CancelJobRequest = Message<"jobs.v1.CancelJobRequest"> & {
  /**
   * @generated from field: string jobId = 1;
   */
  jobId: string;
};

/**
 * Describes the message jobs.v1.CancelJobRequest.
 * Use `create(CancelJobRequestSchema)` to create a new message.
 */
export const CancelJobRequestSchema: GenMessage<CancelJobRequest> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 2);

/**
 * @generated from message jobs.v1.CancelJobResponse
 */
export type CancelJobResponse = Message<"jobs.v1.CancelJobResponse"> & {
};

/**
 * Describes the message jobs.v1.CancelJobResponse.
 * Use `create(CancelJobResponseSchema)` to create a new message.
 */
export const CancelJobResponseSchema: GenMessage<CancelJobResponse> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 3);

/**
 * @generated from message jobs.v1.JobLogRequest
 */
export type JobLogRequest = Message<"jobs.v1.JobLogRequest"> & {
  /**
   * @generated from field: string jobId = 1;
   */
  jobId: string;
};

/**
 * Describes the message jobs.v1.JobLogRequest.
 * Use `create(JobLogRequestSchema)` to create a new message.
 */
export const JobLogRequestSchema: GenMessage<JobLogRequest> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 4);

/**
 * @generated from message jobs.v1.JobLogsResponse
 */
export type JobLogsResponse = Message<"jobs.v1.JobLogsResponse"> & {
  /**
   * @generated from field: jobs.v1.JobStatus jobInfo = 1;
   */
  jobInfo?: JobStatus;

  /**
   * @generated from field: string logs = 2;
   */
  logs: string;
};

/**
 * Describes the message jobs.v1.JobLogsResponse.
 * Use `create(JobLogsResponseSchema)` to create a new message.
 */
export const JobLogsResponseSchema: GenMessage<JobLogsResponse> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 5);

/**
 * @generated from message jobs.v1.JobStatus
 */
export type JobStatus = Message<"jobs.v1.JobStatus"> & {
  /**
   * @generated from field: string job_id = 1;
   */
  jobId: string;

  /**
   * @generated from field: string machine_id = 2;
   */
  machineId: string;

  /**
   * @generated from field: string container_id = 3;
   */
  containerId: string;

  /**
   * @generated from field: string status = 4;
   */
  status: string;

  /**
   * @generated from field: string status_message = 5;
   */
  statusMessage: string;

  /**
   * @generated from field: string output_file_path = 6;
   */
  outputFilePath: string;

  /**
   * @generated from field: string tmp_job_folder_path = 7;
   */
  tmpJobFolderPath: string;

  /**
   * @generated from field: int64 job_timeout = 8;
   */
  jobTimeout: bigint;
};

/**
 * Describes the message jobs.v1.JobStatus.
 * Use `create(JobStatusSchema)` to create a new message.
 */
export const JobStatusSchema: GenMessage<JobStatus> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 6);

/**
 * @generated from message jobs.v1.MachineLimits
 */
export type MachineLimits = Message<"jobs.v1.MachineLimits"> & {
  /**
   * @generated from field: float CPUCores = 1;
   */
  CPUCores: number;

  /**
   * @generated from field: int64 memoryInMb = 2;
   */
  memoryInMb: bigint;

  /**
   * @generated from field: int64 PidLimit = 3;
   */
  PidLimit: bigint;
};

/**
 * Describes the message jobs.v1.MachineLimits.
 * Use `create(MachineLimitsSchema)` to create a new message.
 */
export const MachineLimitsSchema: GenMessage<MachineLimits> = /*@__PURE__*/
  messageDesc(file_jobs_v1_jobs, 7);

/**
 * @generated from service jobs.v1.JobService
 */
export const JobService: GenService<{
  /**
   * @generated from rpc jobs.v1.JobService.NewJob
   */
  newJob: {
    methodKind: "unary";
    input: typeof NewJobRequestSchema;
    output: typeof NewJobResponseSchema;
  },
  /**
   * @generated from rpc jobs.v1.JobService.StreamStatus
   */
  streamStatus: {
    methodKind: "server_streaming";
    input: typeof JobLogRequestSchema;
    output: typeof JobLogsResponseSchema;
  },
  /**
   * @generated from rpc jobs.v1.JobService.CancelJob
   */
  cancelJob: {
    methodKind: "unary";
    input: typeof CancelJobRequestSchema;
    output: typeof CancelJobResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_jobs_v1_jobs, 0);

