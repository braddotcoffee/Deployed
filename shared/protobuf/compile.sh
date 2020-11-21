#! /bin/sh
protoc -I=. --go_out=../.. deployment.proto