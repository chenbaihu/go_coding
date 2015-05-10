package main

import (
	"fmt"
	"math/rand"
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
