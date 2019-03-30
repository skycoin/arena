format:
	goimports -w ./cmd

lint:
	golangci-lint run --skip-dirs vendor