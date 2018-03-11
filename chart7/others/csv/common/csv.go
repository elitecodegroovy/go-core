package main

import (
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"io"
	"encoding/json"
)


type User struct {
	Username  string   `json:"username"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	State string `json:"state"`
	City  string `json:"city"`
	Detail string `json:"detail"`
}

//读取csv文件并转化为json格式输出
func HandleCSVFile(){
	csvFile, _ := os.Open( "user.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var peoples []User
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("read error " + err.Error())
		}
		peoples = append(peoples, User{
			Username: line[0],
			Address: &Address{
				State: line[1],
				City:  line[2],
				Detail:line[3],
			},
		})
	}
	peopleJson, _ := json.Marshal(peoples)
	log.Println(string(peopleJson))
}
func main(){
	HandleCSVFile()
}
