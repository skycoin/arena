# README: RUNNING SKYCOIN WALLET IN TERMINAL

 Running the Skycoin Wallet in Linux environments is fairly simple. To do so, you will first need to install Go1.10+, then retrieve Skycoin with Go.

**Step 1:**

On OSX (Paste these commands in a MacOS Terminal Prompt): 
	1. Install homebrew.
		
    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
2. Install the latest version of Go.
`brew install go`
3. Install Mercurial and Bazaar.
`brew install mercurial bzr`

On Linux (Paste these commands in Terminal)
	Ubuntu/Debian:
	
1. Paste the following commands to install Go, Mercurial and Bazaar.
		
		sudo apt-get update && sudo apt-get upgrade -y
		sudo apt-get install -y curl git mercurial make binutils gcc bzr bison libgmp3-dev screen gcc build-essential
Centos/Fedora

1. Paste the following commands to install Go, Mercurial and Bazaar.
	
		sudo yum update -y && sudo yum upgrade -y
		sudo yum install -y git curl make gcc mercurial binutils bzr bison screen
		if [[ "$(cat /etc/redhat-release | grep -o CentOS)" == "CentOS" ]]; then sudo yum install -y build-essential libgmp3-dev; else sudo yum groupinstall -y "Development Tools" "Development Libraries" && sudo yum install -y gmp; fi;

Manual installation on Linux/Unix Operating Systems is also an option, using the following commands:

1. Declare which version of Go you wish to download.
	
		cd ~
		export GOV=1.11.1 # golang version
2. Use the following commands to download and uncompress the Golang source.

		curl -sS https://storage.googleapis.com/golang/go$GOV.linux-amd64.tar.gz > go$GOV.linux-amd64.tar.gz
		tar xvf go$GOV.linux-amd64.tar.gz
		rm go$GOV.linux-amd64.tar.gz	
3. Install Go.

			sudo mv go /usr/local/go
			sudo ln -s /usr/local/go/bin/go /usr/local/bin/go
			sudo ln -s /usr/local/go/bin/godoc /usr/local/bin/godoc
			sudo ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt

4. The $GOPATH environment variable specifies the location of your workplace. By default, this is a directory named go, with a path of $HOME/go on Unix. You now need to specify your workplace directory's inner folders with the following commands.
				
			mkdir -p $HOME/go
			mkdir -p $HOME/go/bin
			mkdir -p $HOME/go/src
			mkdir -p $HOME/go/pkg
5. After setting up your $GOPATH variable, add it to ~/.bashrc.
		
			source ~/.bashrc
			
6. Export these paths.
	
			export GOROOT=/usr/local/go
			export GOPATH=$HOME/go
			export GOBIN=$GOPATH/bin
			export PATH=$PATH:$GOBIN
7. Test if you installed Go successfully using the hello.go application described below.
	https://golang.org/doc/install#testing

**Step 2:**

1. Retrieve the Skycoin with Go. This will download the Skycoin git to $GOPATH/src/github.com/skycoin/skycoin. If this doesn't work, you can clone the repo directly.
	
		go get github.com/skycoin/skycoin/cmd/...
	Direct Clone: 
		
		git clone https://github.com/skycoin/skycoin
2. Run Skycoin from the command line.
	
			cd $GOPATH/src/github.com/skycoin/skycoin
			make run-client
Now you can optionally:

1. Show the Skycoin node options.
	
			cd $GOPATH/src/github.com/skycoin/skycoin
			make run-help
2.  Run Skycoin with options.
			
			cd $GOPATH/src/github.com/skycoin/skycoin
			make run-help

