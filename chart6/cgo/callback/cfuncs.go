
package main

/*
#include <stdio.h>

// The  gateway function
int call_c_func_in_go(int in)
{
	printf("Go源码中的C程序，输入参数 %d\n", in);
	int callGoFunc(int);
	return callGoFunc(in);
}
*/
import "C"
