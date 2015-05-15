package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type signalHandler func(s os.Signal, arg interface{})

type signalSet struct {
	m map[os.Signal]signalHandler
}

func signalSetNew() *signalSet {
	ss := new(signalSet)
	ss.m = make(map[os.Signal]signalHandler)
	return ss
}

func (set *signalSet) register(s os.Signal, handle signalHandler) {
	if _, find := set.m[s]; !find {
		set.m[s] = handle
	}
}

func (set *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handle available for signal %v", sig)
	}
}

func sysSignalHandleDemo() {
	ss := signalSetNew()
	handler := func(s os.Signal, arg interface{}) {
		fmt.Printf("handler signal: %v\n", s)
	}

	ss.register(syscall.SIGINT, handler)
	ss.register(syscall.SIGUSR1, handler)
	ss.register(syscall.SIGUSR2, handler)

	for {
		//func Notify(c chan<- os.Signal, sig …os.Signal)
		//该函数会将进程收到的系统Signal转发给channel c。转发哪些信号由该函数的可变参数决定，如果你没有传入sig参数，那么Notify会将系统收到的所有信号转发给c。
		//如果你像下面这样调用Notify：
		//signal.Notify(c, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2) 则Go只会关注你传入的Signal类型，其他Signal将会按照默认方式处理，大多都是进程退出。
		c := make(chan os.Signal)
		var sigs []os.Signal
		for sig := range ss.m {
			sigs = append(sigs, sig)
			signal.Notify(c, sig)
		}
		sig := <-c // 将在这里阻塞等待通知

		err := ss.handle(sig, nil)
		if err != nil {
			fmt.Printf("unknown signal received: %v\n", sig)
			os.Exit(1)
		}
	}
}

func main() {
	go sysSignalHandleDemo()
	time.Sleep(time.Hour) // make the main goroutine wait!
}
