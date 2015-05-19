package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	//"strings"
	//"encoding/base64"
	"flag"
	"goapp/lcsdispatcher/stat"
	"sync"
	"sync/atomic"
	"time"
)

func createHttpCli(timeout time.Duration) *http.Client {
	//http://www.tuicool.com/articles/rmaYBz
	var httpCli *http.Client = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					fmt.Println("dail timeout", err)
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: timeout,
		},
	}
	return httpCli
}

func HTTPPost(httpCli *http.Client, url string, data []byte) ([]byte, error) {
	r, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/octet-stream")
	r.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)")
	if !keepAlive {
		r.Header.Set("Connection", "close")
	} else {
		r.Header.Set("Connection", "keep-alive")
		r.Header.Set("keep-alive", "timeout=2000")
	}

	resp, err := httpCli.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid HTTP Code: [%v]", resp.StatusCode)
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func makeReqData() []byte {
	data := bytes.NewBufferString("")
	for i := 0; i < dataLen; i++ {
		data.WriteByte(byte(i))
	}
	return data.Bytes()
}

func httpPostTest(wg *sync.WaitGroup, timerStatHelper *stat.StatHelper) {
	defer wg.Done()
	reqData := makeReqData()
	for {
		if (reqNum != 0) && (atomic.LoadInt64(&hasSendNum) > reqNum) {
			//finish
			break
		}

		httpCli := createHttpCli(time.Duration(timeoutMs) * time.Millisecond)
		for j := 0; j < reqNumPerConn; j++ {
			atomic.AddInt64(&hasSendNum, 1)

			timerStatHelper.AddCount("httppost")
			t := time.Now()
			rspData, err := HTTPPost(httpCli, url, reqData)
			timerStatHelper.AddTimeStat("onetrip", time.Since(t))
			if err != nil {
				//fmt.Println("HttpPost faile:", err)
				timerStatHelper.AddCount("httpostfail")
				continue
			}
			if bytes.Compare(reqData, rspData) == 0 {
				timerStatHelper.AddCount("httpostsucc")
				//fmt.Printf("HttpPost succ\treqData=%d\trspData=%d\n", len(reqData), len(rspData))
				//fmt.Printf("HttpPost succ\treqData=%s\trspData=%s\n", base64.StdEncoding.EncodeToString(reqData), base64.StdEncoding.EncodeToString(rspData))
			} else {
				timerStatHelper.AddCount("httpostfail2")
			}
		}
	}
}

var reqNum int64
var reqNumPerConn int
var url string
var timeoutMs int
var duration int
var concurrence int
var keepAlive bool
var dataLen int

var hasSendNum int64

func initFlag() {
	flag.Int64Var(&reqNum, "reqNum", 0, "Number of requests to perform, 0 mean not limit")
	flag.IntVar(&reqNumPerConn, "reqNumPerConn", 1, "req num per connect")
	flag.StringVar(&url, "url", "http://testserverhost:8080/echo.php", "http server url address")
	flag.IntVar(&timeoutMs, "timeoutMs", 500, "The timeout MS")
	flag.IntVar(&duration, "d", 1, "The duration of dump stat info")
	flag.IntVar(&concurrence, "c", 1, "concurrence")
	flag.IntVar(&dataLen, "dataLen", 4096, "request datalen")
	flag.BoolVar(&keepAlive, "keepAlive", true, "http keep alive")

	flag.Parse()
}

func main() {
	initFlag()
	hasSendNum = 0

	timerStatHelper := stat.NewStatHelper()
	timerStatHelper.SetTimerDump(time.Duration(duration)*time.Second, func() {
		b := bytes.NewBuffer([]byte{})
		timerStatHelper.DumpCount(b)
		timerStatHelper.DumpTimeCost(b)
		fmt.Printf("%v\n", string(b.Bytes()))
	})

	var wg sync.WaitGroup
	for c := 0; c < concurrence; c++ {
		go httpPostTest(&wg, timerStatHelper)
		wg.Add(1)
	}
	wg.Wait()
}
