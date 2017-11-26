package main

import "fmt"

func main(){
	var x [5]int             // 定义一个包含5个元素的数值
	fmt.Println(x[0])        // 输出数组x[0]的值
	fmt.Println(x[len(x)-1]) // 输出数组最后一个元素的值

	// 遍历数值中的各个元素值，包括数组的索引号
	for i, v := range x {
		fmt.Printf("x[%d]:%d\t", i, v)
	}
	fmt.Println("\n --------------------")
	// 仅遍历数组的各个元素值，忽略数值的索引号
	for _, v := range x {
		fmt.Printf("%d\t", v)
	}

	var a [5]int = [5]int{0, 1, 2, 3, 4}
	var b [3]int = [3]int{1, 2}
	fmt.Println("\n", b[2])       // "0"
	fmt.Printf("a[5]:%v", a) //输出数组a

	c := [...]int{0, 1, 2, 3}
	fmt.Printf("%T\n", c) // "[4]int"

	d1 := [2]int{1, 2}
	d2 := [...]int{1, 2}
	d3 := [2]int{1, 3}
	fmt.Println(d1 == d2, d1 == d3, d2 == d3) // "true false false"
	d := [3]int{1, 2}
	//fmt.Println(a == d) //编译出错，无法比较[2]int 和 [3]int
}
