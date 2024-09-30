#!/usr/bin/env node

import {Command} from 'commander';
import {DockerApi} from 'leviathan-client';
import inquirer from 'inquirer';

const program = new Command();

program
    .version('1.0.0')
    .description('A CLI to interact with the Leviathan API');

const baseUrl = "http://localhost:9221"

// const coursesApi = new CoursesApi(undefined, "http://localhost:9221");
const dockerApi = new DockerApi(undefined, baseUrl);
const dockerEndpoints = {
    "Get Container info": async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dockerApi.dockerContainerIdGet(containerId as string);
        console.log(`Status: ${result.status}, message: ${result.statusText}`);
    },
    'Delete Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);

        const result = await dockerApi.dockerContainerIdDelete(parseInt(containerId));
        console.log(`Status: ${result.status}, message: ${result.statusText}`);
    },
    'Start Docker Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);
        const result = await dockerApi.dockerContainerIdStartGet(parseInt(containerId));
        console.log(`Status: ${result.status}, message: ${result.statusText}`);
    },
    'Stop Docker Container': async () => {
        const {containerId} = await inquirer.prompt([
            {type: 'input', name: 'containerId', message: 'Enter the container ID:'}
        ]);
        const result = await dockerApi.dockerContainerIdStopGet(parseInt(containerId));
        console.log(`Status: ${result.status}, message: ${result.statusText}`);
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