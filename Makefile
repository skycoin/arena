deps:
	go mod download

test:
	@go test ./...

build:
	@go build cmd/cli/main.go
