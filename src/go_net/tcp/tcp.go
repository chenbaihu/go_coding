package main

import (
	"bytes"
	"fmt"
	//"io/ioutil"
	"net"
	"os"
	//"strings"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"goapp/lcsdispatcher/stat"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func createTcpCli(timeout time.Duration, tcpaddr string) (net.Conn, error) {
	//fmt.Println("createTcpCli timeout ", timeout)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func tcpSendRecv(conn net.Conn, data []byte, timeout time.Duration) ([]byte, error) {
	totalLens := 0
	for totalLens < len(data) {
		lens, err := conn.Write(data[totalLens:])
		if err != nil {
			fmt.Println("conn Write ", err.Error())
			//conn.Close()
			return nil, err
			//continue
		}
		totalLens += lens
	}

	head := make([]byte, 4)
	_, err := conn.Read(head)
	if err != nil {
		fmt.Println("conn Read Head", err.Error())
		//conn.Close()
		return nil, err
	}

	var rspDataLen uint32
	headBuf := bytes.NewReader(head)
	//err = binary.Read(headBuf, binary.LittleEndian, &rspDataLen)
	err = binary.Read(headBuf, binary.BigEndian, &rspDataLen)
	if err != nil {
		fmt.Println("rspDataLen ", err.Error())
	}
	//fmt.Printf("====================rspDataLen=%d=========================\n", rspDataLen)

	body := make([]byte, rspDataLen)
	var totalLenr uint32 = 0
	for totalLenr < rspDataLen {
		lenr, err := conn.Read(body[totalLenr:])
		if err != nil {
			fmt.Println("conn Read Body", err.Error())
			//conn.Close()
			return nil, err
			//continue
		}
		totalLenr += (uint32)(lenr)
	}
	//fmt.Printf("====================rspDataLen=%u\tbody=%v\n", rspDataLen, body[0:totalLenr])
	return body[0:totalLenr], nil
}

func makeReqData() []byte {
	body := bytes.NewBufferString(inputData)
	if len(inputData) == 0 {
		for i := 0; i < dataLen; i++ {
			body.WriteByte(byte(i))
		}
	}

	head := bytes.NewBufferString("")

	if len(inputData) == 0 {
		//err := binary.Write(head, binary.BigEndian, dataLen)
		_ = binary.Write(head, binary.BigEndian, (uint32)(dataLen))
	} else {
		_ = binary.Write(head, binary.BigEndian, (uint32)(len(inputData)))
	}

	data := append(head.Bytes(), body.Bytes()[:]...)
	return data

	//data := bytes.NewBufferString("")
	//data.WriteString(head.String())
	//data.WriteString(body.String())
	//return data.Bytes()
}

func tcpTest(wg *sync.WaitGroup, timerStatHelper *stat.StatHelper) {
	defer wg.Done()
	reqData := makeReqData()

	//if disPlay {
	//	fmt.Printf("tcp reqData=%s\n", base64.StdEncoding.EncodeToString(reqData[4:]))
	//}
	for {
		if (reqNum != 0) && (atomic.LoadInt64(&hasSendNum) > reqNum) {
			//finish
			break
		}

		beg := time.Now()

		tcpCli, err := createTcpCli(time.Duration(timeoutMs)*time.Millisecond, hostPort)
		if err != nil {
			//TOTO
		}
		for j := 0; j < reqNumPerConn; j++ {
			if atomic.LoadInt64(&hasSendNum) >= reqNum {
				os.Exit(0)
				break
			}

			atomic.AddInt64(&hasSendNum, 1)

			timerStatHelper.AddCount("tcpreq")
			rspData, err := tcpSendRecv(tcpCli, reqData, time.Duration(timeoutMs)*time.Millisecond)
			if err != nil {
				fmt.Println("tcp req faile:", err)
				timerStatHelper.AddCount("tcpreqfail")
				break
			}
			if bytes.Compare(reqData[4:], rspData) == 0 {
				timerStatHelper.AddCount("tcpreqsucc")
				if disPlay {
					fmt.Printf("tcp reqDataBase64=%s\trspDataBase64=%s\n", base64.StdEncoding.EncodeToString(reqData[4:]), base64.StdEncoding.EncodeToString(rspData))
				}
			} else {
				timerStatHelper.AddCount("tcpreqfail2")
			}
		}
		tcpCli.Close()

		costMs := float64(time.Since(beg).Nanoseconds()) / 1000000.0
		fmt.Printf("costMs=%f \n", costMs)

		if disPlay {
			fmt.Printf("================================tcp finish reqNumPerConn=%d close conn\n", reqNumPerConn)
		}
	}
}

var reqNum int64
var reqNumPerConn int
var hostPort string
var timeoutMs int
var duration int
var concurrence int
var dataLen int
var disPlay bool
var inputData string

func initFlag() {
	flag.Int64Var(&reqNum, "reqNum", 1, "Number of requests to perform")
	flag.IntVar(&reqNumPerConn, "reqNumPerConn", 1, "req num per connect")
	flag.StringVar(&hostPort, "hostPort", "127.0.0.1:2007", "echo tcp server ip:port")
	flag.IntVar(&timeoutMs, "timeoutMs", 500, "The timeout MS")
	flag.IntVar(&duration, "d", 1, "The duration of dump stat info")
	flag.IntVar(&concurrence, "c", 1, "concurrence")
	flag.IntVar(&dataLen, "dataLen", 4096, "request datalen")
	flag.BoolVar(&disPlay, "disPlay", false, "display req and rsp")
	flag.StringVar(&inputData, "inputData", "", "input data send to server")

	flag.Parse()
}

var hasSendNum int64

func main() {
	initFlag()
	hasSendNum = 0

	runtime.GOMAXPROCS(24)

	timerStatHelper := stat.NewStatHelper()
	timerStatHelper.SetTimerDump(time.Duration(duration)*time.Second, func() {
		b := bytes.NewBuffer([]byte{})
		timerStatHelper.DumpCount(b)
		//timerStatHelper.DumpTimeCost(b)
		fmt.Printf("%v\n", string(b.Bytes()))
	})

	var wg sync.WaitGroup
	for c := 0; c < concurrence; c++ {
		go tcpTest(&wg, timerStatHelper)
		wg.Add(1)
	}
	wg.Wait()
	os.Exit(0)

	//for atomic.LoadInt64(&hasSendNum) < reqNum {
	//	time.Sleep(time.Second)
	//}
	//time.Sleep(time.Second)
}
