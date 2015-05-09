package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	fmt.Printf("%s\n", req.Proto)
	fmt.Printf("%s\n", req.FormValue("a"))
	fmt.Printf("%s\n", req.UserAgent())
	fmt.Printf("%s\n", req.Referer())
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
