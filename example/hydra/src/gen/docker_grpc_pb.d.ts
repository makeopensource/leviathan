// package: docker_rpc.V1
// file: docker.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as docker_pb from "./docker_pb";

interface IDockerServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createContainer: IDockerServiceService_ICreateContainer;
    deleteContainer: IDockerServiceService_IDeleteContainer;
    listContainers: IDockerServiceService_IListContainers;
    echo: IDockerServiceService_IEcho;
}

interface IDockerServiceService_ICreateContainer extends grpc.MethodDefinition<docker_pb.CreateRequest, docker_pb.CreateResponse> {
    path: "/docker_rpc.V1.DockerService/CreateContainer";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<docker_pb.CreateRequest>;
    requestDeserialize: grpc.deserialize<docker_pb.CreateRequest>;
    responseSerialize: grpc.serialize<docker_pb.CreateResponse>;
    responseDeserialize: grpc.deserialize<docker_pb.CreateResponse>;
}
interface IDockerServiceService_IDeleteContainer extends grpc.MethodDefinition<docker_pb.DeleteRequest, docker_pb.DeleteResponse> {
    path: "/docker_rpc.V1.DockerService/DeleteContainer";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<docker_pb.DeleteRequest>;
    requestDeserialize: grpc.deserialize<docker_pb.DeleteRequest>;
    responseSerialize: grpc.serialize<docker_pb.DeleteResponse>;
    responseDeserialize: grpc.deserialize<docker_pb.DeleteResponse>;
}
interface IDockerServiceService_IListContainers extends grpc.MethodDefinition<docker_pb.ListRequest, docker_pb.ListResponse> {
    path: "/docker_rpc.V1.DockerService/ListContainers";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<docker_pb.ListRequest>;
    requestDeserialize: grpc.deserialize<docker_pb.ListRequest>;
    responseSerialize: grpc.serialize<docker_pb.ListResponse>;
    responseDeserialize: grpc.deserialize<docker_pb.ListResponse>;
}
interface IDockerServiceService_IEcho extends grpc.MethodDefinition<docker_pb.EchoRequest, docker_pb.EchoResponse> {
    path: "/docker_rpc.V1.DockerService/Echo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<docker_pb.EchoRequest>;
    requestDeserialize: grpc.deserialize<docker_pb.EchoRequest>;
    responseSerialize: grpc.serialize<docker_pb.EchoResponse>;
    responseDeserialize: grpc.deserialize<docker_pb.EchoResponse>;
}

export const DockerServiceService: IDockerServiceService;

export interface IDockerServiceServer {
    createContainer: grpc.handleUnaryCall<docker_pb.CreateRequest, docker_pb.CreateResponse>;
    deleteContainer: grpc.handleUnaryCall<docker_pb.DeleteRequest, docker_pb.DeleteResponse>;
    listContainers: grpc.handleUnaryCall<docker_pb.ListRequest, docker_pb.ListResponse>;
    echo: grpc.handleUnaryCall<docker_pb.EchoRequest, docker_pb.EchoResponse>;
}

export interface IDockerServiceClient {
    createContainer(request: docker_pb.CreateRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    createContainer(request: docker_pb.CreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    createContainer(request: docker_pb.CreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    deleteContainer(request: docker_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteContainer(request: docker_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteContainer(request: docker_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    listContainers(request: docker_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.ListResponse) => void): grpc.ClientUnaryCall;
    listContainers(request: docker_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.ListResponse) => void): grpc.ClientUnaryCall;
    listContainers(request: docker_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.ListResponse) => void): grpc.ClientUnaryCall;
    echo(request: docker_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: docker_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: docker_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}

export class DockerServiceClient extends grpc.Client implements IDockerServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createContainer(request: docker_pb.CreateRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public createContainer(request: docker_pb.CreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public createContainer(request: docker_pb.CreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public deleteContainer(request: docker_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteContainer(request: docker_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteContainer(request: docker_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public listContainers(request: docker_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public listContainers(request: docker_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public listContainers(request: docker_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public echo(request: docker_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: docker_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: docker_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: docker_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: docker_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: docker_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}
