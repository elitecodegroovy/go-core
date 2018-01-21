package main

//windows 环境下需要gcc.exe

// #include <stdio.h>
// #include <stdlib.h>
//double* get_array(int n) {
	//double *arr;
	//arr = (double*)malloc(n*sizeof(arr));
	//for(int i=0;i < n; i++){
	//	arr[i]=i;
	//}
	//return arr;
//}
import "C"
import (
	"unsafe"
	"fmt"
)

func main(){
	//C array was converted to Go slice
	size := 10
	carr := C.get_array(C.int(size))
	defer C.free(unsafe.Pointer(carr))
	arr := (*[1 << 30]float64)(unsafe.Pointer(carr))[:size:size]
	fmt.Println(arr)
	//output: [0 1 2 3 4 5 6 7 8 9]
}