package main

import (
	"fmt"
)

func main() {
	//1
	array1 := [5]int{1, 2, 3, 4, 5}
	slices := array1[0:5]
	slices = append(slices, 7, 8, 9, 10, 11, 12)
	fmt.Println("slices", slices)

	//2
	slices1 := []int{1, 2, 3, 4, 5}
	slices1 = append(slices1, 7, 8, 9, 10)
	fmt.Println("slices1", slices1)

	//3
	slices2 := make([]int, 10)
	for i := 1; i < 10; i++ {
		slices2[i] = i
	}
	fmt.Println("slices2", slices2)

	slices2 = append(slices2, 15, 16, 17)
	fmt.Println("slices2", slices2)

	//4
	slices2 = append(slices2, slices1...)
	fmt.Println("slices2", slices2)

	//5
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	n1 := copy(s, a[0:])
	fmt.Println("n1", n1)
	fmt.Println("s", s)

	n2 := copy(s, s[2:])
	fmt.Println("n2", n2)
	fmt.Println("s", s)
}
