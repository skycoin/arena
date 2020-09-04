# Running building and running the cli-tool

## Cloning to your system

Clone or download the repository to your GOAPTH

```bash
cd $GOPATH
````

Enter the following folder

```bash
cd src\github.com\skycoin
```
** or mkdir the same

then:

```bash
git clone address\to\this\repo`
```

## Running the command

### Build

to build and check dependencies :
```bash
make all
```
to build the tool standalone:

```bash
make arena
```

### Available commands

* help
```bash
$ arena help
a small cli tool with one `add` command

Usage:
  arena [flags]
  arena [command]

Available Commands:
  add         Commnad used to perform addition
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.add.yaml)
  -h, --help            help for arena
  -t, --toggle          Help message for toggle

Use "arena [command] --help" for more information about a command.
```

* add
  without -f flag uses integers
  with -f flag uses float64

** NOTE: I did not manage to check the float64 for overflows ( still figuring that out :) )

```
$ arena add
Please pass some arguments to add

This command adds the arguments passed following it.
Usage: arena add arg1 arg2 arg3


$ arena add 231 222
Addition of given numbers [231 222] is 453

$ arena add 231 three
error:
strconv.Atoi: parsing "three": invalid syntax while parsing integers, omiting value : three
Addition of given numbers [231] is 231

$ arena add .4 .5 -f
Sum of given float numbers [.4 .5] is 0.900000

```



