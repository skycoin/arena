GOFILES	= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -l -d -e $(GOFILES)

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	vendorcheck ./...
	golangci-lint run --skip-dirs vendor

test:
	go test -v ./...

run:
	go run ./cmd/add_cli/add_cli.go

update_packages: 
	dep ensure -update -v
