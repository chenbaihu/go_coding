/*
   1、启动work协程池
   2、启动http服务
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	_ "os"
	"runtime"
	"runtime/debug"
	"strings"
	_ "sync"
)

type DataInfo struct {
	Logid string
	Cmd   string
	Data  string
}

var reqMapChan = make(chan *DataInfo, 1000000)

type HttpHandler struct {
	Name     string
	Callfunc func(w http.ResponseWriter, r *http.Request, logid string)
}

func (hh *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remoteaddr := strings.Split(r.RemoteAddr, ":")
	remoteip := remoteaddr[0]
	inputIp := r.Header.Get("X-Forwarded-For")
	if inputIp == "" {
		inputIp = remoteip
	}

	//sid := r.URL.Query().Get("sid")

	logid := "logid"

	hh.Callfunc(w, r, logid)
}

func handle_kv_multi_set(w http.ResponseWriter, r *http.Request, logid string) {
	data := r.FormValue("data")
	req := &DataInfo{Logid: logid, Cmd: "mset", Data: data}
	reqMapChan <- req
	io.WriteString(w, "req process ok--------------")
}

func handle_kv_multi_get(w http.ResponseWriter, r *http.Request, logid string) {
	data := r.FormValue("data")
	req := &DataInfo{Logid: logid, Cmd: "mget", Data: data}
	//reqMapChan <- req
	//TODO req
	//阻塞给响应
	io.WriteString(w, "req process ok--------------")
}

func MapWorker() {
	for {
		select {
		case data := <-reqMapChan:
			//dealworkdata(data)
			fmt.Printf("data=%v\n", data)
			//io.WriteString(*data.Rsp, "req process ok--------------")
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// set recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("abort, unknown error, reason:%v, stack:%s\n",
				err, string(debug.Stack()))
		}
	}()

	//pprof for debug
	go func() {
		err := http.ListenAndServe(":9696", nil)
		if err != nil {
			fmt.Printf("err=%v\n", err)
			//TODO err
		}
	}()

	// start woker pool
	for i := 0; i < 5; i++ {
		go MapWorker()
	}

	// start http server
	mux := http.NewServeMux()
	mux.Handle("/kv/multi_get", &HttpHandler{Name: "handle_kv_multi_get", Callfunc: handle_kv_multi_get})
	mux.Handle("/kv/multi_set", &HttpHandler{Name: "handle_kv_multi_set", Callfunc: handle_kv_multi_set})

	err := http.ListenAndServe(":12345", mux)
	fmt.Printf("http listen fail: %s\n", err.Error())
}
