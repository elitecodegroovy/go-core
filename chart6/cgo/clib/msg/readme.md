
##commands

Compile C code `msg` to a library with the following commands.

```
$ gcc -Wall -g -c msg.c -o msg.o
$ ar ruv libmsg.a msg.o
$ ranlib libmsg.a
```
 
 ## go build
 
 Build Go code to a package.
 
 ```
 go install msg
```