// package: jobs.V1
// file: jobs.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class NewJobRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewJobRequest.AsObject;
    static toObject(includeInstance: boolean, msg: NewJobRequest): NewJobRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewJobRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewJobRequest;
    static deserializeBinaryFromReader(message: NewJobRequest, reader: jspb.BinaryReader): NewJobRequest;
}

export namespace NewJobRequest {
    export type AsObject = {
    }
}

export class NewJobResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewJobResponse.AsObject;
    static toObject(includeInstance: boolean, msg: NewJobResponse): NewJobResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewJobResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewJobResponse;
    static deserializeBinaryFromReader(message: NewJobResponse, reader: jspb.BinaryReader): NewJobResponse;
}

export namespace NewJobResponse {
    export type AsObject = {
    }
}

export class JobStatusRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): JobStatusRequest.AsObject;
    static toObject(includeInstance: boolean, msg: JobStatusRequest): JobStatusRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: JobStatusRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): JobStatusRequest;
    static deserializeBinaryFromReader(message: JobStatusRequest, reader: jspb.BinaryReader): JobStatusRequest;
}

export namespace JobStatusRequest {
    export type AsObject = {
    }
}

export class JobStatusResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): JobStatusResponse.AsObject;
    static toObject(includeInstance: boolean, msg: JobStatusResponse): JobStatusResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: JobStatusResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): JobStatusResponse;
    static deserializeBinaryFromReader(message: JobStatusResponse, reader: jspb.BinaryReader): JobStatusResponse;
}

export namespace JobStatusResponse {
    export type AsObject = {
    }
}

export class CancelJobRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CancelJobRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CancelJobRequest): CancelJobRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CancelJobRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CancelJobRequest;
    static deserializeBinaryFromReader(message: CancelJobRequest, reader: jspb.BinaryReader): CancelJobRequest;
}

export namespace CancelJobRequest {
    export type AsObject = {
    }
}

export class CancelJobResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CancelJobResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CancelJobResponse): CancelJobResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CancelJobResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CancelJobResponse;
    static deserializeBinaryFromReader(message: CancelJobResponse, reader: jspb.BinaryReader): CancelJobResponse;
}

export namespace CancelJobResponse {
    export type AsObject = {
    }
}

export class EchoRequest extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): EchoRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EchoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: EchoRequest): EchoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EchoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EchoRequest;
    static deserializeBinaryFromReader(message: EchoRequest, reader: jspb.BinaryReader): EchoRequest;
}

export namespace EchoRequest {
    export type AsObject = {
        message: string,
    }
}

export class EchoResponse extends jspb.Message { 
    getMessageresponse(): string;
    setMessageresponse(value: string): EchoResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EchoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: EchoResponse): EchoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EchoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EchoResponse;
    static deserializeBinaryFromReader(message: EchoResponse, reader: jspb.BinaryReader): EchoResponse;
}

export namespace EchoResponse {
    export type AsObject = {
        messageresponse: string,
    }
}
