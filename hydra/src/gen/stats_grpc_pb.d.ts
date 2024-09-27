// package: stats.V1
// file: stats.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stats_pb from "./stats_pb";

interface IStatsServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    echo: IStatsServiceService_IEcho;
}

interface IStatsServiceService_IEcho extends grpc.MethodDefinition<stats_pb.EchoRequest, stats_pb.EchoResponse> {
    path: "/stats.V1.StatsService/Echo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stats_pb.EchoRequest>;
    requestDeserialize: grpc.deserialize<stats_pb.EchoRequest>;
    responseSerialize: grpc.serialize<stats_pb.EchoResponse>;
    responseDeserialize: grpc.deserialize<stats_pb.EchoResponse>;
}

export const StatsServiceService: IStatsServiceService;

export interface IStatsServiceServer {
    echo: grpc.handleUnaryCall<stats_pb.EchoRequest, stats_pb.EchoResponse>;
}

export interface IStatsServiceClient {
    echo(request: stats_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: stats_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: stats_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stats_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: stats_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stats_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}

export class StatsServiceClient extends grpc.Client implements IStatsServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public echo(request: stats_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: stats_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: stats_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stats_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: stats_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stats_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}
