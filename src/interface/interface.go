package main

import (
	"fmt"
)

type S struct {
	i int
}

//方法定义规则:
//规则1， 不能定义新的方法 在非本地类型int 上
//规则2， 不能定义新的方法 例如：不能在非本地类型net.AddrError 上定义方法
//接收者类型必须是T 或*T，这里的T 是类型名。T 叫做接收者基础类型或 简称基础类型。基础类型一定不能使指针或接口类型，并且定义在与方法 相同的包中。这里S就是T
func (p *S) Put(val int) {
	p.i = val
}

func (p *S) Get() int {
	fmt.Println("*S Get Method Enter")
	return p.i
}

//func (p S) Put(val int) {
//	p.i = val // 不能修改实际调用者的值，p S是值传递，这里只是为了测试
//}
//
//func (p S) Get() int {
//	fmt.Println("S Get Method Enter")
//	return p.i
//}

type R struct {
	f int
}

//方法定义规则:
//规则1， 不能定义新的方法 在非本地类型int 上
//规则2， 不能定义新的方法 例如：不能在非本地类型net.AddrError 上定义方法
func (r *R) Put(val int) {
	r.f = val
}

func (r *R) Get() int {
	fmt.Println("*R Get Method Enter")
	return r.f
}

//func (r R) Put(val int) {
//	r.f = val // 不能修改实际调用者的值，r R是值传递，这里只是为了测试
//}
//
//func (r R) Get() int {
//	fmt.Println("R Get Method Enter")
//	return r.f
//}
//
//接口定义：
//接口定义为一个方法的集合。方法包含实际的代码。换句话说，一个接口就是定义， 而方法就是实现。
//因此，接收者不能定义为接口类型，这样做的话会引起invalid receiver type ... 的编译器错误。
type I interface {
	Put(int)
	Get() int
}

//接口的使用:
func f(p I) {
	switch p.(type) {
	case *S:
		fmt.Println("type is *S")
		break
	case *R:
		fmt.Println("type is *R")
		break
	//case S:
	//	fmt.Println("type is S")
	//	break
	//case R:
	//	fmt.Println("type is R")
	//	break
	default:
		fmt.Println("type is UNKNOW")
		break
	}
	p.Get()
}

func main() {
	var s S
	//f(s)
	f(&s) // 获取s 的地址，而不是s 的值的原因，是因为在s 的指针上定义了方法

	var r R
	//f(r)
	f(&r) // 获取r 的地址，而不是r 的值的原因，是因为在r 的指针上定义了方法
}
