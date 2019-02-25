format:
	goimports -w ./cmd ./*.go

lint:
	golangci-lint run --skip-dirs vendor

test:
	go test

run:
	go run ./cmd/add_cli.go