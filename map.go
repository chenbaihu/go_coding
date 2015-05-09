package main

import (
	"fmt"
)

func main() {
	//map[<from type>]<to type>

	// create map 1
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Println("monthdays:", monthdays)

	// create map 2
	//当只需要声明一个map 的时候，使用make 的形式：
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3
	fmt.Println("make map:", m)

	// modify map
	monthdays["Feb"] = 29 //modify

	// add elment
	monthdays["Unknow"] = 0 //change

	// foreach
	for k, v := range monthdays {
		fmt.Printf("%s\t%d\n", k, v)
	}

	//一个特殊的变量名是_（下划线）。任何赋给它的值都被丢弃
	for _, v := range monthdays {
		fmt.Printf("%d\n", v)
	}

	// find
	v, ok := monthdays["Feb"]
	if ok {
		fmt.Println("Feb:", v)
	}
	v, ok = monthdays["F"]
	if !ok {
		fmt.Println("F Not Find")
	}

	// delete
	delete(monthdays, "Unknow") // delete(m, x) 会删除map 中由m[x] 建立的实例
	//delete(monthdays, "F")
	fmt.Println("monthdays:", monthdays)

	// copy
	mon2 := monthdays
	fmt.Println("mon2:", mon2)
}
