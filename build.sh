#!/bin/bash

protoc -I. --go_out=. --go_opt=paths=source_relative sodor/sodor.proto
protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative sodor/fat_ctrl.proto
protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative sodor/thomas.proto