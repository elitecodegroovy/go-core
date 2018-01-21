package main

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

import (
	"fmt"
	"unsafe"
)
//Calling C function pointers is currently not supported, however you can declare Go variables
// which hold C function pointers and pass them back and forth between Go and C.
func main() {
	//call C's function
	f := C.intFunc(C.get_two)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 2018

	var cmsg *C.char = C.CString("Happy New Year 2018!")
	defer C.free(unsafe.Pointer(cmsg))

	// C string to Go string
	fmt.Printf("%s\n", C.GoString(cmsg))
	//output: Happy New Year 2018!
	//C string, length to Go string
	fmt.Printf("%s\n", C.GoStringN(cmsg, 5))
	//output: Happy
}
