# leviathan

Container Orchestrator/Job Runner replacement for Autolab Tango

## Testing

To test out the dev build, ensure docker is installed on your system

You will need to run 2 services

1. The core leviathan service
    ```
     docker run --rm --network=host -v /var/run/docker.sock:/var/run/docker.sock ghcr.io/makeopensource/leviathan:dev 
    ```

2. Kraken, a simple node server with a frontend, that talks to leviathan, intended to mimic real usage
   ```
     docker run --rm --network=host ghcr.io/makeopensource/leviathan/kraken:dev 
   ```
   
3. Access the frontend at, http://localhost:3000

## Running a job

You will need to fill out 

1. Timeout: timeout for the job
2. Grader: a grader file
3. Makefile: a makefile (or any file really), it is intended as the entrypoint script to the grading job
4. Entry command: this is what starts the grading job, you may call your script above or call something directly
   * Eg python grade.py student.py
5. Student submission
6. Image tag: for your dockerfile
7. Dockerfile: Your grading dockerfile
   * You must ensure you have this line your dockerfile, as this is where all files will be placed 
   * ```WORKDIR /home/```
8. Leviathan will capture the stdout of the container, and attempt to parse the last line of the stdout
      * It expects a json string as the last line, the contents can be anything as long as it is valid json
      * If it is unable to parse the json, it will fail the job
