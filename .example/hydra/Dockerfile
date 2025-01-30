# Build the docker cli for docker consumption

# Build web_gen library
FROM node:20 AS client-builder

WORKDIR /client

COPY .spec/client .

RUN npm install

# Use an official Node runtime as the parent image
FROM node:20

# Set the working directory in the container
WORKDIR /app/cli/hydra/

# Copy the application source code
COPY .example/hydra .

COPY --from=client-builder client client

# Install the application dependencies
RUN npm install

# Set the entrypoint to run the CLI
ENTRYPOINT ["npm", "run", "cli"]