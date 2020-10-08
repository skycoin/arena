GIT ?= git
GO_VARS ?=
GO ?= go
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
ROOT := $(shell echo ${GOPATH}/src/)
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)
APP :=  addnumbers


.PHONY: help clean dependencies test
help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  build        to build the main binary for current platform"
	@echo "  test         to run unittests"

clean:
	rm -f guerrillad

vendor:
	dep ensure

build:
	$(GO_VARS) $(GO) build -o=$(APP) -ldflags="$(LD_FLAGS)" $(ROOT)/arena/cmd/addnumbers/


test:
	$(GO_VARS) $(GO) test -v $(ROOT)/arena/tests/add_test

docker:
	docker build -t $(APP) .
 