package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Welcome golang programmer! "
	fmt.Println("字符串长度：", len(s))
	//c := s[len(s)]  //超出区间的边界
	s1 := s[0:7]
	fmt.Println("s1: ", s1)

	fmt.Println(s[:])  // "Welcome golang programmer!"
	fmt.Println(s[:7]) // "Welcome"
	fmt.Println(s[8:]) // "golang programmer!"

	fmt.Println("s > s1:", s > s1)

	s2 := "prefix string"
	t := s2
	s2 += ", suffix string"
	fmt.Println("t :", t)  //t : prefix string
	fmt.Println("s2:", s2) //s2: prefix string, suffix string
	//s2[0] = 'P'             //编译出错

	s3 := "great 中国"
	fmt.Println(len(s3))                    // "12"
	fmt.Println(utf8.RuneCountInString(s3)) // "8"

	fmt.Println(string(65))     // "A"，65对应的编码字符
	fmt.Println(string(0x4eac)) // "京

	fmt.Println(string(1234567)) // "�"
}
