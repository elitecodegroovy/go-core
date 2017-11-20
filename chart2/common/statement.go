package main

import (
	"fmt"
	"unicode"
	"strconv"
)

const globalN = 100

func fibonacci( v int) int{
	if v == 0 || v == 1 {
		return v
	}

	return fibonacci(v - 2) + fibonacci(v -1)
}

//Upper the string's first letter
func UpperFirstLetter(v ...interface{}) string{
	s, ok := v[0].(string)
	if ok {
		for i , v := range s {
			return string(unicode.ToUpper(v)) + s[i+1:]
		}
	}
	return ""
}

func increase(v *int) {
	*v++
}


func doWithPointer(){
	d := 10
	p := &d
	fmt.Println( "\n" , *p == 10)  				//true
	*p = 0
	fmt.Println( d == 0)              				//true
	fmt.Println( p != nil )           				//true

	var x, y string
	fmt.Println( &x == &x, &x == &y, &x == nil) //true false false

	a := 0
	increase(&a)
	fmt.Println(a)
	//increase(a)     //传递的参数类型错误，目标函数要求是int类型指针
}

func TestNew() {
	var new, old int
	new = 10
	fmt.Println("new - old = ", new-old)

	Test1()
	Test2()
}
func Test1(){
	brands := []string{"奔驰", "宝马", "奥迪"}
	for x := 0; x < len(brands); x++ {
		fmt.Print(brands[x]) //输出：奔驰宝马奥迪
	}
	//fmt.Println("i: ", x) // 函数内部，x的值为1
}

var x = 2
func Test2() {
	x := 1
	brands := []string{"奔驰", "宝马", "奥迪"}
	fmt.Println("")
	for x := 0 ; x < len(brands); x++ {
		fmt.Print(brands[x])//输出：宝马宝马宝马
	}
	fmt.Println("i: ", x)   // 函数内部，x的值为1
}

func fibonacciChan(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
	fmt.Println("fibonacci x :", x)
}

func doWithChannel(){
	c := make(chan int)
	quit := make(chan int)
	//produce data
	go func(){
		for i:= 0; i < 10; i++ {
			fmt.Println("channel data item ", <- c)
		}
		quit <- 0
	}()
	fibonacciChan(c, quit)
}

func main(){
	const localM = 10
	const i, j = 10, 12
	fmt.Printf("globaleN value %d \n", globalN)
	fmt.Printf("localM value %d \n" , localM)

	fmt.Printf("fibonacci %d: %d \n", i,  fibonacci(i))
	fmt.Printf("fibonacci %d: %d \n", j, fibonacci(j))

	//declare variables
	var s string
	var x, y, z int
	var m, n, f = true, 2.9, "string"
	fmt.Print(s, x,y,z, m, n, f)



	doWithPointer()

	TestNew()
}
