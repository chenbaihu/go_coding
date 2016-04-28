GOPATH:=$(CURDIR)
export GOPATH

all: build

install:
	go install produce_demo

build: build-go_produce_demo

build-go_produce_demo:
	go build -o bin/produce_demo ./src/produce_demo

clean:
	@rm -f bin/produce_demo

cleanlog:
	@rm -f log/*log*
