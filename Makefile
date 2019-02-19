# Go parameters

tools:
	go get golang.org/x/tools/cmd/goimports

format: 
	goimports -l -w $$(find . -type f -name '.go' -not -path "./vendor/*")