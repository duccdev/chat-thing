#!/usr/bin/env bash
set -e

if [ "$1" = "" ] || [ "$2" = "" ]; then
  echo "Usage: $0 <os> <arch>" >&2
  exit 1
fi

set -o xtrace
GOOS="$1" GOARCH="$2" go build -ldflags="-w -s" -o "build/chat-thing-$GOOS-$GOARCH"
