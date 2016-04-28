GOPATH:=$(CURDIR)
export GOPATH

all: build

install:
	go install proto3_map 

build: build-go_proto3_map

build-go_proto3_map:
	go build -o bin/go_proto3_map ./src/proto3_map

clean:
	@rm -f bin/go_proto3_map

cleanlog:
	@rm -f log/*log*
