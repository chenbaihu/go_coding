package main

import (
	"fmt"
	"strconv"
)

//interface{}  空interface
//空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类
//例如:
//var args []interface{}
//args = []interface{}{"a", 5, 5.5}

type List []interface{}

type Person struct {
	name string
	age  int
}

//打印
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	//
	var args []interface{}
	args = []interface{}{"a", 5, 5.5}
	fmt.Printf("args=%v\n", args)

	//
	var args2 []int
	args2 = []int{5, 10, 1000}
	fmt.Printf("args2=%v\n", args2)

	//
	list := make(List, 3)
	list[0] = 1       //an int
	list[1] = "Hello" //a string
	list[2] = Person{"Dennis", 70}

	// if else
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Println("list[%d] is of a different type", index)
		}
	}

	// switch
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
