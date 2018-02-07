package main

import "fmt"

type XYZ struct {
	x string
	y int
	z float64
}

func main() {
	var v XYZ
	fmt.Printf("v 类型：%T, v 值：%#v \n", v, v)
	a := "abc"
	b := 3


	ms := &XYZ{
		x: "X Spark",
		y: 100,
		z: 901.10335,
	}
	fmt.Println(ms, a, b)
}
