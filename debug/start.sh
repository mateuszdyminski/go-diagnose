#!/bin/bash

while true; do sleep 1; dlv connect :2345 --init <(echo quit -c) && exit; done &

dlv exec --headless --listen=:2345 --accept-multiclient --api-version=2 --log /Users/md/workspace/go/src/github.com/mateuszdyminski/go-diagnose/debug/app