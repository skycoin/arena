test:
	go test -v ./...
build:
	go build main.go
lint:
	golangci-lint run -v ./...

