GOFILES	= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -l $(GOFILES)

test: ## Runs test cases
	go test ./...

run: ## Runs the program
	go run main.go
