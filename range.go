package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d", "e", "f"}

	for k, v := range list {
		fmt.Printf("%d\t%s\n", k, v)
	}
}
