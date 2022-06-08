# gRPC Server

Contains public sector information licensed under the Open Government Licence v3.0.

## Overview

This is a small gRPC server for learning.

## Description  

This is a small project for learning pruporses. It creates a gRPC server that serves data about London Fire Brigade fleet and also about animal rescue incidents attended by LFB.
The datasets used are public available under [Open Government Licence v3.0](https://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/)

## Make commands
- `make init`  
Inits the repository removing undesirable files and update dependencies. 

- `make build`  
Builds the application. Executables are put into `build` directory

- `make docker-build`  
Builds the docker image

- `make proto`  
Generate the protobuf stubs from proto definitions

- `make test`  
Run all project tests
