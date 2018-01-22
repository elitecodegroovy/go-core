package main

import "fmt"

type Line struct {
	Name 	string
}

type Point struct {
	x, y float64
}

type Multiline struct {
	Line                //嵌套
	times 	int
	isTwist bool
}

type Rectangle struct {
	Line
	Multiline
	centre		  Point   //标准组合结构类型
	width, height float64

}


func main(){
	var rectangle = Rectangle{
		Line{"直线"},
		Multiline{Line{"直线2"}, 2, true},
		Point{122, 190},
		10, 15.9,
	}

	fmt.Println("", rectangle.Name)
	fmt.Println(rectangle.Multiline)
	fmt.Println(rectangle.Multiline.Name)
}