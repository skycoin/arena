# arena [hari version]

## Requirement

- Go version 1.11.+ (we are using go modules feature: https://github.com/golang/go/wiki/Modules)

## Structure

```
.
├── README.md
├── bin
│   └── arena
├── cmd
│   └── add
│       └── add.go
├── go.mod
├── go.sum
├── main.go
├── run_demos.sh
└── utils
    ├── math.go
    └── validation.go
```

## Installation

- download repositories

```shell
git clone https://github.com/skycoin/arena
```

- checkout to related branch (example: branch hari_test)

```shell
git checkout hari_test
```

- execute arena demos

```shell
sh run_demos.sh
```

#### Sample Result:

```shell
Build Arena...
-------------------------------------
Arena demos.
-------------------------------------
1. Run arena:

./bin/arena

>> Result:

Usage: arena [--version] [--help] <command> [<args>]

Available commands are:
    add    Runs add.


-------------------------------------
2. Get version:

./bin/arena --version

>> Result:

1.0.0

-------------------------------------
3. Run help:

./bin/arena --help add

>> Result:

Usage:
   arena add <number1> <number2>

Sample:
   arena add 1 2

-------------------------------------
4. Run add (OK, Scenario 1):

./bin/arena add 1 2

>> Result:

 1 + 2 = 3.00

-------------------------------------
5. Run add (OK, Scenario 2):

./bin/arena add 1,2 2.3

>> Result:

ERROR: Invalid number!

Usage:
   arena add <number1> <number2>

Sample:
   arena add 1 2


-------------------------------------
6. Run add (ERROR, Scenario 1):

./bin/arena add

>> Result:

ERROR: Invalid input parameters!

Usage:
   arena add <number1> <number2>

Sample:
   arena add 1 2


-------------------------------------
7. Run add (ERROR, Scenario 2):

./bin/arena add a b

>> Result:

ERROR: Invalid number!

Usage:
   arena add <number1> <number2>

Sample:
   arena add 1 2


-------------------------------------
```
