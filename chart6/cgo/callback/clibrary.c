#include <stdio.h>

#include "clibrary.h"

void do_c_func(callback_func callback)
{
	int arg = 2;
	printf("callback 参数 = %d\n", arg);
	int response = callback(2);
	printf("do_c_func return %d\n", response);
}