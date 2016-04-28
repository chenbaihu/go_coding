GOPATH:=$(CURDIR)
export GOPATH

all: build
build: 
	go build -o bin/produce_demo            ./src/produce
	go build -o bin/consume_low_level_demo  ./src/consume_low_level

clean:
	@rm -f bin/produce_demo
	@rm -f bin/consume_low_level_demo

cleanlog:
	@rm -f log/*log*
