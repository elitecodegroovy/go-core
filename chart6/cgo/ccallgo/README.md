You need a gcc compiler(for instance, mingw-w64) if your code are running in windows OS.

Go into the directory ccallgo and do as the following commands :
``````
>gcc -c ccallgo.c
ccallgo.c: In function 'CFunc':
ccallgo.c:7:2: warning: implicit declaration of function 'CalledByCFunc' [-Wimplicit-function-declaration]
  CalledByCFunc();
  ^~~~~~~~~~~~~
>go build
ccallgo.c: In function 'CFunc':
ccallgo.c:7:2: warning: implicit declaration of function 'CalledByCFunc' [-Wimplicit-function-declaration]
  CalledByCFunc();
  ^~~~~~~~~~~~~
>./ccallgo
go main calls the C's func CFunc
go main's func  CalledByCFunc
call Go's func from the C's func

```