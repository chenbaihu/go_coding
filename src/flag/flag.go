package main

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	// os 中的Args
	// like ./flag -async -main_conf=/user/local/nginx/conf/nginx.conf -port=8080 -percent=0.9 arg1 arg2 arg3
	fmt.Printf("Usage os args=%s -help or -h\n", os.Args[0])
	for _, args := range os.Args {
		fmt.Printf("Usage Show os args=%s\n", args)
	}
	fmt.Printf("os Args===================end\n")
	return
}

var Port int
var Percent float64
var MainConf string
var Async bool

func InitFlag() {
	// flag 中的Flag和Arg:
	// reference: http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/
	// ./flag -help or ./flag -h will show usage
	// flag.Xxx()，其中Xxx可以是Int、String等；返回一个相应类型的指针
	// flag.XxxVar()，将flag绑定到一个变量上
	// like ./flag -Port=8080 -Percent=0.9 -MainConf=/user/local/nginx/conf/nginx.conf -Async
	// 当遇到第一个non-falg参数将不再解析，例如 ./flag arg1 -async 此时async不会作为flag解析，而是作为参数
	flag.IntVar(&Port, "Port", 80, "http server listen port")
	flag.Float64Var(&Percent, "Percent", 0.8, "percent")
	flag.StringVar(&MainConf, "MainConf", "/home/chenbaihu/nginx.conf", "nginx.conf file")
	flag.BoolVar(&Async, "Async", false, "sync or async, default async")

	if !flag.Parsed() {
		flag.Parse()
	}
	// if like ./flag -Port=8080 -Percent=0.9 -MainConf=/user/local/nginx/conf/nginx.conf -Async will return 4
	fmt.Printf("InitFlag: there are %d flag input param\n", flag.NFlag())
}

func main() {
	Usage()
	InitFlag()
	ShowFlags()
	fmt.Printf("flag falg===================end\n")

	// like ./flag arg1 arg2 arg3
	// arg not start with "-"
	fmt.Printf("there are %d non-flag input param\n", flag.NArg()) // if ./flag arg1 arg2 arg3 will return 3
	fmt.Printf("%v\n", flag.Arg(0))                                // flag1
	for i, param := range flag.Args() {
		fmt.Printf("#%d :%s\n", i, param)
	}
	fmt.Printf("flag args===================end\n")
}
