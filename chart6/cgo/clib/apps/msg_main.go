// +build !windows

package main

import "github.com/elitecodegroovy/go-core/chart6/cgo/clib/msg"

func main() {
	resultMsg := msg.CreateMessage("Welcome to cgo world!")
	msg.DisplayMessage(resultMsg)
	msg.FreeMessage(resultMsg)
}
