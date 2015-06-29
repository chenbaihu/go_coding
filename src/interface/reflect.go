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

		//得到实际的值(方法1)
		v1 := reflect.ValueOf((i).(*Person))
		fmt.Printf("reflect.ValueOf v1=%v\n", v1)

		//得到实际的值(方法2)
		v2 := reflect.ValueOf(i)
		fmt.Printf("reflect.ValueOf v2=%v\n", v2)

		if v1 == v2 {
			fmt.Printf("v1==v2\n")
		}

		//得到类型的元数据(方法1)
		t1 := reflect.TypeOf(i)
		fmt.Println("reflect.TypeOf t1=", t1)

		//得到类型的元数据(方法2)
		myref := v2.Elem()
		t2 := myref.Type()
		fmt.Println("reflect.TypeOf t2=", t2)

		for i := 0; i < myref.NumField(); i++ {
			fmt.Printf("i=%d\n", i)
			fields := myref.Field(i).String() // name := v2.Elem().Field(0).String()
			fmt.Printf("i=%d\tfields=%s\n", i, fields)

			field := myref.Field(i) // name := v2.Elem().Field(0)
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
