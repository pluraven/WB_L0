#!/bin/bash

function cleanup {
    echo "Killed all child processes"
    if [ -e "$GO_PROGRAM_PATH/main" ]; then
        rm -f "$GO_PROGRAM_PATH/main"
    fi
    pkill -P $$
}

trap cleanup EXIT

GO_PROGRAM_PATH="$HOME/Documents/WBTech/L0"

echo "Launching nats-streaming-server..."
exec nats-streaming-server &

sleep 1

echo "Launching service..."
if [ ! -e "$GO_PROGRAM_PATH/main" ]; then
    go build -o "$GO_PROGRAM_PATH/main" "$GO_PROGRAM_PATH/cmd"
fi
"$GO_PROGRAM_PATH/main"
