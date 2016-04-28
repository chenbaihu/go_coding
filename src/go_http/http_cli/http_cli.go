package httpclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func HttpPost(url string, data string, connTimeoutMs int, serveTimeoutMs int) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Duration(connTimeoutMs)*time.Millisecond)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(time.Duration(serveTimeoutMs) * time.Millisecond))
				return c, nil
			},
		},
	}

	body := strings.NewReader(data)
	reqest, _ := http.NewRequest("POST", url, body)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//reqest.Header.Set("Content-Type", "application/octet-stream")
	reqest.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)")
	//reqest.Header.Set("Connection", "close")
	response, err := client.Do(reqest)
	if err != nil {
		err = errors.New(fmt.Sprintf("http failed, POST url:%s, reason:%s", url, err.Error()))
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http status code error, POST url:%s, code:%d", url, response.StatusCode))
		return nil, err
	}

	res_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("cannot read http response, POST url:%s, reason:%s", url, err.Error()))
		return nil, err
	}
	return res_body, nil
}

func HttpGet(url string, connTimeoutMs int, serveTimeoutMs int) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Duration(connTimeoutMs)*time.Millisecond)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(time.Duration(serveTimeoutMs) * time.Millisecond))
				return c, nil
			},
		},
	}

	reqest, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(reqest)
	if err != nil {
		err = errors.New(fmt.Sprintf("http failed, GET url:%s, reason:%s", url, err.Error()))
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http status code error, GET url:%s, code:%d", url, response.StatusCode))
		return nil, err
	}

	res_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("cannot read http response, GET url:%s, reason:%s", url, err.Error()))
		return nil, err
	}
	return res_body, nil
}

func timeoutClient(timeout time.Duration) *http.Client {
	//http://www.tuicool.com/articles/rmaYBz
	return &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(timeout)
				c, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
}

// RemoteAddr 从HTTP头信息中取得客户端IP地址  X-Real-IP
func RemoteAddr(r *http.Request) string {
	if r == nil {
		return "0.0.0.0"
	}

	addrHeaderKeys := []string{
		"X-Real-IP",
		"HTTP_CLIENT_IP",
		"HTTP_X_FORWARDED_FOR",
		"HTTP_X_FORWARDED",
		"HTTP_X_CLUSTER_CLIENT_IP",
		"HTTP_FORWARDED_FOR",
		"REMOTE_ADDR",
	}
	for _, header := range addrHeaderKeys {
		if addr := r.Header.Get(header); addr != "" {
			return addr
		}
	}

	// The HTTP server in this package
	// sets RemoteAddr to an "IP:port" address before invoking a
	// handler.
	return strings.SplitN(r.RemoteAddr, ":", 2)[0]
}
