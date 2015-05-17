package log

// #cgo CFLAGS:  -I/home/chenbaihu/code_test/go_coding/src/cgo/log/clog/ -I./clog
// #cgo LDFLAGS: -L/home/chenbaihu/code_test/go_coding/src/cgo/log/clog/ -L./clog -lclog
// #cgo LDFLAGS: -Wl,-rpath=/home/chenbaihu/code_test/go_coding/src/cgo/log/clog -Wl,-rpath=./clog
// #include <stdlib.h>
// #include "clog.h"
// #include "clog/clog.h"
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

type MyLog struct {
	filename string
	inited   bool
}

func NewMyLog(filename string) *MyLog {
	mylog := new(MyLog)
	mylog.filename = filename
	mylog.inited = true
	return mylog
}

func CloseMyLog(mylog *MyLog) {
	mylog.filename = ""
	mylog.inited = false
}

const (
	LOG_ALL   int = 0
	LOG_TRACE int = 0
	LOG_DEBUG int = 10000
	LOG_INFO  int = 20000
	LOG_WARN  int = 30000
	LOG_ERROR int = 40000
	LOG_FATAL int = 50000
	LOG_OFF   int = 60000
)

func (log *MyLog) LogDebug(name string, file string, line int, format string, args ...interface{}) {
	log.LogAll(name, LOG_DEBUG, file, line, format, args)
}

func (log *MyLog) LogAll(name string, level int, file string, line int, format string, args ...interface{}) {
	if !log.inited {
		fmt.Println("LogAll log not inited")
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	content := C.CString(strings.Replace(fmt.Sprintf(format, args...), "%", "%%", -1))
	//content := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(content))

	C.LogDebug(cname, C.int(level), cfile, C.int(line), content)
}
