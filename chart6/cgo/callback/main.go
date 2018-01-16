
package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lclibrary

#include "clibrary.h"

int call_c_func_in_go(int in); // Forward declaration.
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export callGoFunc
func callGoFunc(in int) int {
	fmt.Printf("callGoFunc调起:输入参数 %d\n", in)
	return in + 1
}

func main() {
	fmt.Printf("Go main()函数 \n")
	C.do_c_func((C.callback_func)(unsafe.Pointer(C.call_c_func_in_go)))
}
