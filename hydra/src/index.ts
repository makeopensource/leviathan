#!/usr/bin/env node

import {Command} from 'commander';
import inquirer from 'inquirer';
import {createPromiseClient} from "@connectrpc/connect";
import {DockerService} from "leviathan-generated-sdk/src/generated/docker_rpc/v1/docker_pb";
import {readFileAsBytes} from "./utils";
import {LabService} from "leviathan-generated-sdk/src/generated/labs/v1/labs_pb";
import path from "path";
import { createConnectTransport } from "@connectrpc/connect-node";
import { createClient } from "@connectrpc/connect";

const program = new Command();

program
    .version('1.0.0')
    .description('A CLI to interact with the Leviathan API');

const dockerEndpoints = {
    "List images": async () => {
        const result = await dkClient.listImages({})
        for (const image of result.images) {
            console.log("Machine", image.id)
            for (const metadata of image.metadata) {
                console.log(metadata)
            }
        }
    },
    "List containers": async () => {
        const result = await dkClient.listContainers({})
        for (const image of result.containers) {
            console.log("Machine", image.id)
            for (const metadata of image.metadata) {
                console.log(metadata)
            }
        }
    },
    "Get Container info": async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        // const result = await dkClient.createContainer({})
    },
    'Delete Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        // const result = await dkClient.createContainer({})
    },
    'Start Docker Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dkClient.startContainer({combinedId: containerId})
    },
    'Stop Docker Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dkClient.stopContainer({combinedId: containerId})
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
            const result = await dkClient.createNewImage({imageTag: imageTag, file: payload});
            console.log("Sent create docker image", result)
        } catch (error) {
            console.error(error)
            return
        }
    },
    'Create Docker container': async () => {
        const getImageList = await dkClient.listImages({})

        let fullImageList = {};

        for (const image of getImageList.images) {
            for (const metadata of image.metadata) {
                // @ts-ignore
                fullImageList[metadata.RepoTags[0]] = metadata.RepoTags[0]
            }
        }

        // list docker images
        const {imageTag} = await inquirer.prompt([
            {
                type: 'list',
                name: 'imageTag',
                message: 'Choose an image to use:',
                choices: [...Object.keys(fullImageList)]
            }
        ]);

        // select image
        const res = await dkClient.createContainer({imageTag})
    },

    'Get Container Logs': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        try {
            const result = dkClient.getContainerLogs({combinedId: containerId});
            for await (const logs of result) {
                console.log("Logs");
                console.log(logs.logs)
            }
        } catch (error) {
            console.error(error)
            return
        }
    },
};

const labClient = createPromiseClient(LabService, transport)
const labEndpoints = {
    "Create lab": async () => {
        const {labName} = await inquirer.prompt([
            {type: 'input', name: 'labName', message: 'Enter lab name:'}
        ]);
        const {filepath} = await inquirer.prompt([
            {type: 'input', name: 'filepath', message: 'Enter grader file path:'}
        ]);

        const filename = path.basename(filepath)
        const contents = await readFileAsBytes(filepath);

        const payload = new FileUpload({filename: filename, content: contents})
        const payload2 = new FileUpload({filename: filename + "2", content: contents})
        
        const result = await labClient.newLab({LabName: labName, graderFile: payload, makeFile: payload2})
    },
    "Edit lab": async () => {
        const {labName} = await inquirer.prompt([
            {type: 'input', name: 'labName', message: 'Enter lab name:'}
        ]);
        const {filepath} = await inquirer.prompt([
            {type: 'input', name: 'filepath', message: 'Enter grader file path:'}
        ]);

        const filename = path.basename(filepath)
        const contents = await readFileAsBytes(filepath);

        const payload = new FileUpload({filename: filename, content: contents})
        const result = await labClient.editLab({LabName: labName, graderFile: payload})
    },
    "Delete lab": async () => {
        const {labName} = await inquirer.prompt([
            {type: 'input', name: 'labName', message: 'Enter grader file name:'}
        ]);
        const result = await labClient.deleteLab({LabName: labName})
    },
}


async function main() {
    const allEndpoints = {"Docker endpoints": dockerEndpoints, "Lab endpoint": labEndpoints}

    while (true) {
        const {endpoint} = await inquirer.prompt([
            {
                type: 'list',
                name: 'endpoint',
                message: 'Choose an endpoint to call:',
                choices: [...Object.keys(allEndpoints), 'Exit']
            }
        ]);

        if (endpoint === 'Exit') {
            console.log('Goodbye!');
            break;
        }


        const {action} = await inquirer.prompt([
            {
                type: 'list',
                name: 'action',
                message: 'Choose an endpoint to call:',
                // @ts-ignore
                choices: [...Object.keys(allEndpoints[endpoint]), 'Exit']
            }
        ]);


        try {
            const act = action as string
            // @ts-ignore
            await allEndpoints[endpoint][action]();
        } catch (error) {
            // @ts-ignore
            console.error('An error occurred:', error.message);
        }
        console.log('\n');

    }
}

program.action(main);

program.parse(process.argv);