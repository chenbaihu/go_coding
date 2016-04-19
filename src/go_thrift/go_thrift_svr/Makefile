GOPATH:=$(CURDIR)
export GOPATH

all: build

install:
	go install go_thrift_svr 

build: build-go_thrift_svr

build-go_thrift_svr:
	go build -o bin/go_thrift_svr ./src 

clean:
	@rm -f bin/go_thrift_svr

cleanlog:
	@rm -f log/*log*
