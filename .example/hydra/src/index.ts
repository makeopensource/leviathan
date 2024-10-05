#!/usr/bin/env node

import {Command} from 'commander';
import {DockerService} from "leviathan-generated-sdk/src/generated/docker_rpc/v1/docker_connect";
import inquirer from 'inquirer';
import {createConnectTransport} from "@connectrpc/connect-node";
import {createPromiseClient} from "@connectrpc/connect";

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

        const result = await dkClient.echo({})
    },
    // todo
    // 'Create Docker image': async () => {
    //     const file = fs.readFileSync('../ex-Dockerfile', 'utf8');
    //     // const result = await dockerApi.dockerImagesCreatePost(file ,"test:latest");
    //     // console.log(JSON.stringify(result, null, 2));
    // },
};

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