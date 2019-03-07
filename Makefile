default: clean format build test

format:
	go fmt

build: 
	go build

test:
	go test -v	

clean:
	go clean

	