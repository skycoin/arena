format:
	gofmt -w *.go && goimports -l -w .

test:
	cd ./testdata && go test -run ''
