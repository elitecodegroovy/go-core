package main

import (
	"fmt"
	"sort"
)

func doMap(){
	fmt.Println()
	fruit := map[string]int {
		"apple":10,
		"orange":2,
		"banana":100,
		"grape":300,
	}
	fruit["peach"] = 10
	fruit["litchi"] = 200
	for k, v := range fruit {
		fmt.Printf(" %s, quanity: %d\n", k, v)
	}
	fmt.Println()
	var keys []string
	for k, _ := range fruit {
		keys = append(keys, k)
	}
	//排序keys
	sort.Strings(keys)
	for i:=0 ; i < len(keys); i++ {
		fmt.Printf(" %s, quanity: %d\n", keys[i], fruit[keys[i]])
	}

	var authors map[string]string
	fmt.Println("init map :", authors == nil, ", map size 0: ", len(authors) == 0)
	//authors["John.Lau"] = "Go核心技术编程"
	fmt.Println("", authors)

}

func compareInt(x, y map[string]int) bool{
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
func main(){
	a := make(map[int]string)
	a[1] = "张三"
	a[2] = "李四"
	fmt.Printf("map a : %v", a)

	b := map[int]string {
		3:"张瑞",
		4:"李明",
	}
	fmt.Printf("\nmap b:%v", b)
	fmt.Printf("\nmap b[3]:%s", b[3]) //map b[3]:张瑞
	fmt.Printf("\nmap b[2]:%s", b[2]) //map b[3]:张瑞

	if name, ok := b[3]; ok {
		//TODO ...
		fmt.Println("name :", name)
	}
	//delete one element
	delete(b, 3)
	fmt.Printf("\nmap b:%v", b)
	doMap()
}
