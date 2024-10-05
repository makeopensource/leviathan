"use strict";
// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts"
// @generated from file labs/v1/labs.proto (package labs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck
Object.defineProperty(exports, "__esModule", { value: true });
exports.LabService = void 0;
const labs_pb_js_1 = require("./labs_pb.js");
const protobuf_1 = require("@bufbuild/protobuf");
/**
 * @generated from service labs.v1.LabService
 */
exports.LabService = {
    typeName: "labs.v1.LabService",
    methods: {
        /**
         * @generated from rpc labs.v1.LabService.NewLab
         */
        newLab: {
            name: "NewLab",
            I: labs_pb_js_1.NewLabRequest,
            O: labs_pb_js_1.NewLabResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
        /**
         * @generated from rpc labs.v1.LabService.EditLab
         */
        editLab: {
            name: "EditLab",
            I: labs_pb_js_1.EditLabRequest,
            O: labs_pb_js_1.EditLabResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
        /**
         * @generated from rpc labs.v1.LabService.DeleteLab
         */
        deleteLab: {
            name: "DeleteLab",
            I: labs_pb_js_1.DeleteLabRequest,
            O: labs_pb_js_1.DeleteLabResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
        /**
         * @generated from rpc labs.v1.LabService.Echo
         */
        echo: {
            name: "Echo",
            I: labs_pb_js_1.EchoRequest,
            O: labs_pb_js_1.EchoResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
    }
};
