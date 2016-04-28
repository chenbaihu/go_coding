GOPATH:=$(CURDIR)
export GOPATH

all: build
build: 
	go build -o bin/produce_high_level_demo ./src/produce_high_level
	go build -o bin/consume_low_level_demo  ./src/consume_low_level

clean:
	@rm -f bin/produce_high_level_demo
	@rm -f bin/consume_low_level_demo

cleanlog:
	@rm -f log/*log*
