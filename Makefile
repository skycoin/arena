format:
	goimports -l -w .

lint:
	golangci-lint run

install:
	go install

.PHONY: \
	format \
	lint \
	install \