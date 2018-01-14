// +build !windows

package main

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int get_two()
// {
//	    return 2018;
// }
import "C"

//Calling C function pointers is currently not supported, however you can declare Go variables
// which hold C function pointers and pass them back and forth between Go and C.
func main() {
	f := C.intFunc(C.get_two)
	fmt.Println(int(C.bridge_int_func(f)))

	var cmsg *C.char = C.CString("Happy New 2018!")
	defer C.free(unsafe.Pointer(cmsg))

	//Turning C arrays into Go slices
	//var theCArray *C.YourType = C.getTheArray()
	//length := C.getTheArrayLength()
	//slice := (*[1 << 30]C.YourType)(unsafe.Pointer(theCArray))[:length:length]
	//fmt.Println("Turning C arrays into Go slices : ", slice)
}
