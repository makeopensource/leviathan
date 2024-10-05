# OpenApi Spec

This is where the spec is defined and the code is generated.

# What is this

We implement the api using the [spec](./leviathan.yaml), this contains the all paths and types defined in
the [open API spec](https://swagger.io/specification/v3/) format.
This allows us to autogenerate the client and server code, in a (relatively)typesafe manner, the spec only defines the
types it is up to us to follow it.

## Generated code usage

The server files are automatically moved to [internal/generated-server](../internal/generated-server), can be directly
used.

To use the client TS code install it via:

```
npm install 'https://gitpkg.vercel.app/makeopensource/leviathan/.spec/client?master'
```

This install the generated code on the ```master``` branch.

## Directory walkthrough

The folder contains the following folders and files

* [leviathan.yaml](./leviathan.yaml) - This is the actual spec file where the definitions are written.
* [config-go-server.yml](./config-go-server.yml) - Configuration options for generating the go server
  code, [Possible options](https://openapi-generator.tech/docs/generators/go-gin-server)
* [config-ts-client.yml](./config-ts-client.yml) - Configuration options for generating the typescript client
  code, [Possible options](https://openapi-generator.tech/docs/generators/typescript-axios)
* [Makefile](./Makefile) - Makefile to run code generation and other helpful commands.

##### Generation directories

YOU SHOULD NEVER MODIFY THE FILES OR CODE IN THIS DIRECTORY.

* [server](./server) - This is where the generated server-side code is outputted.
* [client](./client) - This is where the generated client-side code is outputted.

## Development setup

1. 

##### Node setup

```
npm i @connectrpc/protoc-gen-connect-es -g
```

```
npm i @bufbuild/protoc-gen-es -g
```

nce you have set this up, you should be good to go.

## Code generation

We have included a handy makefile to generate client and server code.

To generate the server side go code,
This will generate the code, then copy the resulting api stubs
to [internal/generated-server](../internal/generated-server)

Before you run the commands make sure you are in the [spec](.) directory to access the makefile.

```
make gensrv
```

To generate the client side TS code

```
make genclient
```
