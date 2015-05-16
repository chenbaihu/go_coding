gcflag:=-gcflags "-N -l"
all:
	go build $(gcflag) timer_select.go
	go build $(gcflag) timer_after.go
	go build $(gcflag) timer_channel.go

clean:
	@rm -rf timer_select 
	@rm -rf timer_after
	@rm -rf timer_channel
