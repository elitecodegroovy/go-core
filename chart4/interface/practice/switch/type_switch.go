package main

import (
	"fmt"
	"reflect"
	"os"
)

func printTypeInfo(x interface{}) {
	switch t := x.(type) {
	case nil :
		fmt.Printf(" %#v is a nil.\n", t)
	case bool:
		fmt.Printf(" %#v is a boolean.\n", t)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		fmt.Printf(" %#v is an integer.\n", t)
	case float32, float64, complex64, complex128:
		fmt.Printf(" %#v is a floating-point.\n", t)
	case string:
		fmt.Printf(" %#v is a string.\n", t)
	case []int:
		fmt.Printf(" %#v is a slice int.\n", t)
	case map[string]string:
		fmt.Printf(" %#v is a map[string]string.\n", t)
	case chan int, chan int8, chan int16, chan int32, chan int64,
			chan uint, chan uint8, chan uint16, chan uint32, chan uint64:
		fmt.Printf(" %#v is a channel int.\n", t)
	case chan string:
		fmt.Printf("%#v is a channel string.\n", t)
	case chan float32, chan float64, chan complex64, chan complex128:
		fmt.Printf("%#v is a channel float.\n", t)
	case []interface{}:
		fmt.Printf("%#v is a interface{}.\n", t)
	default:
		switch reflect.TypeOf(x).Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan, reflect.Func, reflect.Ptr:
			fmt.Printf(" %#v is a slice, array, map, or channel.\n", t)
		default:
			fmt.Printf("[%#v]Type handler was not implemented.\n", t)
		}
	}
}


func main(){
	var b bool
	printTypeInfo(b)
	printTypeInfo(1024)
	printTypeInfo(3.1415926)
	printTypeInfo("creating-a-microservice-with-golang-and-goa")
	printTypeInfo([]int{1, 2, 3, 4, 5, 6})

	services := make(map[string]string)
	services["1"] = "Authentication Microservice"
	printTypeInfo(services)

	strChan := make(chan string,  2)
	defer close(strChan)
	strChan <- "serverless-golang-api-with-aws-lambda"
	printTypeInfo(strChan)
	<- strChan

	var c interface{}
	c = os.Stdout
	printTypeInfo(c)

	var f interface{}
	printTypeInfo(f)

	printTypeInfo(map[int]int {1:10, 2:20})
}
