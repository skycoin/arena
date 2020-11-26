format:
	@echo ">> Formatting all code using goimports"
	@goimports -w .

dep:
	@echo ">> Downloading Dependencies"
	@go mod download

run-sum:
	@echo ">> Running operation to add two numbers"
	@go run github.com/bobbiprathama/arena/cmd sum --first-number 10 --second-number 12

run-test-unit: dep
	@echo ">> Running Unit Test"
	@go test -tags=unit -cover -count=1 -failfast -covermode=atomic ./...