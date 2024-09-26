// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var jobs_pb = require('./jobs_pb.js');

function serialize_jobs_V1_CancelJobRequest(arg) {
  if (!(arg instanceof jobs_pb.CancelJobRequest)) {
    throw new Error('Expected argument of type jobs.V1.CancelJobRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_CancelJobRequest(buffer_arg) {
  return jobs_pb.CancelJobRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_CancelJobResponse(arg) {
  if (!(arg instanceof jobs_pb.CancelJobResponse)) {
    throw new Error('Expected argument of type jobs.V1.CancelJobResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_CancelJobResponse(buffer_arg) {
  return jobs_pb.CancelJobResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_EchoRequest(arg) {
  if (!(arg instanceof jobs_pb.EchoRequest)) {
    throw new Error('Expected argument of type jobs.V1.EchoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_EchoRequest(buffer_arg) {
  return jobs_pb.EchoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_EchoResponse(arg) {
  if (!(arg instanceof jobs_pb.EchoResponse)) {
    throw new Error('Expected argument of type jobs.V1.EchoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_EchoResponse(buffer_arg) {
  return jobs_pb.EchoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_JobStatusRequest(arg) {
  if (!(arg instanceof jobs_pb.JobStatusRequest)) {
    throw new Error('Expected argument of type jobs.V1.JobStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_JobStatusRequest(buffer_arg) {
  return jobs_pb.JobStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_JobStatusResponse(arg) {
  if (!(arg instanceof jobs_pb.JobStatusResponse)) {
    throw new Error('Expected argument of type jobs.V1.JobStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_JobStatusResponse(buffer_arg) {
  return jobs_pb.JobStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_NewJobRequest(arg) {
  if (!(arg instanceof jobs_pb.NewJobRequest)) {
    throw new Error('Expected argument of type jobs.V1.NewJobRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_NewJobRequest(buffer_arg) {
  return jobs_pb.NewJobRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jobs_V1_NewJobResponse(arg) {
  if (!(arg instanceof jobs_pb.NewJobResponse)) {
    throw new Error('Expected argument of type jobs.V1.NewJobResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jobs_V1_NewJobResponse(buffer_arg) {
  return jobs_pb.NewJobResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// import 'types.proto';
//
var JobServiceService = exports.JobServiceService = {
  newJob: {
    path: '/jobs.V1.JobService/NewJob',
    requestStream: false,
    responseStream: false,
    requestType: jobs_pb.NewJobRequest,
    responseType: jobs_pb.NewJobResponse,
    requestSerialize: serialize_jobs_V1_NewJobRequest,
    requestDeserialize: deserialize_jobs_V1_NewJobRequest,
    responseSerialize: serialize_jobs_V1_NewJobResponse,
    responseDeserialize: deserialize_jobs_V1_NewJobResponse,
  },
  jobStatus: {
    path: '/jobs.V1.JobService/JobStatus',
    requestStream: false,
    responseStream: false,
    requestType: jobs_pb.JobStatusRequest,
    responseType: jobs_pb.JobStatusResponse,
    requestSerialize: serialize_jobs_V1_JobStatusRequest,
    requestDeserialize: deserialize_jobs_V1_JobStatusRequest,
    responseSerialize: serialize_jobs_V1_JobStatusResponse,
    responseDeserialize: deserialize_jobs_V1_JobStatusResponse,
  },
  cancelJob: {
    path: '/jobs.V1.JobService/CancelJob',
    requestStream: false,
    responseStream: false,
    requestType: jobs_pb.CancelJobRequest,
    responseType: jobs_pb.CancelJobResponse,
    requestSerialize: serialize_jobs_V1_CancelJobRequest,
    requestDeserialize: deserialize_jobs_V1_CancelJobRequest,
    responseSerialize: serialize_jobs_V1_CancelJobResponse,
    responseDeserialize: deserialize_jobs_V1_CancelJobResponse,
  },
  echo: {
    path: '/jobs.V1.JobService/Echo',
    requestStream: false,
    responseStream: false,
    requestType: jobs_pb.EchoRequest,
    responseType: jobs_pb.EchoResponse,
    requestSerialize: serialize_jobs_V1_EchoRequest,
    requestDeserialize: deserialize_jobs_V1_EchoRequest,
    responseSerialize: serialize_jobs_V1_EchoResponse,
    responseDeserialize: deserialize_jobs_V1_EchoResponse,
  },
};

exports.JobServiceClient = grpc.makeGenericClientConstructor(JobServiceService);
