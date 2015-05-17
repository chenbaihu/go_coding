gflags := -gcflags "-N -l"
all:
	go build  $(gflags) tcp_echo.go 
	go build  $(gflags) finger_server.go 
	go build  $(gflags) tcpcli.go 

clean:
	@rm -rf tcp_echo.go
	@rm -rf finger_server
	@rm -rf tcpcli 


