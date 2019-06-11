#!/bin/bash
while true; do sleep 1; dlv connect :2345 --init <(echo quit -c) && exit; done