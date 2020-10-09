build:
	go build -o bin/commandline commandline.go

run:
	go run commandline.go $a $b
