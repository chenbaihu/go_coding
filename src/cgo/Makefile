gcflag:=-gcflags "-N -l"
all:
	go build $(gcflag) go_references_c.go
	go build $(gcflag) go_references_c_use_define_function.go
	go build $(gcflag) go_references_clib.go

clean:
	@rm -rf go_references_c 
	@rm -rf go_references_c_use_define_function
	@rm -rf go_references_clib 
