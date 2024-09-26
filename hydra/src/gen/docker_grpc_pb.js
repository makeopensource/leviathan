// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var docker_pb = require('./docker_pb.js');

function serialize_docker_rpc_V1_CreateRequest(arg) {
  if (!(arg instanceof docker_pb.CreateRequest)) {
    throw new Error('Expected argument of type docker_rpc.V1.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_CreateRequest(buffer_arg) {
  return docker_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_CreateResponse(arg) {
  if (!(arg instanceof docker_pb.CreateResponse)) {
    throw new Error('Expected argument of type docker_rpc.V1.CreateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_CreateResponse(buffer_arg) {
  return docker_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_DeleteRequest(arg) {
  if (!(arg instanceof docker_pb.DeleteRequest)) {
    throw new Error('Expected argument of type docker_rpc.V1.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_DeleteRequest(buffer_arg) {
  return docker_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_DeleteResponse(arg) {
  if (!(arg instanceof docker_pb.DeleteResponse)) {
    throw new Error('Expected argument of type docker_rpc.V1.DeleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_DeleteResponse(buffer_arg) {
  return docker_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_EchoRequest(arg) {
  if (!(arg instanceof docker_pb.EchoRequest)) {
    throw new Error('Expected argument of type docker_rpc.V1.EchoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_EchoRequest(buffer_arg) {
  return docker_pb.EchoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_EchoResponse(arg) {
  if (!(arg instanceof docker_pb.EchoResponse)) {
    throw new Error('Expected argument of type docker_rpc.V1.EchoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_EchoResponse(buffer_arg) {
  return docker_pb.EchoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_ListRequest(arg) {
  if (!(arg instanceof docker_pb.ListRequest)) {
    throw new Error('Expected argument of type docker_rpc.V1.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_ListRequest(buffer_arg) {
  return docker_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_docker_rpc_V1_ListResponse(arg) {
  if (!(arg instanceof docker_pb.ListResponse)) {
    throw new Error('Expected argument of type docker_rpc.V1.ListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_docker_rpc_V1_ListResponse(buffer_arg) {
  return docker_pb.ListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var DockerServiceService = exports.DockerServiceService = {
  createContainer: {
    path: '/docker_rpc.V1.DockerService/CreateContainer',
    requestStream: false,
    responseStream: false,
    requestType: docker_pb.CreateRequest,
    responseType: docker_pb.CreateResponse,
    requestSerialize: serialize_docker_rpc_V1_CreateRequest,
    requestDeserialize: deserialize_docker_rpc_V1_CreateRequest,
    responseSerialize: serialize_docker_rpc_V1_CreateResponse,
    responseDeserialize: deserialize_docker_rpc_V1_CreateResponse,
  },
  deleteContainer: {
    path: '/docker_rpc.V1.DockerService/DeleteContainer',
    requestStream: false,
    responseStream: false,
    requestType: docker_pb.DeleteRequest,
    responseType: docker_pb.DeleteResponse,
    requestSerialize: serialize_docker_rpc_V1_DeleteRequest,
    requestDeserialize: deserialize_docker_rpc_V1_DeleteRequest,
    responseSerialize: serialize_docker_rpc_V1_DeleteResponse,
    responseDeserialize: deserialize_docker_rpc_V1_DeleteResponse,
  },
  listContainers: {
    path: '/docker_rpc.V1.DockerService/ListContainers',
    requestStream: false,
    responseStream: false,
    requestType: docker_pb.ListRequest,
    responseType: docker_pb.ListResponse,
    requestSerialize: serialize_docker_rpc_V1_ListRequest,
    requestDeserialize: deserialize_docker_rpc_V1_ListRequest,
    responseSerialize: serialize_docker_rpc_V1_ListResponse,
    responseDeserialize: deserialize_docker_rpc_V1_ListResponse,
  },
  echo: {
    path: '/docker_rpc.V1.DockerService/Echo',
    requestStream: false,
    responseStream: false,
    requestType: docker_pb.EchoRequest,
    responseType: docker_pb.EchoResponse,
    requestSerialize: serialize_docker_rpc_V1_EchoRequest,
    requestDeserialize: deserialize_docker_rpc_V1_EchoRequest,
    responseSerialize: serialize_docker_rpc_V1_EchoResponse,
    responseDeserialize: deserialize_docker_rpc_V1_EchoResponse,
  },
};

exports.DockerServiceClient = grpc.makeGenericClientConstructor(DockerServiceService);
