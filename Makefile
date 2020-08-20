build:
	echo "Building the binaries"
	go build -o bin/main main.go

run:
	go run main.go

test:
	echo "Running tests"
	go test -v ./...

lint:
	golangci-lint run ./...

format:
	goimports -l -w .

all:
	lint build run
