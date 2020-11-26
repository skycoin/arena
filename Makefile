format:
	goimports -l -w .

lint:
	golangci-lint run

.PHONY: \
	format \