package main

import (
	"fmt"
	"math/cmplx"
)

func doFloat(){
	var i int32 = 1
	var j int16 = 2
	//var x int = j + i // 编译出错，类型不同
	fmt.Println("i:", i, ", j:", j)

	f1 := 2.15 // 浮点数
	i1 := int(f1)
	fmt.Println(f1, " ", i1) // "2.15 3"
	f1 = 9.99
	fmt.Println(int(f1)) // "9"

	o := 0777
	fmt.Printf("%d %[1]o %#[1]o\n", o)       // "511 777 0777"
	y := int64(0xbadbed)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", y)//12245997 badbed 0xbadbed 0XBADBED

	// 1 << 24
	var f2 float32 = 16777216
	//f2:  1.6777216e+07 , f2 + 1= 1.6777216e+07
	fmt.Println("f2: ", f2, ", f2 + 1=", f2+1)
	// "f2 == f2+1 : true"
	fmt.Println("f2 == f2+1 :", f2 == f2+1)

	var f3 float32 = 1 << 25
	fmt.Println("f3 :", f3)
	var f4 = f3 + 1
	fmt.Println("f3 == f4 :", f3 == f4)

}

func doComplex(){
	var x complex64 = complex(2, 3) // 2+3i
	var y complex64 = complex(1, 9) // 1+9i
	fmt.Println(x*y) // "(-25+21i)"
	fmt.Println(real(x*y)) // "-25"
	fmt.Println(imag(x*y)) // "21"

	fmt.Println(12i * 10i) // "(-120+0i)"

	fmt.Println(cmplx.Sqrt(-1)) //复数-1的开方根： "(0+1i)"
	fmt.Println(cmplx.Exp(2))  //e的2次方："(7.38905609893065+0i)"

}

func main(){
	doFloat()
	doComplex()
}
