package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Go 的I/O 核心是接口io.Reader 和io.Writer。
	buffer := make([]byte, 1024)
	l, _ := os.Stdin.Read(buffer)
	if l == 0 {
		fmt.Println("os.Stdin.Read failed")
		return
	}
	filename := strings.Trim(string(buffer[:l]), string("\n"))
	fmt.Printf("os.Stdin.Read succ, filename:%s\n", filename)

	buf := make([]byte, 1024)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("os.Open filename:", filename, "failed")
		return
	}
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
