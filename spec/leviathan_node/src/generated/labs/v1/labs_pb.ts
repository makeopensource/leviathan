// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file labs/v1/labs.proto (package labs.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { FileUpload, MachineLimits } from "../../types/v1/types_pb";
import { file_types_v1_types } from "../../types/v1/types_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file labs/v1/labs.proto.
 */
export const file_labs_v1_labs: GenFile = /*@__PURE__*/
  fileDesc("ChJsYWJzL3YxL2xhYnMucHJvdG8SB2xhYnMudjEi0QEKCkxhYlJlcXVlc3QSDwoHbGFiTmFtZRgBIAEoCRIUCgxlbnRyeUNvbW1hbmQYAiABKAkSGgoSdGltZUxpbWl0SW5TZWNvbmRzGAQgASgEEi4KDW1hY2hpbmVMaW1pdHMYBSABKAsyFy50eXBlcy52MS5NYWNoaW5lTGltaXRzEigKCmRvY2tlckZpbGUYBiABKAsyFC50eXBlcy52MS5GaWxlVXBsb2FkEiYKCGpvYkZpbGVzGAcgAygLMhQudHlwZXMudjEuRmlsZVVwbG9hZCIfCg5OZXdMYWJSZXNwb25zZRINCgVsYWJJZBgBIAEoAyJFCg5FZGl0TGFiUmVxdWVzdBINCgVsYWJJZBgBIAEoAxIkCgdsYWJJbmZvGAIgASgLMhMubGFicy52MS5MYWJSZXF1ZXN0IhEKD0VkaXRMYWJSZXNwb25zZSIhChBEZWxldGVMYWJSZXF1ZXN0Eg0KBUxhYklEGAEgASgDIhMKEURlbGV0ZUxhYlJlc3BvbnNlMsgBCgpMYWJTZXJ2aWNlEjgKBk5ld0xhYhITLmxhYnMudjEuTGFiUmVxdWVzdBoXLmxhYnMudjEuTmV3TGFiUmVzcG9uc2UiABI6CgdFZGl0TGFiEhMubGFicy52MS5MYWJSZXF1ZXN0GhgubGFicy52MS5FZGl0TGFiUmVzcG9uc2UiABJECglEZWxldGVMYWISGS5sYWJzLnYxLkRlbGV0ZUxhYlJlcXVlc3QaGi5sYWJzLnYxLkRlbGV0ZUxhYlJlc3BvbnNlIgBCjAEKC2NvbS5sYWJzLnYxQglMYWJzUHJvdG9QAVo1Z2l0aHViLmNvbS9tYWtlb3BlbnNvdXJjZS9sZXZpYXRoYW4vZ2VuZXJhdGVkL2xhYnMvdjGiAgNMWFiqAgdMYWJzLlYxygIHTGFic1xWMeICE0xhYnNcVjFcR1BCTWV0YWRhdGHqAghMYWJzOjpWMWIGcHJvdG8z", [file_types_v1_types]);

/**
 * @generated from message labs.v1.LabRequest
 */
export type LabRequest = Message<"labs.v1.LabRequest"> & {
  /**
   * @generated from field: string labName = 1;
   */
  labName: string;

  /**
   * @generated from field: string entryCommand = 2;
   */
  entryCommand: string;

  /**
   * @generated from field: uint64 timeLimitInSeconds = 4;
   */
  timeLimitInSeconds: bigint;

  /**
   * @generated from field: types.v1.MachineLimits machineLimits = 5;
   */
  machineLimits?: MachineLimits;

  /**
   * @generated from field: types.v1.FileUpload dockerFile = 6;
   */
  dockerFile?: FileUpload;

  /**
   * @generated from field: repeated types.v1.FileUpload jobFiles = 7;
   */
  jobFiles: FileUpload[];
};

/**
 * Describes the message labs.v1.LabRequest.
 * Use `create(LabRequestSchema)` to create a new message.
 */
export const LabRequestSchema: GenMessage<LabRequest> = /*@__PURE__*/
  messageDesc(file_labs_v1_labs, 0);

/**
 * @generated from message labs.v1.NewLabResponse
 */
export type NewLabResponse = Message<"labs.v1.NewLabResponse"> & {
  /**
   * @generated from field: int64 labId = 1;
   */
  labId: bigint;
};

/**
 * Describes the message labs.v1.NewLabResponse.
 * Use `create(NewLabResponseSchema)` to create a new message.
 */
export const NewLabResponseSchema: GenMessage<NewLabResponse> = /*@__PURE__*/
  messageDesc(file_labs_v1_labs, 1);

/**
 * @generated from message labs.v1.EditLabRequest
 */
export type EditLabRequest = Message<"labs.v1.EditLabRequest"> & {
  /**
   * @generated from field: int64 labId = 1;
   */
  labId: bigint;

  /**
   * @generated from field: labs.v1.LabRequest labInfo = 2;
   */
  labInfo?: LabRequest;
};

/**
 * Describes the message labs.v1.EditLabRequest.
 * Use `create(EditLabRequestSchema)` to create a new message.
 */
export const EditLabRequestSchema: GenMessage<EditLabRequest> = /*@__PURE__*/
  messageDesc(file_labs_v1_labs, 2);

/**
 * @generated from message labs.v1.EditLabResponse
 */
export type EditLabResponse = Message<"labs.v1.EditLabResponse"> & {
};

/**
 * Describes the message labs.v1.EditLabResponse.
 * Use `create(EditLabResponseSchema)` to create a new message.
 */
export const EditLabResponseSchema: GenMessage<EditLabResponse> = /*@__PURE__*/
  messageDesc(file_labs_v1_labs, 3);

/**
 * @generated from message labs.v1.DeleteLabRequest
 */
export type DeleteLabRequest = Message<"labs.v1.DeleteLabRequest"> & {
  /**
   * @generated from field: int64 LabID = 1;
   */
  LabID: bigint;
};

/**
 * Describes the message labs.v1.DeleteLabRequest.
 * Use `create(DeleteLabRequestSchema)` to create a new message.
 */
export const DeleteLabRequestSchema: GenMessage<DeleteLabRequest> = /*@__PURE__*/
  messageDesc(file_labs_v1_labs, 4);

/**
 * @generated from message labs.v1.DeleteLabResponse
 */
export type DeleteLabResponse = Message<"labs.v1.DeleteLabResponse"> & {
};

/**
 * Describes the message labs.v1.DeleteLabResponse.
 * Use `create(DeleteLabResponseSchema)` to create a new message.
 */
export const DeleteLabResponseSchema: GenMessage<DeleteLabResponse> = /*@__PURE__*/
  messageDesc(file_labs_v1_labs, 5);

/**
 * @generated from service labs.v1.LabService
 */
export const LabService: GenService<{
  /**
   * @generated from rpc labs.v1.LabService.NewLab
   */
  newLab: {
    methodKind: "unary";
    input: typeof LabRequestSchema;
    output: typeof NewLabResponseSchema;
  },
  /**
   * @generated from rpc labs.v1.LabService.EditLab
   */
  editLab: {
    methodKind: "unary";
    input: typeof LabRequestSchema;
    output: typeof EditLabResponseSchema;
  },
  /**
   * @generated from rpc labs.v1.LabService.DeleteLab
   */
  deleteLab: {
    methodKind: "unary";
    input: typeof DeleteLabRequestSchema;
    output: typeof DeleteLabResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_labs_v1_labs, 0);

