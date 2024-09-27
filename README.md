# leviathan
Container Orchestrator/Job Runner replacement for Autolab Tango

## TODO

* make leviathan ghcr public [docs](https://docs.github.com/en/packages/learn-github-packages/configuring-a-packages-access-control-and-visibility#selecting-the-inheritance-setting-for-packages-scoped-to-an-organization) 
* enable issues on makeopensoruce fork of leviathan [docs](https://stackoverflow.com/questions/16406180/is-there-a-way-to-add-issues-to-a-github-forked-repo-without-modifying-the-orig) 


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

## Things to test against

* fork bombs
* /dev/*** accesses
* zip bombs
* unauth accesses outside the working dir