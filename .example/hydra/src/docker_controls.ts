import {createPromiseClient, Transport} from "@connectrpc/connect";
import {DockerService} from "leviathan-generated-sdk/src/generated/docker_rpc/v1/docker_connect";
import inquirer from "inquirer";
import {FileUpload} from "leviathan-generated-sdk/src/generated/docker_rpc/v1/docker_pb";
import {readFileAsBytes} from "./utils"

export function setupDocker(transport: Transport) {

    // @ts-ignore
    return {dkClient, dockerEndpoints}
}

