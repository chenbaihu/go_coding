package main

// #include <stdlib.h>
// #include <stdio.h>
// #include <unistd.h>
// #include <stdint.h>
// #include <sys/types.h>
// #include <time.h>
// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"

import (
	"./fun"
	"./log"
	"fmt"
	"os"
	"runtime"
	"time"
)

// http://golang.org/cmd/cgo/
// http://blog.sina.com.cn/s/blog_538d55be01015h6g.html

func random() uint32 {
	t := new(C.time_t)
	*t = C.time(t)
	C.srandom(C.uint(*t))
	//C.srand(C.uint(*t))

	for {
		fmt.Printf("This is from C:%d\n", C.random())
		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Printf("This is from C:%s\n", fun.MySecret())
	go func() {
		for {
			str := "hello world go string"
			fun.HelloWorld(str)
			time.Sleep(5 * time.Second)
		}
	}()

	go random()

	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))

	logName := "mylog.mylog_test"
	mylog := log.NewMyLog(logName)
	defer log.CloseMyLog(mylog)

	md5 := "b03c700c5241bc2ec49578ab19b3355f"
	num := 123456
	cv := "win1359"

	_, filename, line, _ := runtime.Caller(0)
	mylog.LogDebug(logName, filename, line, "%s\t%d\t%s", md5, num, cv)

	time.Sleep(time.Hour)
	os.Exit(0)
}
