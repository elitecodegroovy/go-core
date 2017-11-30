package main

import "fmt"

type person struct {
	name string
	age int
}

func doPointer(){
	p := person{name: "李明", age: 20}
	fmt.Printf("姓名：%s, 年龄：%d", p.name, p.age)

	pp := &p
	fmt.Printf("\n姓名：%s, 年龄：%d", pp.name, pp.age )
}

func doArithmetic(){
	b := [...]int{109, 110, 111}
	p := &b
//	p++
	fmt.Println("p:", *p)
}

func main(){
	i := 100
	p := &i
	fmt.Println("*p :", *p)   //输出指针p的值
	doPointer()
	doArithmetic()

}
