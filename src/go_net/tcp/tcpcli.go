package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	//"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	reqArr := []string{"HEAD / HTTP/1.0\r\n\r\n", "GET / HTTP/1.1\r\n\r\n"}

	for _, req := range reqArr {

		_, err = conn.Write([]byte(req))
		checkError(err)

		result, err := ioutil.ReadAll(conn)
		checkError(err)

		fmt.Println("req:", req)
		fmt.Println(string(result))
	}

	//time.Sleep(time.Duration(100) * time.Second)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
