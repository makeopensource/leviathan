// implementation for interacting with the leviathan GRPC API

import grpc from '@grpc/grpc-js';
import {DockerServiceClient} from "./gen/docker_grpc_pb";
import {JobServiceClient} from "./gen/jobs_grpc_pb";
import {CreateRequest} from "./gen/docker_pb";
import {NewJobRequest} from "./gen/jobs_pb";


export function client() {
    const dockerClient = new DockerServiceClient('localhost:50051', grpc.credentials.createInsecure());
    const jobsClient = new JobServiceClient('localhost:50051', grpc.credentials.createInsecure())

    dockerClient.createContainer(new CreateRequest(), (err, resp) => {
        if (err) {
            console.error(err);
        } else {
            console.log("Response :", resp);
        }
    });

    jobsClient.newJob(new NewJobRequest(), (err, resp) => {
        if (err) {
            console.error(err);
        } else {
            console.log("Response :", resp);
        }
    });
}
