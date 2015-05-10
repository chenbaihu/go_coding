package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

var c chan int

//ci := make(chan int )
//cs := make(chan string )
//cf := make(chan interface{})

//ci <- 1    //发送整数1 到channel ci
//<-ci       //从channel ci 接收整数
//i := <-ci  //从channel ci 接收整数，并保存到i 中

func ready(w string) {
	for true {
		i, ok := <-c
		if ok {
			fmt.Printf("ready chan i:%d\tw:%s\n", i, w)
			c <- i
		}
	}
}

func main() {
	fmt.Println("channel test")

	// 虽然goroutine 是并发执行的，但是它们并不是并行运行的。如果不告诉Go 额外的东西，同一时刻只会有一个goroutine 执行。
	// 利用runtime.GOMAXPROCS(n) 可以设置goroutine 并行执行的数量。
	// GOMAXPROCS 设置了同时运行的CPU 的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
	// 当调度得到改进后，这将被移除。
	runtime.GOMAXPROCS(5)

	c = make(chan int)
	go ready("Tea")
	go ready("Coffee")

	for true {
		i := rand.Int()
		c <- i
		r, ok := <-c
		if ok {
			fmt.Printf("main r:%d\n", r)
		}
	}
	close(c)
}
