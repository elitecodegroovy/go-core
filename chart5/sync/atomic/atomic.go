package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

var (
	gValue 		int64 								//全局变量
)

func addWithAtomic(wg *sync.WaitGroup, v int64) {
	defer wg.Done()
	atomic.AddInt64(&gValue, v)
}


func callAtomic() {
	maxLoop := 2000
	var wg sync.WaitGroup
	wg.Add(3 + maxLoop)								//启动3+maxLoop个goroutine

	go addWithAtomic(&wg, 3)
	go addWithAtomic(&wg, 6)
	go addWithAtomic(&wg, 9)

	for i:=0; i < maxLoop; i++ {					//loop add value
		go addWithAtomic(&wg, 1)
	}
	wg.Wait()                                       //等待执行完成

	fmt.Printf("全局gValue:%d", gValue)    //输出：全局gValue:2018
}

func main(){
	callAtomic()
}