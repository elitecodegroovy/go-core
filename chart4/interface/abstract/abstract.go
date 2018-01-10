package main

import (
	"time"
	"strconv"
	"encoding/json"
	"fmt"
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
}