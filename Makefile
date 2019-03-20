#!/usr/bin/make -f

.PHONY: all say_hello clean format test build

GOCMD=go
GO_FILES :=$(shell find . -name '*.go' -not -path './vendor/*' )
GOIMPORTS :=goimports -e
GOTEST= $(GOCMD) test
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

say_hello:
	@echo "Hello World"

format: ## Run goimports on all go files
	$(GOIMPORTS) -w -l $(GO_FILES)

test: ## Run all the tests
	$(GOTEST) -v ./pkg/sum/testdata

clean:
	@echo "cleaning up..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

build: ## Build the app
	dep ensure && go build

