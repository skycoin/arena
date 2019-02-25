# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOLIST=$(GOCMD) list
    GOFORMAT=goimports
    BINARY_NAME=skycoin
    
    all: test build
    build: 
	$(GOBUILD) -o $(BINARY_NAME)
    test: 
	$(GOTEST) -v ./...
    clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
    format:
	$(GOFORMAT) -w main.go
	$(GOFORMAT) -w pkg/validator.go
	$(GOFORMAT) -w pkg/validator_test.go