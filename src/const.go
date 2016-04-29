package main

import (
	"fmt"
)

func main() {
	const (
		kOK = iota // ENUM
		kERR
		kFAIL
	)
	println(kOK)
	println(kERR)
	fmt.Printf("%d\n", kFAIL)

	const (
		Type1 int = 0
		Type2 int = 2
		Type3 int = 4
		Type4 int = Type2 | Type3
	)

	const (
		str1 = "111111"
		str2 = "ccccccc"
		str3 = "dddddddd"
	)
	fmt.Printf("%s\n", str1)
}
