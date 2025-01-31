import {createClient} from "@connectrpc/connect";
import {createConnectTransport} from "@connectrpc/connect-node";
import {CreateContainerRequest, DockerService} from "./generated/docker_rpc/v1/docker_pb";
import {LabService} from "./generated/labs/v1/labs_pb";
import {JobService} from "./generated/jobs/v1/jobs_pb";


export function initGrpcClients(baseUrl: string) {
    const transport = createConnectTransport({
        baseUrl: baseUrl,
        httpVersion: "2"
    });

    const dockerClient = createClient(DockerService, transport)
    const labClient = createClient(LabService, transport)
    const jobClient = createClient(JobService, transport)

    return {
        dockerClient,
        labClient,
        jobClient,
    };
}

