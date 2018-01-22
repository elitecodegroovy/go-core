package main

import "fmt"

type Base struct {
	a 	string
	b 	int
}

type DerivedBase struct {
	Base					//嵌套
	a 	float32				//与基机构重名，覆盖基结构中的
	c 	int64
}

func main(){
	var x DerivedBase
	fmt.Printf("derivedBase a type :%T\n", x.a)      //a type :float32
	fmt.Printf("b type: %T, c type: %T\n", x.b, x.c) //b type: int, c type: int64
	fmt.Printf("base c type: %T\n", x.Base.a)		      //base c type: string
	fmt.Printf("base b type: %T", x.Base.b)		      //base c type: string
}





