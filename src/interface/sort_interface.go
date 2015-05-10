package main

import (
	"fmt"
)

type Xi []int

func (xi Xi) Len() int {
	return len(xi)
}

func (xi Xi) Less(i int, j int) bool {
	if xi[i] <= xi[j] {
		return true
	}
	return false
}

func (xi Xi) Swap(i int, j int) {
	xi[i], xi[j] = xi[j], xi[i]
}

type Xs []string

func (xs Xs) Len() int {
	return len(xs)
}

func (xs Xs) Less(i int, j int) bool {
	if xs[i] <= xs[j] {
		return true
	}
	return false
}

func (xs Xs) Swap(i int, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

//根据规则，单方法接口命名为方法名加上-er 后缀：Reader，Writer，Formatter 等。
//有一堆这样的命名，高效的反映了它们职责和包含的函数名。Read，Write，Close， Flush，String 等等有着规范的声明和含义
type Sorter interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func Sort(x Sorter) {
	for i := 0; i < x.Len(); i++ {
		for j := i + 1; j < x.Len(); j++ {
			if !x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

	fmt.Println("ints:", ints)
	Sort(ints)
	fmt.Println("Sort ints:", ints)

	fmt.Println("strings:", strings)
	Sort(strings)
	fmt.Println("Sort strings:", strings)
}
