PKG 	:= github.com/skycoin/arena
GOFILES	= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: format
format: goimports
	@goimports -local $(PKG) -w -l $(GOFILES)

.PHONY: goimports
goimports:
	@command -v goimports >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "installing goimports"; \
		go get golang.org/x/tools/cmd/goimports; \
	fi
