package main

import "fmt"

func DoChannel() {
	ch := make(chan string)
	ch <- "ch1" //发送表达式
	//....
	//另外的 goroutine
	c := <-ch //接受表达式
	fmt.Printf("received c :%s", c)
	ch <- "another ch2"

	<-ch //接受表达式，抛弃接受的数据

	ch1 := make(chan int)    // 非缓存管道
	ch2 := make(chan int, 0) // 非缓存管道
	ch3 := make(chan int, 3) // 容量为3的缓存管道
	fmt.Printf("%v %v %v", ch1, ch2, ch3)

}

func main() {
	DoChannel()
}
