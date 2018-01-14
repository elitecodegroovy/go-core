
##Preample
In order to use cgo on Windows, you'll also need to first install a gcc compiler (for instance,
 mingw-w64) and have gcc.exe (etc.) in your PATH environment variable before compiling with cgo will work.

##commands
Input the cmds in your wingw windows dos window:
```
$ gcc -c clibrary.c
$##建立档案文件(archive)libclibrary.a
$ ar cru libclibrary.a clibrary.o
$ go build
$ ./ccallbacks

```