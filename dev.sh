#!/usr/bin/env bash
set -e
set -o xtrace

go build -o build/chat-thing-dev
build/chat-thing-dev $@
