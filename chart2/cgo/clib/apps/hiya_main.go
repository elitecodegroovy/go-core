package main

import "github.com/elitecodegroovy/go-core/chart2/cgo/clib/hiya"


func main() {
	msg := hiya.CreateMessage("Welcome to cgo world!")
	hiya.DisplayMessage(msg)
	hiya.FreeMessage(msg)
}
