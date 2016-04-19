GOPATH:=$(CURDIR)
export GOPATH

all: build

install:
	go install go_thrift_cli 

build: build-go_thrift_cli

build-go_thrift_cli:
	go build -o bin/go_thrift_cli ./src 

clean:
	@rm -f bin/go_thrift_cli

cleanlog:
	@rm -f log/*log*
