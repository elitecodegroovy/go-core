package main

import (
	"fmt"
	"time"
)

func doOptsOfTime() {
	// get current timestamp
	currentTime := time.Now().Local()

	//print time
	fmt.Println(currentTime)

	//format Time, string type
	newFormat := currentTime.Format("2006-01-02 15:04:05.000")
	fmt.Println(newFormat)

	//build Time 2016-02-17 23:59:59.999, DateTime type
	myTime := time.Date(0162, time.February, 17, 23, 59, 59, 999, time.UTC)

	//output the myTime
	fmt.Println("MyTime:", myTime.Format("2006-01-02 15:04:05.000"))

	fmt.Println("milliseconds:", time.Now().UnixNano()/int64(time.Millisecond))

	//TODO Changing time layout(form)
	form := "2006-01-02 15:04:05"
	time, err := time.Parse(form, "2017-12-24 19:04:05")
	if err != nil {
		fmt.Println("parsing time error", err)
	}
	fmt.Println("time :", time)
}

func main() {
	doOptsOfTime()
}
