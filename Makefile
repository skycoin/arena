default: clean format build test

# All code should be formatted with goimports.
# Use the testify require or testify assert package for assertions.
update_packages: 
	dep ensure -update -v

#  Additionally, the project should have a Makefile command called
#  make format which formats all of the source with goimports.
format: update_packages
	goimports -w .

build: update_packages
	go build

test:
	go test -v	

clean:
	go clean
	