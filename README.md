# Skycoin: Addition
- [Skycoin: Addition](#skycoin-addition)
  - [Building the source](#building-the-source)
    - [Prerequisites](#prerequisites)
    - [Build](#build)
  - [Usage](#usage)
  - [Testing](#testing)
  - [Future enhancement](#future-enhancement)

## Building the source

### Prerequisites
goLang
dep: `go get -u github.com/golang/dep/cmd/dep`

### Build
Clone the repository
1. cd `$GOPATH/src/github.com`
2. `git clone https://github.com/{username}/arena`

The build process is set to be automatic you just need to install the package using:

`dep ensure`

`make tool`

## Usage

`make add`

Enter number and press enter, enter second number.

## Testing

`make test`

Will run all the tests for all packages (src/*).

## Future enhancement

We can add more operations like subtract, divide, just by adding a specific function in `src/{operation}/{operation}.go`.
And adding them in `cmd/cli.go`