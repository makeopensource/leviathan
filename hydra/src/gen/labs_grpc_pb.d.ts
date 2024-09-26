// package: labs.V1
// file: labs.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as labs_pb from "./labs_pb";

interface ILabServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    newLab: ILabServiceService_INewLab;
    editLab: ILabServiceService_IEditLab;
    deleteLab: ILabServiceService_IDeleteLab;
    echo: ILabServiceService_IEcho;
}

interface ILabServiceService_INewLab extends grpc.MethodDefinition<labs_pb.NewLabRequest, labs_pb.NewLabResponse> {
    path: "/labs.V1.LabService/NewLab";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<labs_pb.NewLabRequest>;
    requestDeserialize: grpc.deserialize<labs_pb.NewLabRequest>;
    responseSerialize: grpc.serialize<labs_pb.NewLabResponse>;
    responseDeserialize: grpc.deserialize<labs_pb.NewLabResponse>;
}
interface ILabServiceService_IEditLab extends grpc.MethodDefinition<labs_pb.EditLabRequest, labs_pb.EditLabResponse> {
    path: "/labs.V1.LabService/EditLab";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<labs_pb.EditLabRequest>;
    requestDeserialize: grpc.deserialize<labs_pb.EditLabRequest>;
    responseSerialize: grpc.serialize<labs_pb.EditLabResponse>;
    responseDeserialize: grpc.deserialize<labs_pb.EditLabResponse>;
}
interface ILabServiceService_IDeleteLab extends grpc.MethodDefinition<labs_pb.DeleteLabRequest, labs_pb.DeleteLabResponse> {
    path: "/labs.V1.LabService/DeleteLab";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<labs_pb.DeleteLabRequest>;
    requestDeserialize: grpc.deserialize<labs_pb.DeleteLabRequest>;
    responseSerialize: grpc.serialize<labs_pb.DeleteLabResponse>;
    responseDeserialize: grpc.deserialize<labs_pb.DeleteLabResponse>;
}
interface ILabServiceService_IEcho extends grpc.MethodDefinition<labs_pb.EchoRequest, labs_pb.EchoResponse> {
    path: "/labs.V1.LabService/Echo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<labs_pb.EchoRequest>;
    requestDeserialize: grpc.deserialize<labs_pb.EchoRequest>;
    responseSerialize: grpc.serialize<labs_pb.EchoResponse>;
    responseDeserialize: grpc.deserialize<labs_pb.EchoResponse>;
}

export const LabServiceService: ILabServiceService;

export interface ILabServiceServer {
    newLab: grpc.handleUnaryCall<labs_pb.NewLabRequest, labs_pb.NewLabResponse>;
    editLab: grpc.handleUnaryCall<labs_pb.EditLabRequest, labs_pb.EditLabResponse>;
    deleteLab: grpc.handleUnaryCall<labs_pb.DeleteLabRequest, labs_pb.DeleteLabResponse>;
    echo: grpc.handleUnaryCall<labs_pb.EchoRequest, labs_pb.EchoResponse>;
}

export interface ILabServiceClient {
    newLab(request: labs_pb.NewLabRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.NewLabResponse) => void): grpc.ClientUnaryCall;
    newLab(request: labs_pb.NewLabRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.NewLabResponse) => void): grpc.ClientUnaryCall;
    newLab(request: labs_pb.NewLabRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.NewLabResponse) => void): grpc.ClientUnaryCall;
    editLab(request: labs_pb.EditLabRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.EditLabResponse) => void): grpc.ClientUnaryCall;
    editLab(request: labs_pb.EditLabRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.EditLabResponse) => void): grpc.ClientUnaryCall;
    editLab(request: labs_pb.EditLabRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.EditLabResponse) => void): grpc.ClientUnaryCall;
    deleteLab(request: labs_pb.DeleteLabRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.DeleteLabResponse) => void): grpc.ClientUnaryCall;
    deleteLab(request: labs_pb.DeleteLabRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.DeleteLabResponse) => void): grpc.ClientUnaryCall;
    deleteLab(request: labs_pb.DeleteLabRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.DeleteLabResponse) => void): grpc.ClientUnaryCall;
    echo(request: labs_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: labs_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: labs_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}

export class LabServiceClient extends grpc.Client implements ILabServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public newLab(request: labs_pb.NewLabRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.NewLabResponse) => void): grpc.ClientUnaryCall;
    public newLab(request: labs_pb.NewLabRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.NewLabResponse) => void): grpc.ClientUnaryCall;
    public newLab(request: labs_pb.NewLabRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.NewLabResponse) => void): grpc.ClientUnaryCall;
    public editLab(request: labs_pb.EditLabRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.EditLabResponse) => void): grpc.ClientUnaryCall;
    public editLab(request: labs_pb.EditLabRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.EditLabResponse) => void): grpc.ClientUnaryCall;
    public editLab(request: labs_pb.EditLabRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.EditLabResponse) => void): grpc.ClientUnaryCall;
    public deleteLab(request: labs_pb.DeleteLabRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.DeleteLabResponse) => void): grpc.ClientUnaryCall;
    public deleteLab(request: labs_pb.DeleteLabRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.DeleteLabResponse) => void): grpc.ClientUnaryCall;
    public deleteLab(request: labs_pb.DeleteLabRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.DeleteLabResponse) => void): grpc.ClientUnaryCall;
    public echo(request: labs_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: labs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: labs_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: labs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: labs_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: labs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}
