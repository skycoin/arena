# Arena

## Tasks

1. Write a Golang program that takes two numbers from the command line, adds them up and prints the result.
2. Fork this repository and make a PR to this repository with the code.
3. Add information for contact in the PR. (Email, Telegram or Wechat.)

## Usage:

Method 1:

```go
go run main.go
```

Method 2: (Using docker image directly):
```docker
docker run -it --rm --name arena -v "$PWD":/go/src/github.com/skycoin/arena -w /go/src/github.com/skycoin/arena golang:1.14.2-stretch go run main.go
```

Method 3: (Using docker file):
```docker
docker build -t skycoin-arena .
```
```docker
docker run -it skycoin-arena arena 
```

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


