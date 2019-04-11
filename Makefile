test:
	go test

build:
	go build -o ./arena cmd/add.go

lint:
	golangci-lint run

format:
	goimports -w -local github.com/skycoin/hardware-wallet-daemon ./cmd arena.go arena_test.go

.DEFAULT_GOAL := build