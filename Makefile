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

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

all:
	format lint build
