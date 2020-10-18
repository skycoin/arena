all: format test build run

format:
	goimports -w ./cmd

test:
	go test ./...

build:
	go build -o ./bin/sum_of_args ./cmd/...

run:
	./bin/sum_of_args 1 2
