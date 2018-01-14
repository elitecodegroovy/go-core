// +build !windows

//test only on linux OS
package msg

/*
#cgo CFLAGS: -I/opt/goapp/src/github.com/elitecodegroovy/go-core/chart6/cgo/clib/msg
#cgo LDFLAGS: -L/opt/goapp/src/github.com/elitecodegroovy/go-core/chart6/cgo/clib/msg -lmsg
#include "msg.h"
*/
import "C"

type Message C.Message

func CreateMessage(msg string) *C.Message {
	cMsg := C.CString(msg)
	return C.create_message(cMsg)
}

func DisplayMessage(msg *C.Message) {
	C.display_message(msg)
}

func FreeMessage(msg *C.Message) {
	C.free_message(msg)
}
