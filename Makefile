.PHONY: all
all: deps arena

# deps ensure
.PHONY: deps
deps:
	@echo "--> Running dependency management deps"
	#@go get -u github.com/golang/dep/cmd/dep
	@dep ensure -update -v

.PHONY: arena
arena:
	@echo "----------------------------"
	@echo "Making arena CLI tool"
	@go install github.com/skycoin/arena
	@echo "Successfully Created !!"
	@echo "----------------------------"
	@echo "Use command arena --help for info"
