package main

import (
	"errors"
	"log"
)

func builtin() {
	//1 builtin func make
	//builtin make function created byte type slice.with default slice length 10
	slice1 := make([]byte, 10)
	log.Printf("slice1: %v, len: %d, cap: %d", slice1, len(slice1), cap(slice1))
	//builtin make function created byte type slice.with initialized length 0 and capacity 20
	slice2 := make([]byte, 0, 20)
	log.Printf("slice2: %v, len: %d, cap: %d", slice2, len(slice2), cap(slice2))

	map1 := make(map[string]int)
	log.Printf("map1: %v, len: %d", map1, len(map1))
	map2 := make(map[string]int, 10)
	//it doesn't support the calling of cap(map2)
	log.Printf("map2: %v, len: %d", map2, len(map2))

	unbuffered := make(chan int)
	log.Printf("unbuffered channel: %v, type: %T, len: %d, cap: %d",
		unbuffered, unbuffered, len(unbuffered), cap(unbuffered))
	buffered := make(chan int, 10)
	log.Printf("buffered channel: %v, type: %T, len: %d, cap: %d",
		buffered, buffered, len(buffered), cap(buffered))

	//2 builtin func new
	num := new(int)
	log.Printf("num type: %T, num: %v, *num: %v", num, num, *num)

	others()
}

func others() {
	//3 builtin func len
	ns := []int{1, 2, 3, 4, 5, 6}
	bs := make([]byte, 10, 100)
	str := " ཀ ཁ ག ངཅ ཆ ཇ ཉ "
	m := make(map[string]int)
	m["hello"] = 1
	channel := make(chan int, 5)
	channel <- 100
	channel <- 600
	var pointer *[5]byte

	log.Printf("int slice len: %d", len(ns))
	log.Printf("byte slice len: %d", len(bs))
	log.Printf("string len: %d", len(str))
	log.Printf("map len: %d", len(m))
	log.Printf("channel len: %d", len(channel))
	log.Printf("pointer len: %d", len(pointer))

	//4 builtin func cap. It didn't work for string or map type
	log.Printf("slice cap: %d", cap(bs))
	log.Printf("channel cap: %d", cap(channel))
	log.Printf("pointer cap: %d", cap(pointer))

	//5  builtin func delete
	m["one"] = 10
	m["two"] = 20
	m["three"] = 30
	log.Printf("original map: %+v", m)
	delete(m, "two")
	log.Printf("map: %+v", m)

	//6 builtin func close
	log.Printf("%d", <-channel) //100
	channel <- 1000
	close(channel)
	log.Printf("%d", <-channel) //600
	log.Printf("%d", <-channel) //1000
	if i, ok := <-channel; ok {
		log.Printf("index %d, %t: , channel is still open.", i, ok)
	} else {
		log.Printf("index %d, %t: , channel is closed..", i, ok)
	}

	//7 builtin func append
	bs2 := append(bs, '1')
	bs2 = append(bs2, '2', '3')
	log.Printf("byte slice append one or more elements :%v", bs2)
	bs1 := []byte{'4', '5', '6', '7', '8', '9', '0'}
	bs2 = append(bs2, bs1...)
	log.Printf("byte slice append []byte :%v", bs2)

	//8 builtin func copy
	ascii := "ASCII"
	copy(bs2, ascii)
	log.Printf("copy new string to slice []byte :%v", bs2)

	//9 10 11 builtin func complex/real/imag
	c1 := 9.9 + 9.8i
	c2 := complex(9.9, 9.8)
	c3 := complex(float32(1.5), float32(0.5))
	log.Printf("c3 type: %T", c3)
	log.Printf("c1 == c2: %v", c1 == c2)
	log.Printf("c1 real: %v", real(c1))
	log.Printf("c1 imag: %v", imag(c1))
	log.Printf("c1 + c2: %v", c1+c2)
	log.Printf("c1 - c2: %v", c1-c2)
	log.Printf("c1 * c2: %v", c1*c2)
	log.Printf("c1 / c2: %v", c1/c2)

	//12 13 builtin func panic/recover
	var i1, i2 = 100, 0
	//do a internal panic
	doRecover(func() {
		log.Printf("i1/i2= %d will cause the panic ", i1/i2)
	})
	//call panic func
	doRecover(func() {
		panic(errors.New("an error occured"))
	})

	//14 15 builtin func print/println
	print("Print is useful for bootstrapping and debugging;",
		" it is not guaranteed to stay in the language.")
	println("Spaces are always added between arguments and a newline is appended.",
		"Println is useful for bootstrapping and debugging;",
		" it is not guaranteed to stay in the language.")
}

func doRecover(f func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("got an error: %v", r)
			return
		}
	}()
	f()
}

func main() {
	builtin()
}
