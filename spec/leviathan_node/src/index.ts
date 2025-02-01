import {createClient} from "@connectrpc/connect";
import {createConnectTransport} from "@connectrpc/connect-node";
import {DockerService} from "./generated/docker_rpc/v1/docker_pb";
import {LabService} from "./generated/labs/v1/labs_pb";
import {JobService} from "./generated/jobs/v1/jobs_pb";
import {StatsService} from "./generated/stats/v1/stats_pb";

export * from "./generated/docker_rpc/v1/docker_pb"
export * from "./generated/jobs/v1/jobs_pb"
export * from "./generated/labs/v1/labs_pb"
export * from "./generated/stats/v1/stats_pb"
export * from "./generated/types/v1/types_pb"
export * from "@connectrpc/connect-node"
export * from "@connectrpc/connect"

export function createGrpcClient(baseUrl: string) {
    const transport = createConnectTransport({
        httpVersion: "2",
        baseUrl: baseUrl,
    });

    const dockerClient = createClient(DockerService, transport)
    const labClient = createClient(LabService, transport)
    const jobClient = createClient(JobService, transport)
    const statsClient = createClient(StatsService, transport)

    return {
        dockerClient,
        labClient,
        jobClient,
        statsClient,
    }
}

