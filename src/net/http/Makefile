gflags := -gcflags "-N -l"
all:
	go build  $(gflags) tcp_echo.go 
	go build  $(gflags) http_echo.go 
	go build  $(gflags) finger_server.go 

clean:
	@rm -rf tcp_echo.go
	@rm -rf http_echo.go
	@rm -rf finger_server


