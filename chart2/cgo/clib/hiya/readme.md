
##commands

Compile C code hiya to a library with the following commands.

```
$ gcc -Wall -g -c hiya.c -o hiya.o
$ ar ruv libhiya.a hiya.o
$ ranlib libhiya.a
```
 
 ## go build
 
 Build Go code to a package.
 
 ```
 go install hiya
```