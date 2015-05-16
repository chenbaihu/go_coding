package main

import (
	"fmt"
	"runtime"
	"time"
)

func runAfter() {
	fmt.Println("run After")
}

func main() {
	// 虽然goroutine 是并发执行的，但是它们并不是并行运行的。如果不告诉Go 额外的东西，同一时刻只会有一个goroutine 执行。
	// 利用runtime.GOMAXPROCS(n) 可以设置goroutine 并行执行的数量。
	// GOMAXPROCS 设置了同时运行的CPU 的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
	// 当调度得到改进后，这将被移除。
	runtime.GOMAXPROCS(5)

	timer := new(time.Timer)
	timer = time.AfterFunc(1*time.Second, func() {
		runAfter()
		timer.Reset((time.Duration)(1 * time.Second))
	})

	time.Sleep(time.Hour)
}
