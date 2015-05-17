package main

// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import "fmt"

// http://golang.org/cmd/cgo/

func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 42
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
