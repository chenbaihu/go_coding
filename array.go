package main

import (
	"fmt"
)

func main() {
	//1
	var arr [100]int
	for i := 0; i < 100; i++ {
		arr[i] = i
	}

	//2
	arr2 := [3][2]int{[...]int{1, 2}, [...]int{3, 4}, [...]int{5, 6}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\n", arr2[i][j])
		}
	}

	//3
	arr3 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\n", arr3[i][j])
		}
	}

	arr4 := [3]string{"aaaaaa", "bbbbb", "cccccc"}
	for i := 0; i < 3; i++ {
		println(arr4[i])
	}
}
