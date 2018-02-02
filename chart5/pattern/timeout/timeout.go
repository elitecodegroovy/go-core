package main

import (
	"time"
	"fmt"
	"errors"
	"sync"
)

func DoTask1(c chan bool) {
	time.Sleep(time.Second * 1)
	c <- true
}

func DoTask2(c chan bool) {
	time.Sleep(time.Second * 2)
	c <- true
}

func DoTask3(c chan bool) {
	time.Sleep(time.Second * 3)
	c <- true
}

var err = errors.New("任务超时")

//使用time.After
func DoJob1(wg *sync.WaitGroup) (bool, error){
	defer wg.Done()
	c1 := make(chan bool, 1)
	defer close(c1)

	go DoTask1(c1)
	t1 := time.Now()
	select {
	case r := <- c1:
		fmt.Println("成功完成Job1, 耗时 ", time.Since(t1) )
		return r, nil
	case <- time.After(2* time.Second):
		fmt.Println("Job1执行任务超时")
		return false, err
	}
}

//使用time.Tick
func DoJob2(wg *sync.WaitGroup) (bool, error){
	defer wg.Done()
	c1 := make(chan bool, 1)
	defer close(c1)
	go DoTask2(c1)

	tick := time.Tick(3 * time.Second)
	t1 := time.Now()
	select {
	case r := <- c1:
		fmt.Println("成功完成Job2, 耗时 ", time.Since(t1) )
		return r, nil
	case <- tick:
		fmt.Println("Job2执行任务超时")
		return false, err
	}
}

func DoJob3(wg *sync.WaitGroup) (bool, error){
	defer wg.Done()
	c1 := make(chan bool, 1)
	defer close(c1)
	go DoTask3(c1)

	timer := time.NewTimer(4 * time.Second)
	t1 := time.Now()
	select {
	case r := <- c1:
		fmt.Println("成功完成Job3, 耗时 ", time.Since(t1) )
		return r, nil
	case <- timer.C:
		fmt.Println("Job3执行任务超时")
		return false, err
	}
}

func TestTimeOut(){
	var wg sync.WaitGroup
	wg.Add(3)
	go DoJob1(&wg)
	go DoJob2(&wg)
	go DoJob3(&wg)
	wg.Wait()
}

func main(){
	TestTimeOut()
}