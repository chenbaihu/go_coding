package main

import (
	"fmt"
	"strings"
)

//func MapTestFun(r []rune) []rune {
//	if r == []rune("||") {
//		return []rune("|")
//	}
//
//	if r == []rune("/") {
//		return []rune("|")
//	}
//	return r
//}

func MapTestFun(r rune) rune {
	if r == rune('/') {
		return rune('|')
	}
	return r
}

func main() {
	str := "a|b|c|d|e|f"
	sa := strings.Split(str, "|") // Split
	fmt.Printf("sa=%v\n", sa)

	for i := 0; i < len(sa); i++ {
		fmt.Printf("%s\n", sa[i])
	}

	str = strings.Join(sa, "||") // Join
	fmt.Printf("%s\n", str)

	str = strings.Replace(str, "||", "/", 2) // Replace
	fmt.Printf("%s\n", str)

	str = strings.Map(MapTestFun, str) // Map
	fmt.Printf("%s\n", str)
}
