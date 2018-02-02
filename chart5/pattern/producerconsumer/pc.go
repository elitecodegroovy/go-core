package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

var closeCon int32 = 0

func doProducer(c chan int, done chan bool) {
	if success := rand.Float32() > 0.5; success {
		c <- rand.Intn(1000)
		done <- true
	}else {
		done <- false
	}

}

func doConsumer(c chan int) {
	for {
		num := <-c
		fmt.Printf("生产者产生的随机数： %d\n", num)
		if atomic.LoadInt32(&closeCon) != 0 {
			close(c)
			break
		}
	}
}

func handlePC(){
	const pNum = 20
	//有缓存的管道，
	results := make(chan int, 10)
	done := make(chan bool)
	//开启生产者
	for i := 0; i < pNum; i++ {
		go doProducer(results, done)
	}
	//开启消费者
	go doConsumer(results)

	//等待所有的生产者结束
	for i := 0; i < pNum; i++ {
		fmt.Println("计算结果：",<-done)
	}
	//close consumer
	atomic.StoreInt32(&closeCon, 1)
	close(done)
	fmt.Println("程序结束...")
}

func main() {
	handlePC()
}
