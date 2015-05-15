package stat

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

type TimerStatus struct {
	mutex     sync.RWMutex
	tsRecodes map[string]*tsRecode

	timer *time.Timer
}

type tsRecode struct {
	count int64
}

func NewTimerStatus() (ts *TimerStatus) {
	ts = new(TimerStatus)
	ts.tsRecodes = make(map[string]*tsRecode)
	return
}

func (ts *TimerStatus) SetTimerDump(duration time.Duration, f func()) {
	ts.timer = time.AfterFunc(duration, func() {
		f()
		ts.mutex.Lock()
		ts.tsRecodes = make(map[string]*tsRecode)
		ts.mutex.Unlock()
	})
}

func (ts *TimerStatus) AddCount(name string) {
	ts.AddCountN(name, 1)
}

func (ts *TimerStatus) AddCountN(name string, count int) {
	cnt := ts.getCountRecord(name)
	atomic.AddInt64(&cnt.count, int64(count))
}

func (ts *TimerStatus) DumpCount(writer io.Writer) error {
	buffio := bufio.NewWriter(writer)
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	for name, count := range ts.tsRecodes {
		if _, err := fmt.Fprintf(buffio, "%s_c=%d\t", name, count.count); err != nil {
			return err
		}
	}
	return buffio.Flush()
}

func (ts *TimerStatus) getCountRecord(name string) *tsRecode {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	r, exists := ts.tsRecodes[name]
	if !exists {
		r = new(tsRecode)
		ts.tsRecodes[name] = r
	}
	return r
}
