#!/bin/sh
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative laforge_proto/laforge_proto.proto