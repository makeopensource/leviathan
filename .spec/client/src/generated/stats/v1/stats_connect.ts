// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts"
// @generated from file stats/v1/stats.proto (package stats.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { EchoRequest, EchoResponse } from "./stats_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * todo
 *
 * @generated from service stats.v1.StatsService
 */
export const StatsService = {
  typeName: "stats.v1.StatsService",
  methods: {
    /**
     * @generated from rpc stats.v1.StatsService.Echo
     */
    echo: {
      name: "Echo",
      I: EchoRequest,
      O: EchoResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

