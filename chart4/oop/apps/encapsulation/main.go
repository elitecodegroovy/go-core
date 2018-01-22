package main

import (
	"github.com/elitecodegroovy/go-core/chart4/oop/encapsulation"
	"fmt"
)

func main(){
	helper := &encapsulation.Helper{
		Id: 1000,
		Name: "灵犀",
	}
	fmt.Printf("%t\n", helper.GetStatus())
	helper.SetStatus(true)
	fmt.Printf("%t\n", helper.GetStatus())

	helper.DoTask(2)
	helper.DoTask(2.19)
	helper.DoTask(true)
	helper.DoTask(map[string]string{"1":"1000"})
	helper.DoTask([]int{1, 2, 3, 4, 5})

	fmt.Println(helper)
}
