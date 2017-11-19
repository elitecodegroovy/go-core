package main

import (
	"fmt"
	"unicode"
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

func TestNew(){
	var new, old int
	new = 10
	fmt.Println("new - old = ", new - old)
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
