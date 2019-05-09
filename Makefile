test:
	go test ./...

format:
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

build:
	go build cmd/add/add.go
