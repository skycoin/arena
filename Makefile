install-linters: ## Install linters
	go get -u golang.org/x/tools/cmd/goimports

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -local github.com/skycoin/arena ./cmd
	goimports -w -local github.com/skycoin/arena ./pkg