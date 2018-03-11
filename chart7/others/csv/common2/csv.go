package main

import (
	"os"
	"log"
	"encoding/csv"
	"bufio"
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
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("read error " + err.Error())
	}
	log.Printf("Lines: %d", len(records))
	for i := range records {
		peoples = append(peoples, User{
			Username: records[i][0],
			Address: &Address{
				State: records[i][1],
				City:  records[i][2],
				Detail:records[i][3],
			},
		})

	}
	peopleJson, _ := json.Marshal(peoples)
	log.Println(string(peopleJson))
}
func main(){
	HandleCSVFile()
}
