package main

import (
	"fmt"
	"os"
)

func main() {
	str := "hello 世界"

	n := len(str)
	fmt.Fprintf(os.Stdout, "str=%s\tlen=%d\n", str, n)

	for i := 0; i < n; i++ {
		fmt.Printf("for use len: i=%d\tstr[i]=%v\tstr[i]=%c\n", i, str[i], str[i])
	}

	for i, v := range str {
		fmt.Printf("for use range: i=%d\tv=%v\tv=%c\n", i, v, v)
	}

	rs := []rune(str)
	n = len(rs)
	for i := 0; i < n; i++ {
		fmt.Printf("rune str: i=%d\trs[i]=%v\trs[i]=%c\n", i, rs[i], rs[i])
	}

	s := str + " hello 中国"
	fmt.Printf("test + str: %s\n", s)

	fmt.Printf("test substr: %s\n", s[5:])
	fmt.Printf("test substr: %s\n", s[5:len(s)])
}
