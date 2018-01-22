package main

import "fmt"

type base struct {
	a 	string
	b 	int
}

type derivedBase struct {
	base					//嵌套
	a 	float32				//与基机构重名，覆盖基结构中的
	c 	int64
}

func main(){
	var x derivedBase
	fmt.Printf("derivedBase a type :%T\n", x.a)      //a type :float32
	fmt.Printf("b type: %T, c type: %T\n", x.b, x.c) //b type: int, c type: int64
	fmt.Printf("base c type: %T", x.base.a)		      //base c type: string
	fmt.Printf("base b type: %T", x.base.b)		      //base b type: int
}





