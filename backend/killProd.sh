#! /bin/bash
PID=$(pgrep deployed-api)

if [ -z "$PID" ]; then
    kill $PID
fi