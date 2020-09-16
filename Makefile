format:
	go get golang.org/x/tools/cmd/goimports
	find . -name \*.go -not -path ./vendor -exec goimports -w {} \;

build_add:
	cd cmd/add && go build -o ../bin/

build_all: build_add

test:
	go test ./...