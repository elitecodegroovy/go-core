package main

import (
	"math"
	"fmt"
	"sort"
)

func sqrtFloat(x, y float64) float64{
	return math.Sqrt(x*x + y*y)
}

func doFun(){
	fmt.Println("3*3+4*4的平方根：", sqrtFloat(3.0, 4.0))
}

func increase(n int) int{
	n++
	return n
}

func decrease(n int) int {
	n--
	return n
}

func doFuncValue(){
	incr := increase
	decr := decrease
	fmt.Printf("incr(5): %d, 函数类型:%T \n", incr(5), incr)
	fmt.Printf("incr(5): %d, 函数类型:%T \n", incr(10), incr)
	fmt.Printf("decr(10): %d, 函数类型:%T \n",  decr(10),  decr)
}
func power(a int) func(int) int {
	return func(b int) int{
		return int(math.Pow(float64(a), float64(b)))
	}
}

func doAnonymous(){
	var pow2 = power(2)
	fmt.Printf("2 ^ 10 = %d\n", pow2(10))
	var pow3 = power(5)
	fmt.Printf("5 ^2 = %d\n", pow3(2))
}

func doAnonymous1(){
	anon := func(_name string) string {
		return "Welcome, " + _name
	}
	anotherFunc(anon)
}

func anotherFunc(f func(string) string) {
	result := f("John")
	fmt.Println(result) //"Welcome, John"
}

func doAnonymous2(){
	numbers := []int{10, 1, -5, 22, 2, 0, 1}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= 7
	})
	fmt.Println("The first number >= 7 is at index:", index)
	fmt.Println("The first number >= 7 is:", numbers[index])
}


func main(){
	doFun()
	doFuncValue()
	doAnonymous()
	doAnonymous1()
	doAnonymous2()
}
