#!/usr/bin/env node

import {Command} from 'commander';
import {DockerService} from "leviathan-generated-sdk/src/generated/docker_rpc/v1/docker_connect";
import inquirer from 'inquirer';
import {createConnectTransport} from "@connectrpc/connect-node";
import {createPromiseClient} from "@connectrpc/connect";
import * as fs from "node:fs";
import {FileUpload} from "leviathan-generated-sdk/src/generated/docker_rpc/v1/docker_pb";

const program = new Command();

program
    .version('1.0.0')
    .description('A CLI to interact with the Leviathan API');

const baseUrl = "http://localhost:9221"

const transport = createConnectTransport({
    baseUrl: baseUrl,
    httpVersion: "2"
});

const dkClient = createPromiseClient(DockerService, transport)

const dockerEndpoints = {
    "Get Container info": async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dkClient.createContainer({})
    },
    'Delete Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dkClient.createContainer({})
    },
    'Start Docker Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dkClient.listContainers({})
    },
    'Stop Docker Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dkClient.stopContainer({containerId})
    },
    'Create Docker image': async () => {
        const {filepath} = await inquirer.prompt([
            {type: 'input', name: 'filepath', message: 'Enter dockerfile name:'}
        ]);

        const {imageTag} = await inquirer.prompt([
            {type: 'input', name: 'imageTag', message: 'Enter tagname name:'}
        ]);

        try {
            const contents = await readFileAsBytes(filepath);
            const payload = new FileUpload({filename: "newDockerfile", content: contents})
            const result = await dkClient.createNewImage({imageTag: imageTag,file: payload});
            console.log("Sent create docker image", result)
        } catch (error) {
            console.error(error)
            return
        }
    },
};

async function readFileAsBytes(filePath: string): Promise<Uint8Array> {
    try {
        const buffer = fs.readFileSync(filePath);
        return new Uint8Array(buffer);
    } catch (error) {
        console.error('Error reading file:', error);
        throw error;
    }
}


async function main() {
    while (true) {
        const {action} = await inquirer.prompt([
            {
                type: 'list',
                name: 'action',
                message: 'Choose an endpoint to call:',
                choices: [...Object.keys(dockerEndpoints), 'Exit']
            }
        ]);

        if (action === 'Exit') {
            console.log('Goodbye!');
            break;
        }

        try {
            const act = action as string
            // @ts-ignore
            await dockerEndpoints[act]();
        } catch (error) {
            // @ts-ignore
            console.error('An error occurred:', error.message);
        }

        console.log('\n');
    }
}

program.action(main);

program.parse(process.argv);