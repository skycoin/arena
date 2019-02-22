SHELL := /bin/bash

.PHONY: \
	imports \

format:
	goimports -w -local github.com/skycoin/arena ./


