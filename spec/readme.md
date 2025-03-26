# Leviathan gRPC Specification

Leviathan uses [gRPC](https://grpc.io/) as its communication protocol, offering several advantages over traditional REST, such as 

* improved performance, 
* bidirectional streaming 
* efficient serialization
* type-safe apis and auto-generated clients

Leviathan uses a variant of gRPC called [ConnectRPC](https://connectrpc.com/docs/introduction/), which simplifies usage
compared to vanilla gRPC while maintaining compatibility and performance benefits.

This directory defines the gRPC specification and serves as the source for code generation.

## Requirements

To generate code for the api, you must install

* [Docker](https://docs.docker.com/engine/install/)
* [Just runner](https://just.systems/man/en/) - cross-platform alternative to makefile

## Code generation

To generate gRPC stubs, in the `spec` folder run,

```
just gen
```

This will:

* This will build the Dockerfile, which installs all the required dependencies for code gen 
* Runs the [gen-stubs.sh](./gen-stubs.sh) script, which calls the code gen CLI 
* The go files will be moved to the [go src](../src)
* The node files will be moved to [leviathan_node](./leviathan_node), which is npm package that can be used by any typescript project.

## Installing clients

### Node

To use the node TS code install via:

```
npm install 'https://gitpkg.vercel.app/makeopensource/leviathan/spec/leviathan_node?master'
```

This install the generated code on the ```master``` branch.

## Directory walkthrough

The folder contains the following folders and files

* [buf.gen.yaml](buf.gen.yaml) - the connect rpc config
* [buf.yaml](buf.yaml) - another connect rpc config
* [Dockerfile](Dockerfile) - Dockerfile that installs all dependencies for code gen
* [gen-stub.sh](gen-stubs.sh) - Script to run the connect rpc code gen cli 
* [Justfile](Justfile) - Justfile to run helpful commands
* [leviathan_node](leviathan_node) - node web client
* [proto](proto) - Contains the protocol definitions 
