// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var labs_pb = require('./labs_pb.js');

function serialize_labs_V1_DeleteLabRequest(arg) {
  if (!(arg instanceof labs_pb.DeleteLabRequest)) {
    throw new Error('Expected argument of type labs.V1.DeleteLabRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_DeleteLabRequest(buffer_arg) {
  return labs_pb.DeleteLabRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_DeleteLabResponse(arg) {
  if (!(arg instanceof labs_pb.DeleteLabResponse)) {
    throw new Error('Expected argument of type labs.V1.DeleteLabResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_DeleteLabResponse(buffer_arg) {
  return labs_pb.DeleteLabResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_EchoRequest(arg) {
  if (!(arg instanceof labs_pb.EchoRequest)) {
    throw new Error('Expected argument of type labs.V1.EchoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_EchoRequest(buffer_arg) {
  return labs_pb.EchoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_EchoResponse(arg) {
  if (!(arg instanceof labs_pb.EchoResponse)) {
    throw new Error('Expected argument of type labs.V1.EchoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_EchoResponse(buffer_arg) {
  return labs_pb.EchoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_EditLabRequest(arg) {
  if (!(arg instanceof labs_pb.EditLabRequest)) {
    throw new Error('Expected argument of type labs.V1.EditLabRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_EditLabRequest(buffer_arg) {
  return labs_pb.EditLabRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_EditLabResponse(arg) {
  if (!(arg instanceof labs_pb.EditLabResponse)) {
    throw new Error('Expected argument of type labs.V1.EditLabResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_EditLabResponse(buffer_arg) {
  return labs_pb.EditLabResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_NewLabRequest(arg) {
  if (!(arg instanceof labs_pb.NewLabRequest)) {
    throw new Error('Expected argument of type labs.V1.NewLabRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_NewLabRequest(buffer_arg) {
  return labs_pb.NewLabRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_labs_V1_NewLabResponse(arg) {
  if (!(arg instanceof labs_pb.NewLabResponse)) {
    throw new Error('Expected argument of type labs.V1.NewLabResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_labs_V1_NewLabResponse(buffer_arg) {
  return labs_pb.NewLabResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var LabServiceService = exports.LabServiceService = {
  newLab: {
    path: '/labs.V1.LabService/NewLab',
    requestStream: false,
    responseStream: false,
    requestType: labs_pb.NewLabRequest,
    responseType: labs_pb.NewLabResponse,
    requestSerialize: serialize_labs_V1_NewLabRequest,
    requestDeserialize: deserialize_labs_V1_NewLabRequest,
    responseSerialize: serialize_labs_V1_NewLabResponse,
    responseDeserialize: deserialize_labs_V1_NewLabResponse,
  },
  editLab: {
    path: '/labs.V1.LabService/EditLab',
    requestStream: false,
    responseStream: false,
    requestType: labs_pb.EditLabRequest,
    responseType: labs_pb.EditLabResponse,
    requestSerialize: serialize_labs_V1_EditLabRequest,
    requestDeserialize: deserialize_labs_V1_EditLabRequest,
    responseSerialize: serialize_labs_V1_EditLabResponse,
    responseDeserialize: deserialize_labs_V1_EditLabResponse,
  },
  deleteLab: {
    path: '/labs.V1.LabService/DeleteLab',
    requestStream: false,
    responseStream: false,
    requestType: labs_pb.DeleteLabRequest,
    responseType: labs_pb.DeleteLabResponse,
    requestSerialize: serialize_labs_V1_DeleteLabRequest,
    requestDeserialize: deserialize_labs_V1_DeleteLabRequest,
    responseSerialize: serialize_labs_V1_DeleteLabResponse,
    responseDeserialize: deserialize_labs_V1_DeleteLabResponse,
  },
  echo: {
    path: '/labs.V1.LabService/Echo',
    requestStream: false,
    responseStream: false,
    requestType: labs_pb.EchoRequest,
    responseType: labs_pb.EchoResponse,
    requestSerialize: serialize_labs_V1_EchoRequest,
    requestDeserialize: deserialize_labs_V1_EchoRequest,
    responseSerialize: serialize_labs_V1_EchoResponse,
    responseDeserialize: deserialize_labs_V1_EchoResponse,
  },
};

exports.LabServiceClient = grpc.makeGenericClientConstructor(LabServiceService);
