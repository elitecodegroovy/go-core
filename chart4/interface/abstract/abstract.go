package main

import (
	"time"
	"strconv"
	"encoding/json"
	"fmt"
	"bytes"
	"os"
)

type Book struct {
	Id int64
	Name string
	PublishTime time.Time
}

type Person struct {
	Id int64			`json:"id"`
	Name string			`json:"name"`
	Age int				`json:"age"`
	Position string		`json:"position"`
}


func (b *Book) String() string{
	return strconv.FormatInt(b.Id, 10) + b.Name + b.PublishTime.String()
}

func (p *Person) String() string{
	if result, err := json.Marshal(p); err != nil {
		fmt.Errorf(" json.Marshal error %s", err.Error())
		return "nil"
	}else {
		return string(result)
	}
}

func DoInterface(){
	var str fmt.Stringer
	var t1 time.Time
	str = t1					    //time.Time拥有String()实现方法
	fmt.Print("\n t1:", str)
	str = new(Book)					//Book拥有String()实现方法
	str = new(bytes.Buffer)			//bytes.Buffer拥有String()实现方法

	//str = new(bytes.Reader)			//无法编译：bytes.Reader没有String()实现方法

	var g interface{}
	g = 1000
	g = 10.99
	g = "generic type"
	g = map[string]int {"one": 1, "two":2}
	g = os.Stdout
	g =  new(time.Timer)
	fmt.Println("g ", g)
}

func main(){
	p := &Person{
		1000,
		"刘继刚",
		30,
		"高级工程师",
	}
	b := &Book {
		10,
		"Go核心技术编程与实践",
		time.Now(),
	}
	fmt.Print("persion: ", p, ", book：", b)
	DoInterface()
}