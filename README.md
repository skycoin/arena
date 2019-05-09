


# How to Run Skycoin Wallet From Source (In (Terminal)

Commands to use to run your Skycoin Wallet from source in terminal. 
## Getting Started

The following commands will get Skycoin Wallet running from terminal / command line.

### Prerequisites

The assumption is that you already have the Skycoin Installation package installed and tested for your version of O/S. If not the following documents will guide you through that process: 
[https://github.com/skycoin/skycoin#run-skycoin-from-the-command-line](https://github.com/skycoin/skycoin#run-skycoin-from-the-command-line)

[https://github.com/skycoin/skycoin/blob/develop/INSTALLATION.md](https://github.com/skycoin/skycoin/blob/develop/INSTALLATION.md)

## Running Skycoin Wallet

The following commands should be used to run the Skycoin wallet :

$ cd $GOPATH/src/github.com/skycoin/skycoin
$ make run-client

## Show Skycoin Node Options

The following commands will display the Skycoin Node options: 

$ cd $GOPATH/src/github.com/skycoin/skycoin
$ make run-help

## Run Skycoin With Options

The following commands will allow you to run Skycoin Wallet with options: 

$ cd $GOPATH/src/github.com/skycoin/skycoin
$ make ARGS="--launch-browser=false -data-dir=/custom/path" run


## End

