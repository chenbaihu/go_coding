package main

import (
	"fmt"
)

///////////////// multi return values
func MultiReturnTest(v1 int, v2 int) (min int, max int) {
	if v1 <= v2 {
		min = v1
		max = v2
		return
	}
	min = v2
	max = v1
	return
}

/////////////////
type Test struct {
	l   int
	arr []int
}

////////////////   可变参数 及 method 及 延迟调用defer
func (t *Test) InitTest(l int, arg ...int) bool {
	defer fmt.Println("InitTest Enter And Return")
	//变量arg 是一个int 类型的slice
	t.arr = append(t.arr, arg...) // append auto allocation space, can't use copy
	fmt.Println("arr:", t.arr)
	t.l = l
	return true
}

func (t Test) ShowTest() {
	fmt.Println("Test.l:", t.l)
	fmt.Println("Test.arr:", t.arr)
}

//////////////////// function type
type Func func(int)
type MapFunc map[string]Func

func hello(val int) {
	fmt.Println("hello world:", val)
}

///////////////////  callback
func callback(val int, fun Func) {
	fun(val)
}

func main() {
	min, max := MultiReturnTest(4, 5)
	fmt.Printf("min=%d\tmax=%d\n", min, max)

	min, max = MultiReturnTest(5, 4)
	fmt.Printf("min=%d\tmax=%d\n", min, max)

	t := new(Test) // eq  var t Test
	if !t.InitTest(5, 1, 2, 3, 4, 5) {
		fmt.Println("t.InitTest failed")
	}
	t.ShowTest()

	index := 1

	var f Func
	f = hello
	f(index)
	index++

	//m_fc := MapFunc{
	m_fc := map[string]Func{
		"func1": hello,
		"func2": hello,
		"func3": hello,
	}

	m_fc["func1"](index)
	index++

	for fn, fc := range m_fc {
		fmt.Printf("fn:%s\n", fn)
		fc(index)
		index++
	}

	callback(index, hello)
}
