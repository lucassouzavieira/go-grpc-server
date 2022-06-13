# gRPC Server

Contains public sector information licensed under the Open Government Licence v3.0. Check [data/LICENSE.md](data/LICENSE.md) to have check out details of the license.

## Overview

This is a small gRPC server for learning.

## Description  

This is a small project for learning pruporses. It creates a gRPC server that serves data about London Fire Brigade fleet and also about animal rescue incidents attended by LFB.
The datasets used are public available under [Open Government Licence v3.0](https://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/)

## CLI

A CLI client is included in this project under `cmd/cli` directory. 
In order to use the CLI, you should build the gRPC server and have it running. You can specify the address for it by passing `-s` parameter 
```
Usage: grpc-client [-s] COMMAND [arg...]

A gRPC CLI client for go-grpc-server project
                 
Options:         
  -s, --server   gRPC server address (default "localhost:8080")
                 
Commands:        
  incidents      Handle LFB incidents info
  fleet          Handle LFB fleet info
                 
Run 'grpc-client COMMAND --help' for more information on a command.
```

## Building
### Make commands
- `make build`  
Builds the application. Executables are put into `build` directory

- `make build-cli`  
Builds the CLI gRPC client for this application. Executables are put into `build` directory

- `make docker-build`  
Builds the docker image

- `make proto`  
Generate the protobuf stubs from proto definitions

- `make test`  
Run all project tests
