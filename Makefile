format:
	goimports .

lint:
	golangci-lint run

build:
	go build -o sum *.go

test:
	go test -race ./...

run: format lint build
	./sum 1.00001 2 3.1 6.002