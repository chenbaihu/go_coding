GOPATH:=$(CURDIR)
export GOPATH

all: build
build: 
	go build -o bin/kv_demo ./src/main

clean:
	@rm -f bin/kv_demo

cleanlog:
	@rm -f log/*log*
