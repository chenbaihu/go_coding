package main

import (
	"fmt"
	//"io"
	"os"
	"strings"
)

func main() {
	buf := make([]byte, 4096)

	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Printf("os.Stdin Read failed")
		return
	}
	fmt.Printf("os.Stdin Read n=%d\tbuf=%s", n, strings.Trim(string(buf), "\n"))

	os.Stdout.Write(buf)
}
