// package: jobs.V1
// file: jobs.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as jobs_pb from "./jobs_pb";

interface IJobServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    newJob: IJobServiceService_INewJob;
    jobStatus: IJobServiceService_IJobStatus;
    cancelJob: IJobServiceService_ICancelJob;
    echo: IJobServiceService_IEcho;
}

interface IJobServiceService_INewJob extends grpc.MethodDefinition<jobs_pb.NewJobRequest, jobs_pb.NewJobResponse> {
    path: "/jobs.V1.JobService/NewJob";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<jobs_pb.NewJobRequest>;
    requestDeserialize: grpc.deserialize<jobs_pb.NewJobRequest>;
    responseSerialize: grpc.serialize<jobs_pb.NewJobResponse>;
    responseDeserialize: grpc.deserialize<jobs_pb.NewJobResponse>;
}
interface IJobServiceService_IJobStatus extends grpc.MethodDefinition<jobs_pb.JobStatusRequest, jobs_pb.JobStatusResponse> {
    path: "/jobs.V1.JobService/JobStatus";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<jobs_pb.JobStatusRequest>;
    requestDeserialize: grpc.deserialize<jobs_pb.JobStatusRequest>;
    responseSerialize: grpc.serialize<jobs_pb.JobStatusResponse>;
    responseDeserialize: grpc.deserialize<jobs_pb.JobStatusResponse>;
}
interface IJobServiceService_ICancelJob extends grpc.MethodDefinition<jobs_pb.CancelJobRequest, jobs_pb.CancelJobResponse> {
    path: "/jobs.V1.JobService/CancelJob";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<jobs_pb.CancelJobRequest>;
    requestDeserialize: grpc.deserialize<jobs_pb.CancelJobRequest>;
    responseSerialize: grpc.serialize<jobs_pb.CancelJobResponse>;
    responseDeserialize: grpc.deserialize<jobs_pb.CancelJobResponse>;
}
interface IJobServiceService_IEcho extends grpc.MethodDefinition<jobs_pb.EchoRequest, jobs_pb.EchoResponse> {
    path: "/jobs.V1.JobService/Echo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<jobs_pb.EchoRequest>;
    requestDeserialize: grpc.deserialize<jobs_pb.EchoRequest>;
    responseSerialize: grpc.serialize<jobs_pb.EchoResponse>;
    responseDeserialize: grpc.deserialize<jobs_pb.EchoResponse>;
}

export const JobServiceService: IJobServiceService;

export interface IJobServiceServer {
    newJob: grpc.handleUnaryCall<jobs_pb.NewJobRequest, jobs_pb.NewJobResponse>;
    jobStatus: grpc.handleUnaryCall<jobs_pb.JobStatusRequest, jobs_pb.JobStatusResponse>;
    cancelJob: grpc.handleUnaryCall<jobs_pb.CancelJobRequest, jobs_pb.CancelJobResponse>;
    echo: grpc.handleUnaryCall<jobs_pb.EchoRequest, jobs_pb.EchoResponse>;
}

export interface IJobServiceClient {
    newJob(request: jobs_pb.NewJobRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.NewJobResponse) => void): grpc.ClientUnaryCall;
    newJob(request: jobs_pb.NewJobRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.NewJobResponse) => void): grpc.ClientUnaryCall;
    newJob(request: jobs_pb.NewJobRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.NewJobResponse) => void): grpc.ClientUnaryCall;
    jobStatus(request: jobs_pb.JobStatusRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.JobStatusResponse) => void): grpc.ClientUnaryCall;
    jobStatus(request: jobs_pb.JobStatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.JobStatusResponse) => void): grpc.ClientUnaryCall;
    jobStatus(request: jobs_pb.JobStatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.JobStatusResponse) => void): grpc.ClientUnaryCall;
    cancelJob(request: jobs_pb.CancelJobRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.CancelJobResponse) => void): grpc.ClientUnaryCall;
    cancelJob(request: jobs_pb.CancelJobRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.CancelJobResponse) => void): grpc.ClientUnaryCall;
    cancelJob(request: jobs_pb.CancelJobRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.CancelJobResponse) => void): grpc.ClientUnaryCall;
    echo(request: jobs_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: jobs_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    echo(request: jobs_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}

export class JobServiceClient extends grpc.Client implements IJobServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public newJob(request: jobs_pb.NewJobRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.NewJobResponse) => void): grpc.ClientUnaryCall;
    public newJob(request: jobs_pb.NewJobRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.NewJobResponse) => void): grpc.ClientUnaryCall;
    public newJob(request: jobs_pb.NewJobRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.NewJobResponse) => void): grpc.ClientUnaryCall;
    public jobStatus(request: jobs_pb.JobStatusRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.JobStatusResponse) => void): grpc.ClientUnaryCall;
    public jobStatus(request: jobs_pb.JobStatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.JobStatusResponse) => void): grpc.ClientUnaryCall;
    public jobStatus(request: jobs_pb.JobStatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.JobStatusResponse) => void): grpc.ClientUnaryCall;
    public cancelJob(request: jobs_pb.CancelJobRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.CancelJobResponse) => void): grpc.ClientUnaryCall;
    public cancelJob(request: jobs_pb.CancelJobRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.CancelJobResponse) => void): grpc.ClientUnaryCall;
    public cancelJob(request: jobs_pb.CancelJobRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.CancelJobResponse) => void): grpc.ClientUnaryCall;
    public echo(request: jobs_pb.EchoRequest, callback: (error: grpc.ServiceError | null, response: jobs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: jobs_pb.EchoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jobs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
    public echo(request: jobs_pb.EchoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jobs_pb.EchoResponse) => void): grpc.ClientUnaryCall;
}
