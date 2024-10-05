// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file stats/v1/stats.proto (package stats.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message stats.v1.EchoRequest
 */
export class EchoRequest extends Message<EchoRequest> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<EchoRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "stats.v1.EchoRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EchoRequest {
    return new EchoRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EchoRequest {
    return new EchoRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EchoRequest {
    return new EchoRequest().fromJsonString(jsonString, options);
  }

  static equals(a: EchoRequest | PlainMessage<EchoRequest> | undefined, b: EchoRequest | PlainMessage<EchoRequest> | undefined): boolean {
    return proto3.util.equals(EchoRequest, a, b);
  }
}

/**
 * @generated from message stats.v1.EchoResponse
 */
export class EchoResponse extends Message<EchoResponse> {
  /**
   * @generated from field: string message_response = 1;
   */
  messageResponse = "";

  constructor(data?: PartialMessage<EchoResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "stats.v1.EchoResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message_response", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EchoResponse {
    return new EchoResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EchoResponse {
    return new EchoResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EchoResponse {
    return new EchoResponse().fromJsonString(jsonString, options);
  }

  static equals(a: EchoResponse | PlainMessage<EchoResponse> | undefined, b: EchoResponse | PlainMessage<EchoResponse> | undefined): boolean {
    return proto3.util.equals(EchoResponse, a, b);
  }
}

