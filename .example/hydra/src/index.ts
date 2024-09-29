#!/usr/bin/env node

import {Command} from 'commander';
import {CoursesApi, DockerApi} from 'leviathan-client'

const program = new Command();

program
    .version('1.0.0')
    .description('A simple TypeScript CLI application')
    .option('-n, --name <name>', 'Your name')
    .option('-g, --greeting <greeting>', 'Custom greeting', 'Hello')
    .action((options) => {
        const name = options.name || 'World';
        console.log(`${options.greeting}, ${name}!`);
    });

program.parse(process.argv);

const courses = new DockerApi(undefined ,"http://localhost:9221");

courses.dockerContainerIdDelete(94882).then(value => {
    console.log(value)
})