// package: labs.V1
// file: labs.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class NewLabRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewLabRequest.AsObject;
    static toObject(includeInstance: boolean, msg: NewLabRequest): NewLabRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewLabRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewLabRequest;
    static deserializeBinaryFromReader(message: NewLabRequest, reader: jspb.BinaryReader): NewLabRequest;
}

export namespace NewLabRequest {
    export type AsObject = {
    }
}

export class NewLabResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewLabResponse.AsObject;
    static toObject(includeInstance: boolean, msg: NewLabResponse): NewLabResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewLabResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewLabResponse;
    static deserializeBinaryFromReader(message: NewLabResponse, reader: jspb.BinaryReader): NewLabResponse;
}

export namespace NewLabResponse {
    export type AsObject = {
    }
}

export class EditLabRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EditLabRequest.AsObject;
    static toObject(includeInstance: boolean, msg: EditLabRequest): EditLabRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EditLabRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EditLabRequest;
    static deserializeBinaryFromReader(message: EditLabRequest, reader: jspb.BinaryReader): EditLabRequest;
}

export namespace EditLabRequest {
    export type AsObject = {
    }
}

export class EditLabResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EditLabResponse.AsObject;
    static toObject(includeInstance: boolean, msg: EditLabResponse): EditLabResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EditLabResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EditLabResponse;
    static deserializeBinaryFromReader(message: EditLabResponse, reader: jspb.BinaryReader): EditLabResponse;
}

export namespace EditLabResponse {
    export type AsObject = {
    }
}

export class DeleteLabRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteLabRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteLabRequest): DeleteLabRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteLabRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteLabRequest;
    static deserializeBinaryFromReader(message: DeleteLabRequest, reader: jspb.BinaryReader): DeleteLabRequest;
}

export namespace DeleteLabRequest {
    export type AsObject = {
    }
}

export class DeleteLabResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteLabResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteLabResponse): DeleteLabResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteLabResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteLabResponse;
    static deserializeBinaryFromReader(message: DeleteLabResponse, reader: jspb.BinaryReader): DeleteLabResponse;
}

export namespace DeleteLabResponse {
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
