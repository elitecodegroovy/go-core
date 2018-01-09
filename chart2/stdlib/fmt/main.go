package main

import (
	"fmt"
	"regexp"
)

func doVerbs() {
	//1 bool verb, string verb, slice verb
	s := "990aa"
	isMatched, _ := regexp.MatchString("^[0-9]{3}", s)
	//输出字符串类型和布尔类型的格式
	fmt.Printf("match the string '%s' : %t\n", s, isMatched)
	//slice verb
	bs := []byte(s)
	fmt.Printf("slice :%v, string '%s' \t %d \t %T\n, ", bs, bs, bs, bs)

	//2 integer verb
	d := 100
	//分别输出整型、二进制和靠右宽度为9个字符的二进制
	fmt.Printf("integer : %d, binary :%b|%9b\n", d, d, d)
	//分别输出靠左宽度为9个字符的二进制、靠右宽度为9个字符的二进制并且使用0填充空白位
	fmt.Printf("%-9b|%09b\n", d, d)
	//八进制与二进制类似, + 标记强制以符号标记输出
	fmt.Printf("%o|%#o|%# 9o|%#+ 9o|%+09o|\n", d, d, d, d, -d)
	//十六进制与二进制类似
	fmt.Printf("%x|%X|%9x|%09x|%#09X|0x%09X|\n", d, d, d, d, d, d)

	//3 character verb
	//数值和字符输出。ascii 编码的值转化为字符 'rune'
	fmt.Printf("%U，GuangZhou '%c%c'\n", '\u5e7f', '\u5e7f', '\u5dde')

	//4 float verb
	f := 0.006
	fmt.Printf("|%9.3f|%.2f|%.1e", f, f, f)
}

func main() {
	doVerbs()
}
