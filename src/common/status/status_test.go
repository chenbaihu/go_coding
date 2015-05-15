package stat

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	var wg sync.WaitGroup
	ts := NewTimerStatus()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < i; n++ {
				ts.AddCountN("test", n)
				ts.AddCountN("test2", n)
			}
		}()
	}
	wg.Wait()
	w := bufio.NewWriter(os.Stdout)

	ts.DumpCount(w)
	fmt.Fprintf(w, "\n")
	w.Flush()
}

func doSomething(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m += i
		for j := 0; j < 30; j++ {
			strconv.Itoa(m) // cost some CPU time
		}
	}
	return m
}
