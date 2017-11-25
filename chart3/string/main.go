package main

import "fmt"

func main(){
	s := "Welcome golang programmer! "
	fmt.Println("字符串长度：", len(s))
	//c := s[len(s)]  //超出区间的边界
	s1 := s[0:7]
	fmt.Println("s1: ", s1)

	fmt.Println(s[:]) // "Welcome golang programmer!"
	fmt.Println(s[:7]) // "Welcome"
	fmt.Println(s[8:]) // "golang programmer!"
}