// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file jobs/v1/jobs.proto (package jobs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * todo figure out request/response
 *
 * @generated from message jobs.v1.NewJobRequest
 */
export class NewJobRequest extends Message<NewJobRequest> {
  constructor(data?: PartialMessage<NewJobRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "jobs.v1.NewJobRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NewJobRequest {
    return new NewJobRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NewJobRequest {
    return new NewJobRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NewJobRequest {
    return new NewJobRequest().fromJsonString(jsonString, options);
  }

  static equals(a: NewJobRequest | PlainMessage<NewJobRequest> | undefined, b: NewJobRequest | PlainMessage<NewJobRequest> | undefined): boolean {
    return proto3.util.equals(NewJobRequest, a, b);
  }
}

/**
 * @generated from message jobs.v1.NewJobResponse
 */
export class NewJobResponse extends Message<NewJobResponse> {
  constructor(data?: PartialMessage<NewJobResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "jobs.v1.NewJobResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NewJobResponse {
    return new NewJobResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NewJobResponse {
    return new NewJobResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NewJobResponse {
    return new NewJobResponse().fromJsonString(jsonString, options);
  }

  static equals(a: NewJobResponse | PlainMessage<NewJobResponse> | undefined, b: NewJobResponse | PlainMessage<NewJobResponse> | undefined): boolean {
    return proto3.util.equals(NewJobResponse, a, b);
  }
}

/**
 * @generated from message jobs.v1.JobStatusRequest
 */
export class JobStatusRequest extends Message<JobStatusRequest> {
  constructor(data?: PartialMessage<JobStatusRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "jobs.v1.JobStatusRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JobStatusRequest {
    return new JobStatusRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JobStatusRequest {
    return new JobStatusRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JobStatusRequest {
    return new JobStatusRequest().fromJsonString(jsonString, options);
  }

  static equals(a: JobStatusRequest | PlainMessage<JobStatusRequest> | undefined, b: JobStatusRequest | PlainMessage<JobStatusRequest> | undefined): boolean {
    return proto3.util.equals(JobStatusRequest, a, b);
  }
}

/**
 * @generated from message jobs.v1.JobStatusResponse
 */
export class JobStatusResponse extends Message<JobStatusResponse> {
  constructor(data?: PartialMessage<JobStatusResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "jobs.v1.JobStatusResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JobStatusResponse {
    return new JobStatusResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JobStatusResponse {
    return new JobStatusResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JobStatusResponse {
    return new JobStatusResponse().fromJsonString(jsonString, options);
  }

  static equals(a: JobStatusResponse | PlainMessage<JobStatusResponse> | undefined, b: JobStatusResponse | PlainMessage<JobStatusResponse> | undefined): boolean {
    return proto3.util.equals(JobStatusResponse, a, b);
  }
}

/**
 * You can add filters here if needed
 *
 * @generated from message jobs.v1.CancelJobRequest
 */
export class CancelJobRequest extends Message<CancelJobRequest> {
  constructor(data?: PartialMessage<CancelJobRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "jobs.v1.CancelJobRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CancelJobRequest {
    return new CancelJobRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CancelJobRequest {
    return new CancelJobRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CancelJobRequest {
    return new CancelJobRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CancelJobRequest | PlainMessage<CancelJobRequest> | undefined, b: CancelJobRequest | PlainMessage<CancelJobRequest> | undefined): boolean {
    return proto3.util.equals(CancelJobRequest, a, b);
  }
}

/**
 * @generated from message jobs.v1.CancelJobResponse
 */
export class CancelJobResponse extends Message<CancelJobResponse> {
  constructor(data?: PartialMessage<CancelJobResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "jobs.v1.CancelJobResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CancelJobResponse {
    return new CancelJobResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CancelJobResponse {
    return new CancelJobResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CancelJobResponse {
    return new CancelJobResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CancelJobResponse | PlainMessage<CancelJobResponse> | undefined, b: CancelJobResponse | PlainMessage<CancelJobResponse> | undefined): boolean {
    return proto3.util.equals(CancelJobResponse, a, b);
  }
}

