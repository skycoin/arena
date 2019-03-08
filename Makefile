# Go parameters

TOOLS = golang.org/x/tools/cmd/goimports 

tools: ; $(info $(M) building tools…)
	go get -v $(TOOLS)

format: 
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

test:
	go test ./...

add:
	go run main.go -operation 1