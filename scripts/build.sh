#!/bin/sh
mkdir -p $PWD/bin
GOBIN=$PWD/bin go install ./cmd/adder
