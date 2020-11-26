format:
	goimports -l -w .

lint:
	golangci-lint run

install:
	go install

test:
	go install && go test ./cmd/



.PHONY: \
	format \
	lint \
	install \
	test \