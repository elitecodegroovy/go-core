package main

import (
	"fmt"
	"regexp"
)

//Go RegExp Quick Reference
//	[abc]	A single character of: a, b or c
//	[^abc]	Any single character except: a, b, or c
//	[a-z]	Any single character in the range a-z
//	[a-zA-Z]	Any single character in the range a-z or A-Z
//	^	Start of line
//	$	End of line
//	\A	Start of string
//	\z	End of string
//	.	Any single character
//	\s	Any whitespace character
//	\S	Any non-whitespace character
//	\d	Any digit
//	\D	Any non-digit
//	\w	Any word character (letter, number, underscore)
//	\W	Any non-word character
//	\b	Any word boundary
//	(...)	Capture everything enclosed
//	(a|b)	a or b
//	a?	Zero or one of a
//	a*	Zero or more of a
//	a+	One or more of a
//	a{3}	Exactly 3 of a
//	a{3,}	3 or more of a
//	a{3,6}	Between 3 and 6 of a

func matchStr() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the dot.
	var validID = regexp.MustCompile(`^[a-zA-Z0-9_]+\.?[a-zA-Z0-9_]+@[a-z0-9]+\.[a-z0-9]{1,3}`)

	fmt.Println(validID.MatchString("adam.com@163.com"))
	fmt.Println(validID.MatchString("eve@126.cn"))
	fmt.Println(validID.MatchString("eve_jigang@126.cn"))
	fmt.Println(validID.MatchString("eve-jigang@126.cn")) //false
	fmt.Println(validID.MatchString("eve[90]@126.cn"))    //false

	//another way
	if match, err := regexp.MatchString("^[a-zA-Z0-9_]+\\.?[a-zA-Z0-9_]+@[a-z0-9]+\\.[a-z0-9]{1,3}",
		"eve90@123.cn"); err != nil {
		fmt.Errorf("err %s", err.Error())
	} else {
		fmt.Println("match :", match)
	}
}

func FindStrIndex() {
	re := regexp.MustCompile("ble?")
	fmt.Println(re.FindStringIndex("tabblett"))
	fmt.Println(re.FindStringIndex("tablett"))
}

// FindStringSubmatch returns a slice of strings holding the text of the
// leftmost match of the regular expression in s and the matches, if any, of
// its subexpressions, as defined by the 'Submatch' description in the
// package comment.
// A return value of nil indicates no match.
func FindStrSubmatch() {
	re := regexp.MustCompile("a(x*)byc")
	fmt.Printf("%q\n", re.FindStringSubmatch("-----axxxxbyc-----axbyc--------ooooooo-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-ax-"))
}

func FindAllMatch() {
	re := regexp.MustCompile("or?")
	fmt.Println(re.FindAllString("originalOrage", -1))
	fmt.Println(re.FindAllString("pollpoor", -1))
	fmt.Println(re.FindAllString("boolboodoorauthor", -1))
}

func FindAllSubmatch() {
	re := regexp.MustCompile("(xx*)b")
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-7777xxb999xxxxxb999-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-9axxb-", -1))
}

// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
func DoRegexSplit() {
	a := regexp.MustCompile(",")
	fmt.Println(a.Split("ban,an,a", -1))
	fmt.Println(a.Split("a,b,v,h,l", 2))
}
func main() {
	//匹配case，返回是否匹配成功
	matchStr()

	//查询匹配到的索引号
	FindStrIndex()

	//查询匹配的子匹配
	FindStrSubmatch()

	//查询所有匹配
	FindAllMatch()

	//查询所有子匹配
	FindAllSubmatch()

	//Split正则表达式
	DoRegexSplit()

	//online go regexp tool
	//https://regex-golang.appspot.com/assets/html/index.html
}
