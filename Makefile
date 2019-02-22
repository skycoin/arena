build:
	@echo "--> App building "
	go build -o arena .

run:
	./arena

format:
	goimports