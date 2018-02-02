package main

import (
	"sync"
	"sync/atomic"
	"fmt"
	"time"
)

var (
	gValue 		int64 								//全局变量
)
//原子递增变量gValue的值
func addWithAtomic(wg *sync.WaitGroup, v int64) {
	time.Sleep(1000 * time.Millisecond)
	defer wg.Done()
	atomic.AddInt64(&gValue, v)
	if atomic.LoadInt64(&gValue) == 1000 {
		atomic.StoreInt64(&gValue, 0)
	}
}
//仅仅执行5秒钟
func readValue(wg *sync.WaitGroup){
	defer wg.Done()
	tick := time.Tick(1000 * time.Millisecond)
	timeout := time.After(3 * time.Second)
	for {
		select {
		case <- timeout:
			fmt.Println("Time is over!")
			return
		case <- tick :
			fmt.Printf("变量gValue: %d\n", atomic.LoadInt64(&gValue))
		}
	}
}

func callAtomic() {
	maxLoop := 2000
	var wg sync.WaitGroup
	wg.Add(4 + maxLoop)								//启动4+maxLoop个goroutine
	//读取gValue值
	go readValue(&wg)

	go addWithAtomic(&wg, 3 )
	go addWithAtomic(&wg, 6)
	go addWithAtomic(&wg, 9)

	for i:=0; i < maxLoop; i++ {					//loop add value
		go addWithAtomic(&wg, 1)
	}
	wg.Wait()                                       //等待执行完成

	fmt.Printf("修改之前全局gValue:%d\n", gValue)    //输出：全局gValue:2018
	atomic.StoreInt64(&gValue, gValue+ 2000)
	fmt.Printf("修改之后全局gValue:%d", gValue)    //输出：全局gValue:2018
}

func main(){
	callAtomic()
}


