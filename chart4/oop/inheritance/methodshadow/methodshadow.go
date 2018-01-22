package main

import "fmt"

type Base struct {
	a 	string
	b 	int
}

func (base *Base)PrintInfo(){
	fmt.Printf("a: %s, b: %d\n", base.a, base.b)
}

func (base *Base) SayMsg(){
	fmt.Println("I'am a Base's method--SayMsg")
}

type DerivedBase struct {
	Base					//嵌套
	a 	float32				//与基机构重名，覆盖基结构中的
	c 	int64
}

func (base *DerivedBase) SayMsg(){
	fmt.Println("I'am a DerivedBase's method--SayMsg")
}

func main(){
	var x = &DerivedBase{Base{"基机构", 100}, 11.9, 900}
	x.PrintInfo()   //a: 基机构, b: 100
	x.SayMsg()		//I'am a DerivedBase's method--SayMsg
	x.Base.SayMsg() //I'am a Base's method--SayMsg
}