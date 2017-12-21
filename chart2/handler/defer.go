package main

import "fmt"

func main() {
	callF()
	fmt.Println("main函数正常运行结束")
}

func callF() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	callUnkownF(0)
	fmt.Println("callF函数正常运行结束")
}

//未知的第三方库，为了防止程序停止运行，需要捕获panic
func callUnkownF(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in callUnkownF", i)
	fmt.Println("递增参数i:", i)
	callUnkownF(i + 1)
}
