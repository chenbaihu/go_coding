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

		t1 := reflect.TypeOf(i) //得到类型的元数据
		fmt.Println("reflect.TypeOf t1=", t1)

		//v := reflect.ValueOf((i).(*Person)) //得到实际的值
		v := reflect.ValueOf(i) //得到实际的值
		fmt.Printf("reflect.ValueOf v=%v\n", v)

		myref := v.Elem()
		t2 := myref.Type() //得到类型的元数据
		fmt.Println("reflect.TypeOf t2:", t2)

		for i := 0; i < myref.NumField(); i++ {
			fmt.Printf("i=%d\n", i)
			fields := myref.Field(i).String() // name := v.Elem().Field(0).String()
			fmt.Printf("i=%d\tfields=%s\n", i, fields)

			field := myref.Field(i) // name := v.Elem().Field(0)
			//fmt.Printf("%d. %s %s = %v \n", i, t2.Field(i).Name, field.Type(), field.Interface())
			fmt.Printf("i=%d\tname=%s\ttype=%s\tfields=%s\n", i, t2.Field(i).Name, field.Type(), field.String())
		}

	}
}

func main() {
	var person Person
	person.name = "chenbaihu"
	person.age = 27
	show(&person)
}
