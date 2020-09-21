# Arena

## Tasks

1. Write a Golang program that takes two numbers from the command line, adds them up and prints the result.
2. Fork this repository and make a PR to this repository with the code.
3. Add information for contact in the PR. (Email, Telegram or Wechat.)

##Usage

Method 1(using golang):

1.Build the golang Application and generate binary

```go build main.go```

2.To run the Application run below command:

```./main add --firstNumber=<enter any integer number> --secondNumber=<enter any integer number> ```


Method 2(Using Docker and Makefile):

1.Build the docker image using below command:

```make build```

2.Run the application using below command:

```make run firstnumber=<enter any number> secondnumber=<enter any number>```
## Guidelines
Here are some guidelines that will give you a clearer picture of what is expected in terms of coding style and conventions.

1: Branch and PR guide:
https://github.com/skycoin/skycoin/wiki/Github-branch-and-PR-guide

2: Idiomatic Go guide:
https://github.com/skycoin/skycoin/wiki/Idiomatic-Go-%28in-Skycoin%29

3: Dependency management:
https://github.com/skycoin/skycoin/wiki/Dependency-management-in-Skycoin

4: Linter:
https://github.com/golangci/golangci-lint


