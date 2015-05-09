package main

import (
	"fmt"
)

func main() {
	const (
		kOK = iota
		kERR
		kFAIL
	)
	println(kOK)
	println(kERR)
	fmt.Printf("%d\n", kFAIL)

	const (
		str1 = "111111"
		str2 = "ccccccc"
		str3 = "dddddddd"
	)
	fmt.Printf("%s\n", str1)
}
