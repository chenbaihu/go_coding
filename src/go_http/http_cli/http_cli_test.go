package httpclient

import (
	"fmt"
	"net/http"
	"testing"
	_ "time"
)

func TestNotFound(t *testing.T) {
	//_, err := HttpPost("http://so.com/404", []byte("golang"), 100, 200)
	_, err := HttpPost("http://so.com/404", "golang", 100, 200)
	if err == nil {
		t.Fail()
	}
}

func TestTimeout(t *testing.T) {
	//_, err := HttpPost("https://twitter.com", []byte("golang"), 100, 200)
	_, err := HttpPost("https://twitter.com", "golang", 100, 200)
	if err == nil {
		t.Fatalf("兲朝可以访问twitter.com???????????")
	}
}

func TestNormal(t *testing.T) {
	//_, err := HttpPost("http://www.haosou.com", []byte("q=golang"), 100, 200)
	_, err := HttpPost("http://www.haosou.com", "q=golang", 100, 200)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoteAddr(t *testing.T) {
	ipMap := map[string]string{
		"X-Real-IP":                "1.1.1.1",
		"HTTP_CliENT_IP":           "2.2.2.2",
		"HttP_X_FORWARDED_FOR":     "3.3.3.3",
		"HTTP_X_ForWARDED":         "4.4.4.4",
		"HTTP_X_CLUstER_CLIENT_IP": "5.5.5.5",
		"HTTP_FORWARdeD_FOR":       "6.6.6.6",
		"REMOTE_ADdR":              "7.7.7.7",
	}

	r, _ := http.NewRequest("GET", "http://example.com", nil)
	for k, v := range ipMap {
		r.Header.Add(k, v)
		if RemoteAddr(r) != v {
			t.Fatalf("%v != %v", k, v)
		}
		fmt.Printf("RemoteAddr:r=%v\tip=%s\n\n", r, RemoteAddr(r))
		r.Header.Del(k)
	}
}
