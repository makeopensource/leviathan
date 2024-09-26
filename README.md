# leviathan
Container Orchestrator/Job Runner replacement for Autolab Tango

## Dev stuff

To generate go GRPC code

```bash
protoc --go_grpc_out=. --go_out=. .\proto\*.proto
```

To generate the node GRPC code

install the following dependencies

```
npm install -D grpc-tools grpc_tools_node_protoc_ts
```

```
protoc --plugin=protoc-gen-ts_proto=node_modules\.bin\protoc-gen-ts.cmd --ts_proto_out=src/generated --ts_proto_opt=import_style=commonjs,outputServices=grpc-js,env=node,useOptionals=messages,exportCommonSymbols=false,esModuleInterop=true -I ../proto ../proto/*.proto
```

