package cgp

import (
	"fmt"
	"runtime/debug"
	"sync"
	"testing"
	_ "time"
)

type MGetReq struct {
	A int
	B int
}

func NewMGetReq() *MGetReq {
	var r MGetReq
	return &r
}

type MGetRsp struct {
	S int
}

func NewMGetRsp() *MGetRsp {
	var r MGetRsp
	return &r
}

type MGetMsg struct {
	Req *MGetReq
	Rsp *MGetRsp
	*sync.WaitGroup
}

func NewMGetMsg() *MGetMsg {
	var mget_msg MGetMsg
	mget_msg.Req = NewMGetReq()
	mget_msg.Rsp = NewMGetRsp()
	return &mget_msg
}

func (mget_msg *MGetMsg) Cmd() error {
	mget_msg.Rsp.S = mget_msg.Req.A + mget_msg.Req.B
	return nil
}

var Cgp *ChanGoroutinePool

func Test_cgp(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("abort, unknown error, errmsg:%v, stack:%s",
				err, string(debug.Stack()))
		}
	}()

	Cgp = NewChanGoroutinePool()
	if Cgp.Start() != nil {
		fmt.Printf("cgp.Start failed")
		return
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		fmt.Printf("--------------------i=%d test-----------------\n", i)
		wg.Add(1)
		go MultiMGetMsg(i, &wg)
	}
	wg.Wait()

	if Cgp.Stop() != nil {
		fmt.Printf("Cgp.Stop failed")
	}
}

func MultiMGetMsg(i int, wg *sync.WaitGroup) {
	mget_msg_arr := make([]*MGetMsg, 10)
	msg_wg := sync.WaitGroup{}
	for i := range mget_msg_arr {
		mget_msg_arr[i] = NewMGetMsg()
		mget_msg_arr[i].Req.A = i
		mget_msg_arr[i].Req.B = i
		mget_msg_arr[i].WaitGroup = &msg_wg
		Cgp.PushJob(mget_msg_arr[i])
	}
	msg_wg.Wait()

	for i := range mget_msg_arr {
		req := mget_msg_arr[i].Req
		rsp := mget_msg_arr[i].Rsp
		fmt.Printf("i=%d req.A=%d req.B=%d Rsp.S=%d\n", i, req.A, req.B, rsp.S)
	}
	wg.Done()
}
