ARENA_SKYCOIN?=0.0.0
COMMIT=`git rev-parse --short HEAD`

build: build_binary

build_mac:
	@GOOS=darwin GOARCH=amd64 \
	go build -v --ldflags "-w \
	-X github.com/mrcaelumn/arena/version.Version=$(ARENA_SKYCOIN) \
	-X github.com/mrcaelumn/arena/version.GitCommit=$(COMMIT)" .

install-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go install -v --ldflags "-w \
	-X github.com/mrcaelumn/arena/version.Version=$(ARENA_SKYCOIN) \
	-X github.com/mrcaelumn/arena/version.GitCommit=$(COMMIT)" .

build_binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o arena -a --ldflags "-w \
	-X github.com/mrcaelumn/arena/version.Version=$(ARENA_SKYCOIN) \
	-X github.com/mrcaelumn/arena/version.GitCommit=$(COMMIT)" .

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet -v $(shell go list ./... | grep -v /vendor/)

clean:
	@rm -rf build
	@rm -rf arena*

.PHONY: test vet build build_binary clean
