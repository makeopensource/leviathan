// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var stats_pb = require('./stats_pb.js');

function serialize_stats_V1_EchoRequest(arg) {
  if (!(arg instanceof stats_pb.EchoRequest)) {
    throw new Error('Expected argument of type stats.V1.EchoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stats_V1_EchoRequest(buffer_arg) {
  return stats_pb.EchoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stats_V1_EchoResponse(arg) {
  if (!(arg instanceof stats_pb.EchoResponse)) {
    throw new Error('Expected argument of type stats.V1.EchoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_stats_V1_EchoResponse(buffer_arg) {
  return stats_pb.EchoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// import "types.proto";
//
// todo
var StatsServiceService = exports.StatsServiceService = {
  echo: {
    path: '/stats.V1.StatsService/Echo',
    requestStream: false,
    responseStream: false,
    requestType: stats_pb.EchoRequest,
    responseType: stats_pb.EchoResponse,
    requestSerialize: serialize_stats_V1_EchoRequest,
    requestDeserialize: deserialize_stats_V1_EchoRequest,
    responseSerialize: serialize_stats_V1_EchoResponse,
    responseDeserialize: deserialize_stats_V1_EchoResponse,
  },
};

exports.StatsServiceClient = grpc.makeGenericClientConstructor(StatsServiceService);
