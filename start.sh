#!/bin/bash

function cleanup {
    echo "Killed all child processes"
    if [ -e "$GO_PROGRAM_PATH/main" ]; then
        rm -f "$GO_PROGRAM_PATH/main"
    fi
    pkill -P $$
}

trap cleanup EXIT

NATS_STREAMING_SERVER_PATH=nats-streaming-server

GO_PROGRAM_PATH="$(dirname "$0")"

echo "Launching nats-streaming-server..."
exec "$NATS_STREAMING_SERVER_PATH" &

sleep 1

echo "Launching service..."
go build -o main "$GO_PROGRAM_PATH/cmd"
if [ -e "$GO_PROGRAM_PATH/main" ]; then
    "$GO_PROGRAM_PATH/main"
fi
