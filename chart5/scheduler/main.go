package main

import (
	"fmt"
	"time"
)

func doTask(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func doScheduler1() {
	go doTask("go schedule 1")
	doTask("schedule 1")
}

func doScheduler2() {
	doTask("schedule 2")
	go doTask("go schedule 2")
}

func main() {
	//doScheduler1()
	doScheduler2()
}
