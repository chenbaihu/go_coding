package main

import (
    _ "errors"
    "fmt"
    _ "net/http"
    "os"
    "thriftclient"
    "runtime"
    _ "time"
    "flag"
    "idl/gen-go/rt"
)

var (
    RTIp string
    RTPort int
    RTTimeout int
)

func InitFlag() {
    flag.StringVar(&RTIp,   "Ip",        "127.0.0.1", "rtproxy server ip addr")
    flag.IntVar(&RTPort,    "Port",      11311,       "rtproxy listen port")
    flag.IntVar(&RTTimeout, "TimeOutMS", 100,         "rtproxy timeout ms")
    if !flag.Parsed() {
        flag.Parse()
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    InitFlag()
    fmt.Printf("ip=%s port=%d timeout_ms=%d\n", RTIp, RTPort, RTTimeout);

    req := &rt.ComputeReq{TypeA1:1, JobId:100, CityId:5, MapsplitId:1}
    req.OidList = append(req.OidList, 1, 2, 3)
    req.DidList = append(req.DidList, 10, 20, 30)

    cli, transport, err := thriftclient.GetRTClient(RTIp, RTPort, (int64)(RTTimeout))  
    if err!=nil { 
        fmt.Printf("GetRTClient err=%s\n", err);
        os.Exit(-2) 
    }
    defer transport.Close()

    rsp, err2 := cli.Compute(req)
    if err2!=nil {
        fmt.Printf("Compute err=%s\n", err2);
        os.Exit(-3) 
    }
    fmt.Printf("rsp=%v\n", rsp)
}
