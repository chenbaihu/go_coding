package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type T func()

func timerDemo(input chan interface{}) {
	var count int64
	t1 := time.NewTimer(time.Second * 5)

	for {
		select {
		case msg := <-input:
			atomic.AddInt64(&count, int64(1))
			//println(msg)
			switch msg.(type) {
			case string:
				fmt.Printf("msg=%v\n", msg)
			case int:
				fmt.Printf("msg=%v\n", msg)
			case T:
				fmt.Printf("msg=%v\n", msg)
				//msg()
			default:
				fmt.Printf("msg=%v not know type\n", msg)
				//(T)(msg)()
			}

		case <-t1.C:
			fmt.Printf("5s timer, count=%d\n", atomic.SwapInt64(&count, int64(0)))
			t1.Reset(time.Second * 5)
		}
	}
}

func testFunc() {
	fmt.Printf("testFunc hello world")
}

func main() {
	// 虽然goroutine 是并发执行的，但是它们并不是并行运行的。如果不告诉Go 额外的东西，同一时刻只会有一个goroutine 执行。
	// 利用runtime.GOMAXPROCS(n) 可以设置goroutine 并行执行的数量。
	// GOMAXPROCS 设置了同时运行的CPU 的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
	// 当调度得到改进后，这将被移除。
	runtime.GOMAXPROCS(5)

	tick := time.Tick(1 * time.Second)

	input := make(chan interface{})
	go timerDemo(input)

	for {
		_ = <-tick
		input <- "this is a timer test"
		input <- 5
		input <- testFunc
	}
}
