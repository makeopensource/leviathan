---
sidebar_position: 1
slug: /
title: About
---

# Leviathan

A container orchestrator and job runner for executing code in docker containers.

## Overview

Leviathan is the code execution engine for the [DevU project](https://github.com/makeopensource/devU).
It securely runs submitted code in isolated containers, providing evaluation results and feedback.
Designed as a modern replacement for [Autolab Tango](https://github.com/autolab/Tango).

Leviathan maintains full compatibility with existing Autolab/Tango graders. Any graders written for Tango will run on Leviathan without modification.

## Features
- Secure container-based execution
- Resource limiting and management
- gRPC API for type-safe, efficient communication

## How it Works

You will need to fill out

1. Timeout: timeout for the job
2. Job files - this contains any files require for your grading job:
    * this should contain your grader file, student submission
    * entry scripts: unlike autolab,
        * makefiles are optional, you can include a *.sh, python, ruby script or whatever tool you prefer
        * you can also directly call your grader using the command line by specifying it in the entry command
            * Eg python grade.py student.py
        * You must ensure the required tools/dependencies are installed in the dockerfile
3. Entry command: this is what starts the grading job
    * Eg python grade.py student.py
4. Image tag: for your dockerfile
5. Dockerfile: Your grading dockerfile
    * You must ensure you have this line your dockerfile, as this is where all files will be placed
    * ```WORKDIR /home/```
6. Leviathan will capture the stdout of the container, and attempt to parse the last line of the stdout
    * It expects a json string as the last line, the contents can be anything as long as it is valid json
    * If it is unable to parse the json, it will fail the job
