
BIN=arena_sum.exe
GOSRCS=sum.go

.PHONY: all

all: fmt
	go build -o $(BIN) $(GOSRCS)

fmt:
	go get golang.org/x/tools/cmd/goimports
	$(GOPATH)/bin/goimports -w $(GOSRCS)

clean:
	rm -rf $(BIN)
