#! /bin/bash

./killProd.sh
go build -o deployed-api .
PORT=32980 PRODUCTION=true ./deployed-api &
disown -h %1