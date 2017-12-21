package main

import "fmt"

func main() {
	doArray()
}

func doArray() {
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
	fmt.Println("\n", b[2])  // "0"
	fmt.Printf("a[5]:%v", a) //输出数组a

	c := [...]int{0, 1, 2, 3}
	fmt.Printf("%T\n", c) // "[4]int"

	d1 := [2]int{1, 2}
	d2 := [...]int{1, 2}
	d3 := [2]int{1, 3}
	fmt.Println(d1 == d2, d1 == d3, d2 == d3) // "true false false"
	d := [3]int{1, 2}
	//fmt.Println(a == d) //编译出错，无法比较[2]int 和 [3]int
	fmt.Println("d:", d)

	d5 := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := d5[2:]
	fmt.Println("before d5:", d5, ", slice s1:", s1) //数组和切片共享｛2, 3, 4, 5, 6, 7｝
	d5[2] = 10
	fmt.Println("after d5:", d5, ", slice s1:", s1) //修改数组的值，切片的元素也将会改变，因为它们共享的元素地址相同

	a1 := [...]int{10, 9, 8, 7, 6, 5}
	reverseInt(a1[:])
	fmt.Println(a1) // "[5 6 7 8 9 10]"

	s2 := []int{10, 9, 8, 7, 6, 5}
	reverseInt(s2)
	fmt.Println(s2)                                 // ""
	fmt.Println("len:", len(s2), ", cap:", cap(s2)) // ""
	fmt.Printf("%T", s2)

	var x1 []int
	x1 = append(x1, 11)
	x1 = append(x1, 12, 13)
	x1 = append(x1, 14, 15, 16)
	x1 = append(x1, x1...) // append the slice x1
	fmt.Println(x1)        // "[]int[11 12 13 14 15 16 11 12 13 14 15 16]"

	langs := []string{"C", "", "C++", "Golang"}
	fmt.Printf("%q\n", rmEmpty(langs)) // `["C" "C++" "Golang"]`
	fmt.Printf("%q\n", langs)          // `["C" "C++" "Golang" "Golang"]

	stack()
	Doslice()
}

//inverse int array
func reverseInt(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func rmEmpty(str []string) []string {
	i := 0
	for _, s := range str {
		if s != "" {
			str[i] = s
			i++
		}
	}
	return str[:i]
}

func rmEmpty2(str []string) []string {
	result := str[:0]
	for _, s := range str {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

func stack() {
	stack := []string{}
	s1 := "BENZ"
	//push s1
	stack = append(stack, s1)
	fmt.Printf("stack: %v \n", stack)
	s2 := "tesla model 3"
	stack = append(stack, s2)
	//get stack the last element
	x := stack[len(stack)-1]
	fmt.Println("x :", x)

	// 去除最后一个元素后的stack
	stack = stack[:len(stack)-1]
	fmt.Printf("stack[len-1] : %v", stack)
}

func Doslice() {
	fmt.Println()
	s := []string{"one", "two", "three", "four"}
	fmt.Println(rmStr(s, 2)) //删除第二个元素后的切片： [one two four]
}

func rmStr(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
