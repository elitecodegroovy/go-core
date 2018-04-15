package main

import (
	"fmt"
	"bytes"
	"strconv"
	"strings"
)

func doConcatenation1(){
	str := "abc "+ "cba"
	fmt.Println("str=", str)
}

func doConcatenation2(){
	var buffer bytes.Buffer
	for i:=0; i < 1000; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	fmt.Println(buffer.String())
}

func doConcatenation3() {
	var strs []string
	for i := 0; i < 1000; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	fmt.Println(strings.Join(strs, ""))
}

//Only works for Go1.10 version
func doConcatenation4(){
	var b strings.Builder
	for i := 0; i < 1000; i++ {
		fmt.Fprintf(&b, "%d", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}

func convertInt(i int) string{
	return strconv.Itoa(i)
}

func convertInt64(i int64) string {
	return fmt.Sprintf("%d", i)
}

func testConverting(){
	var i int = 100
	var j int64 = 2000000
	fmt.Println("i:", convertInt(i))
	fmt.Println("j:", convertInt64(j))
}

func doString() {
	s := "Cloud native go programming"
	b := []byte(s)
	fmt.Println("byte[]:", b)
	s = string(b)
	fmt.Println("string b:", s)

}
func main(){
	doConcatenation1()
	doConcatenation2()
	doConcatenation3()
	doConcatenation4()
	testConverting()
	doString()
}