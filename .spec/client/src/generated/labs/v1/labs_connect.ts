// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts"
// @generated from file labs/v1/labs.proto (package labs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteLabRequest, DeleteLabResponse, EditLabResponse, LabRequest, NewLabResponse } from "./labs_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service labs.v1.LabService
 */
export const LabService = {
  typeName: "labs.v1.LabService",
  methods: {
    /**
     * @generated from rpc labs.v1.LabService.NewLab
     */
    newLab: {
      name: "NewLab",
      I: LabRequest,
      O: NewLabResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc labs.v1.LabService.EditLab
     */
    editLab: {
      name: "EditLab",
      I: LabRequest,
      O: EditLabResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc labs.v1.LabService.DeleteLab
     */
    deleteLab: {
      name: "DeleteLab",
      I: DeleteLabRequest,
      O: DeleteLabResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

