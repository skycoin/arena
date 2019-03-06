format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -l github.com/skycoin/arena ./arena ./

test: ## Runs test cases
	go test ./...

run: ## Runs the program
	go run main.go
