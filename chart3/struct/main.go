package main

import (
	"fmt"
	"time"
)

type Book struct {
	Id      int64
	Name    string   //书名
	Authors []string //作者
	ISBN    string   //ISBN标记
}

type PublishedBook1 struct {
	Id            int64
	PublishedTime time.Time //出版时间
	Name          string    //出版者的名称
	book          Book
}

type PublishedBook2 struct {
	Id            int64
	PublishedTime time.Time //出版时间
	Name          string    //出版者的名称
	Book
}

func (pub1 *PublishedBook1) String() string {
	t1 := pub1.PublishedTime.Format("2006-01-02 15:04:05.000")
	return "PublishedTime:" + t1 + ", Name: " + pub1.Name + ", book Name" + pub1.book.Name + ", " + ", book ISBN:" + pub1.book.ISBN
}

func (pub2 *PublishedBook2) String() string {
	t1 := pub2.PublishedTime.Format("2006-01-02 15:04:05.000")
	return "PublishedTime:" + t1 + ", Name:" + pub2.Name + ", book Name" + pub2.Book.Name + ", " + ", book ISBN:" + pub2.ISBN
}

func execStruct() {
	form := "2006-01-02 15:04:05"
	timeStr := "2017-12-12 19:04:05"
	t, err := time.Parse(form, timeStr)
	if err != nil {
		fmt.Println("parsing time error", err)
	}
	pBook1 := PublishedBook1{Name: "人人出版社1",
		book: Book{Name: "Go网络编程",
			Authors: []string{"John.Lau", "John"},
			ISBN:    "16334300730000"},
		PublishedTime: t,
	}
	fmt.Printf("publishedBook1: %+v \n", pBook1)

	pBook2 := PublishedBook2{100, t, "人人出版社2",
		Book{1, "Go网络编程",
			[]string{"John.Lau", "John"},
			"16334300730001"},
	}
	fmt.Printf("publishedBook2: %+v \n", pBook2)
}

func main() {
	execStruct()
}
