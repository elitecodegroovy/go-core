package main

import (
	"fmt"
)

func execInterface(){
	var a interface{} = []string{"213", "2", "100"}
	fmt.Printf("a type: %T, value :%v,", a,  a,)
}

func main(){
	execInterface()

}
