package main

import (
	"github.com/elitecodegroovy/go-core/chart4/oop/encapsulation"
)

func main(){
	var helper *encapsulation.Helper
	helper.DoTask(2)
	helper.DoTask(2.19)
	helper.DoTask(true)
	helper.DoTask(map[string]string{"1":"1000"})
	helper.DoTask([]int{1, 2, 3, 4, 5})

}
