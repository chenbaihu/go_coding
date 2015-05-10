package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int32
}

func show(i interface{}) {
	switch i.(type) {
	case *Person:

		t := reflect.TypeOf(i) //得到类型的元数据
		fmt.Println("reflect.TypeOf t:", t)

		v := reflect.ValueOf(i) //得到实际的值
		fmt.Printf("reflect.ValueOf v:%v\n\n\n", v)

		//fmt.Println("NumField=", v.NumField())

		//for index := 0; index < v.NumField(); index++ {
		for index := 0; index < 2; index++ {
			tag := t.Elem().Field(index).Tag
			fmt.Println("t.Elem().Field tag:", index, tag)

			filed := v.Elem().Field(index)
			//filed:= v.Elem().Field(index).String()
			fmt.Println("v.Elem().Field", index, "\t", filed)
		}
	}
}

func main() {
	var person Person
	person.name = "chenbaihu"
	person.age = 27
	show(&person)
}
