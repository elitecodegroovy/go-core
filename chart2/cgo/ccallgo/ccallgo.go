package main

import "fmt"

/*
#include <stdio.h>
extern void CFunc();
*/
import "C"

//export CalledByCFunc
func CalledByCFunc() {
	fmt.Println("go main's func  CalledByCFunc")
}

func CCallGo() {
	fmt.Println("go main calls the C's func CFunc")
	C.CFunc()
}

//Go makes its functions available to C code through use of a special //export
func main(){
	CCallGo()
}
