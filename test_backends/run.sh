#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
pids=()

cleanup() {
    echo "Stopping test backends..."
    for pid in "${pids[@]}"; do
        kill "$pid" 2>/dev/null || true
    done
    wait 2>/dev/null || true
}

trap cleanup EXIT INT TERM

start_backend() {
    local name="$1"
    echo "Starting $name..."
    go run "$SCRIPT_DIR/$name" &
    pids+=("$!")
}

start_backend alpha
start_backend beta
start_backend charlie

echo "Backends are running. Press Ctrl+C to stop them."
wait
