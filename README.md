# Arena
## Program
Description: This code will take exact two command line input and print the sum of these two.

# How to run

1. Without makefile
	a. go build -o bin/commandline commandline.go
	b. bin/commandline 10 15 
	(Note: 10 and 15 are two variables you can pass of your choice)
	
		or
	a. go run commandline.go 10 15

2. With make file
	a. make -f Makefile
	b. make run a=10 b=15
