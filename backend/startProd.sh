#! /bin/bash

go build -o deployed-api .
PORT=32980 ./deployed-api &
disown -h %1