package main

import (
	"bytes"
	"common/status"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"hash/crc32"
	"net"
	_ "os"
	"sync"
	_ "sync/atomic"
	"time"
)

var reqNum int
var hostPort string
var timeoutMs int
var duration int
var concurrence int

func initFlag() {
	flag.IntVar(&reqNum, "reqNum", 1, "Number of requests to perform")
	flag.StringVar(&hostPort, "hostPort", "127.0.0.1:1053", "The hostname and port of the udp server")
	flag.IntVar(&timeoutMs, "timeoutMs", 1, "The timeout MS")
	flag.IntVar(&duration, "d", 1, "The duration of dump stat info")
	flag.IntVar(&concurrence, "c", 1, "concurrence")

	flag.Parse()
}

// like ./udpcli -hostPort=127.0.0.1:1053 -reqNum=100000
func main() {
	initFlag()

	ts := status.NewTimerStatus()
	ts.SetTimerDump(time.Duration(duration)*time.Second, func() {
		b := bytes.NewBuffer([]byte{})
		ts.DumpCount(b)
		fmt.Printf("%v\n", string(b.Bytes()))
	})

	var wg sync.WaitGroup
	for c := 0; c < concurrence; c++ {
		go request(&wg, ts)
		wg.Add(1)
	}

	wg.Wait()
}

func makeData(body []byte) []byte {

	c32 := crc32.ChecksumIEEE(body)
	c32byte := new(bytes.Buffer)

	//0a 01 01 20 00 ee 00 01  00 00 6d 64
	_ = binary.Write(c32byte, binary.BigEndian, c32)
	cheader := []byte{0x00, 0x00, 0x00}
	cheader = append(cheader, c32byte.Bytes()[:]...)
	cheader = append(cheader, byte(0x00), byte(0x00), byte(0x00), byte(0x00))
	cheader = append(cheader, body[:]...)

	return cheader
}

func request(wg *sync.WaitGroup, ts *status.TimerStatus) {
	defer wg.Done()

	addr, err := net.ResolveUDPAddr("udp", hostPort)
	if err != nil {
		fmt.Println("server address error. It MUST be a format like this hostname:port", err)
		return
	}

	// Create a udp socket and connect to server
	socket, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Printf("connect to udpserver %v failed : %v\n", addr.String(), err.Error())
		return
	}
	defer socket.Close()

	data_tmp := "hello udp server"
	data_str, err := base64.StdEncoding.DecodeString(data_tmp)
	msg := make([]byte, len(data_str))
	msg = []byte(data_str)
	data := make([]byte, 1500)

	for i := 0; i < reqNum; i++ {

		t := time.Now()
		socket.SetDeadline(t.Add(time.Duration(1000 * time.Millisecond)))

		// send data to server
		_, err = socket.Write(msg)
		ts.AddCount("send")
		if err != nil {
			fmt.Println("send data error ", err)
			ts.AddCount("sendfail")
			continue
		}

		// recv data from server
		//read, remoteAddr, err := socket.ReadFromUDP(data)
		read, _, err := socket.ReadFromUDP(data)
		ts.AddCount("recv")
		if err != nil {
			fmt.Println("recv data error ", err)
			ts.AddCount("recvfail")
			continue
		}
		data = data[:read]

		//fmt.Println("RSP=%s", string(data))

		ts.AddCount("sendsucc")
		ts.AddCount("recvsucc")
	}
}
