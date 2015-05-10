package main

import (
	"fmt"
)

type S struct {
	i int
}

func (p *S) Put(val int) {
	p.i = val
}

func (p *S) Get() int {
	fmt.Println("S Get Method Enter")
	return p.i
}

type R struct {
	f int
}

func (r *R) Put(val int) {
	r.f = val
}

func (r *R) Get() int {
	fmt.Println("R Get Method Enter")
	return r.f
}

type I interface {
	Put(int)
	Get() int
}

func f(p I) {
	switch p.(type) {
	case *S:
		fmt.Println("type is *S")
		p.Get()
		break
	case *R:
		fmt.Println("type is *R")
		p.Get()
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
}

func main() {
	var s S
	//f(s) // 在s 上没有定义方法,接口也就不可以用
	f(&s) // 获取s 的地址，而不是s 的值的原因，是因为在s 的指针上定义了方法

	var r R
	//f(r) // 在r 上没有定义方法，接口也就不可以使用
	f(&r) // 获取r 的地址，而不是r 的值的原因，是因为在r 的指针上定义了方法
}
