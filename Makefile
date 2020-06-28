GOPATH=$(shell pwd)/vendor

go-format:
	@GOPATH=$(GOPATH) goimports -w ./
