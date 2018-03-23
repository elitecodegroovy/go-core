package main

import (
	"reflect"
	"fmt"
)

func main(){
	tag := reflect.StructTag(`xml:"gopher" json:"go_programmer"`)
	fmt.Println(tag.Get("json"), tag.Get("xml"))

}
