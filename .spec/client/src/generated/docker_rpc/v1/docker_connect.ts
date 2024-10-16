// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts"
// @generated from file docker_rpc/v1/docker.proto (package docker_rpc.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateContainerRequest, CreateContainerResponse, DeleteContainerRequest, DeleteContainerResponse, GetContainerLogRequest, GetContainerLogResponse, ListContainersRequest, ListContainersResponse, ListImageRequest, ListImageResponse, NewImageRequest, NewImageResponse, StartContainerRequest, StartContainerResponse, StopContainerRequest, StopContainerResponse } from "./docker_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service docker_rpc.v1.DockerService
 */
export const DockerService = {
  typeName: "docker_rpc.v1.DockerService",
  methods: {
    /**
     * @generated from rpc docker_rpc.v1.DockerService.CreateContainer
     */
    createContainer: {
      name: "CreateContainer",
      I: CreateContainerRequest,
      O: CreateContainerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.DeleteContainer
     */
    deleteContainer: {
      name: "DeleteContainer",
      I: DeleteContainerRequest,
      O: DeleteContainerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.ListContainers
     */
    listContainers: {
      name: "ListContainers",
      I: ListContainersRequest,
      O: ListContainersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.StartContainer
     */
    startContainer: {
      name: "StartContainer",
      I: StartContainerRequest,
      O: StartContainerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.StopContainer
     */
    stopContainer: {
      name: "StopContainer",
      I: StopContainerRequest,
      O: StopContainerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.GetContainerLogs
     */
    getContainerLogs: {
      name: "GetContainerLogs",
      I: GetContainerLogRequest,
      O: GetContainerLogResponse,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.CreateNewImage
     */
    createNewImage: {
      name: "CreateNewImage",
      I: NewImageRequest,
      O: NewImageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_rpc.v1.DockerService.ListImages
     */
    listImages: {
      name: "ListImages",
      I: ListImageRequest,
      O: ListImageResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

