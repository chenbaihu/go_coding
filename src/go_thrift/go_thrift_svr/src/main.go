package main

import (
    _ "errors"
    "fmt"
    _ "net/http"
	"git.apache.org/thrift.git/lib/go/thrift"
    "os"
    "thriftserver"
    "runtime"
    "time"
    "flag"
    "idl/gen-go/rt"
	"strconv"
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

    handle := thriftserver.NewMapServiceHandle()

    tf := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    pf := thrift.NewTBinaryProtocolFactoryDefault()

    taddr := RTIp + ":" + strconv.Itoa(RTPort)

    ss, err := thrift.NewTServerSocketTimeout(taddr, time.Millisecond*time.Duration(RTTimeout))
    if err != nil {
        fmt.Printf("taddr=%s RTTimeout=%d NewTServerSocketTimeout failed", taddr, RTTimeout)
        os.Exit(-2) 
    }   

    processor := rt.NewMapServiceProcessor(handle)
    server := thrift.NewTSimpleServer4(processor, ss, tf, pf) 
    
    fmt.Printf("taddr=%s RTTimeout=%d succ", taddr, RTTimeout)

    server.Serve()
    os.Exit(0) 
}
