package main

import (
	"fmt"
	"time"
)

func getStrFormatTime(format string) string {
	// 获取当前的时间
	currentTime := time.Now()

	//获取一个字符串类型的时间格式
	strTime := currentTime.Format(format)
	fmt.Println("当期格式化时间： ", strTime)

	return strTime
}

func do4() {
	d := []int{5, 7, 13, 17}
	for i, v := range d {
		fmt.Printf("index %d ,value:%d \n", i, v)
	}
}

func main() {
	//函数可视
	getStrFormatTime("2006-01-02 15:04:05.000")

	//变量可视
	fmt.Println("\n 二月份的值：", time.January)

	//无法获取time包中的months变量，因为它是小写字母开头的变量，无法可视
	//fmt.Println(" months[0]", time.months[0])
	//s1 := "Months"
	//fmt.Println(s1, "swapCase: ", f1(s1))
	do4()
}
