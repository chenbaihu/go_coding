package fun

//#include "fun.h"
//#include <stdlib.h>
import "C"

// http://blog.sina.com.cn/s/blog_538d55be01015h6g.html
//再次提醒大家：import "C" 一定要紧跟C语言代码注释结束的最后一行，绝对不能空出一行，也不能和其他包合并写到import小括号内。

import (
	"unsafe"
)

func MySecret() string {
	return (C.GoString(C.MySecret()))
}

func HelloWorld(str string) {
	cstr := C.CString(str)
	C.HelloWorld(cstr)
	defer C.free(unsafe.Pointer(cstr)) // 如果不释放内容，可以看到内存增长非常快
}

// Go string to C string
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
//func C.CString(string) *C.char

// C string to Go string
//func C.GoString(*C.char) string

// C string, length to Go string
//func C.GoStringN(*C.char, C.int) string

// C pointer, length to Go []byte
//func C.GoBytes(unsafe.Pointer, C.int) []byte
