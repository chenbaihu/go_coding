package main

import (
	"fmt"
	"strconv"
	"sync"
)

type DataInfo struct {
	v    int
	data string
}

var queue = make(chan *DataInfo, 100000)
var quit chan int

var dataPool *sync.Pool
var producti int64

func init() {
	dataPool = &sync.Pool{New: func() interface{} {
		d := DataInfo{v: 1, data: "hello world"}
		return &d
	}}
	producti = 0
}

func product() {
	for {
		d := dataPool.Get().(*DataInfo)
		d.v = 2
		d.data = "hello world" + strconv.FormatInt(producti, 10)
		queue <- d
		producti++
		fmt.Printf("product d=%v\n", d)
		if producti > 100000000 {
			quit <- 1
		}
	}
}

func consume() {
	for {
		d := <-queue
		fmt.Printf("consume d=%v", d)
		dataPool.Put(d)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go product()
	}

	for i := 0; i < 5; i++ {
		go consume()
	}
	_ = <-quit
}
