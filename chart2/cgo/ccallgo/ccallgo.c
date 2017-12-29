#include <stdio.h>

#include "_cgo_export.h"

void CFunc() {
	printf("call Go's func from the C's func \n");
	CalledByCFunc();
}