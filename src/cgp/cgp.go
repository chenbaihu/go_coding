package cgp

import (
	"errors"
	_ "fmt"
	_ "global"
	_ "math/rand"
	"sync/atomic"
	_ "time"
)

type ChanJob interface {
	Cmd() error
	Add(delta int)
	Done()
}

type ChanGoroutinePool struct {
	//cgprand *rand.Rand   //多协程访问需要加锁
	CP    []chan ChanJob
	CI    int64 //roundrobin
	start bool
}

func NewChanGoroutinePool() *ChanGoroutinePool {
	var cgp ChanGoroutinePool
	return &cgp
}

func (cgp *ChanGoroutinePool) Start() error {
	if cgp.start {
		return nil
	}
	cgp.CP = make([]chan ChanJob, 10) //10个channel
	for ci := range cgp.CP {
		cgp.CP[ci] = make(chan ChanJob, 100000) //每个channel缓存区是10w
		for gi := 0; gi < 100; gi++ {           //每个channel下面100个线程
			go ChanHandle(cgp.CP[ci])
		}
	}
	cgp.start = true
	return nil
}

func ChanHandle(cg chan ChanJob) {
	fmt.Printf("ChanHanle start...")
	for {
		cj, cs := <-cg
		if cs == false {
			break
		}
		cj.Cmd()
		cj.Done()
	}
	fmt.Printf("ChanHanle stop...")
}

func (cgp *ChanGoroutinePool) Stop() error {
	if !cgp.start {
		return nil
	}
	for ci := range cgp.CP {
		close(cgp.CP[ci])
	}
	cgp.start = false
	return nil
}

func (cgp *ChanGoroutinePool) PushJob(cj ChanJob) error {
	cj.Add(1)

	//ci := cgprand.Intn(len(cgp.CP)) //据说随机在多协程环境下要加锁
	ci := atomic.AddInt64(&cgp.CI, 1) % int64(len(cgp.CP))
	//fmt.Printf("-----------------------------ci=%d\n", ci)

	select {
	case cgp.CP[ci] <- cj:
	default:
		fmt.Printf("PushJob cj=%v error", cj)
		return errors.New("PushJob error")
	}
	return nil
}
