GOPATH:=$(CURDIR)
export GOPATH

all: build
build: 
	go build -o bin/produce_demo            ./src/produce
	go build -o bin/consume_low_level_demo  ./src/consume_low_level
	go build -o bin/consume_group_demo      ./src/consume_group

clean:
	@rm -f bin/produce_demo
	@rm -f bin/consume_low_level_demo
	@rm -f bin/consume_group_demo

cleanlog:
	@rm -f log/*log*
