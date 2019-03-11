# Makefile

export GOPATH := $(shell pwd)

all:
	echo $$GOPATH
	go get -d
	go run *.go 1 2

build:
	echo $$GOPATH
	go get -d
	go build -o arena