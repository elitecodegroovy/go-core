package main

import (
	"fmt"
	"log"
	"strconv"
)

func doFormat() {
	i := 1000
	si := strconv.Itoa(i)
	fmt.Printf("si : %s", si)
	if i64, err := strconv.ParseInt(si, 10, 32); err != nil {
		log.Printf("error ParseInt %s", err.Error())
	} else {
		log.Printf("i64 value :%d", i64)
	}

	//convert2Str
	convert2Str()
}

func convert2Str() {
	fmt.Printf("\n formatInt %s", strconv.FormatInt(100, 2))
}

func main() {
	doFormat()
}
