FROM node:23-alpine

# Set the working directory in the container
WORKDIR /app/

# Copy the application source code
COPY . .

# Install the application dependencies
RUN npm install

# Set the entrypoint to run the CLI
ENTRYPOINT ["npm", "run", "cli"]