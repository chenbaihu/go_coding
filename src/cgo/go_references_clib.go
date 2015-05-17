package main

import (
	"fmt"
	//"time"
)

// #include <stdlib.h>
// #include <stdio.h>
// #include <unistd.h>
// #include <stdint.h>
// #include <sys/types.h>
// #include <time.h>
import "C"

func random() uint32 {
	return uint32(C.random())
}

func main() {
	t := new(C.time_t)
	*t = C.time(t)
	C.srandom(C.uint(*t))
	//C.srand(C.uint(*t))
	for {
		fmt.Printf("This is from C:%d\n", random())
		//time.Sleep(time.Second)
	}
}
