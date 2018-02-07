package json

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port	 string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Errorf("LoadConfiguration :%s", err.Error())
		return config
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

type Message struct {
	Id 			int64 		`json:"id"`
	CreatedTime int64 		`json:"createdTime"`
	Msg 		string 		`json:"msg"`
}

func WriteFile(msg interface{}, filename string)(bool, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
		return false, err
	}
	if err = ioutil.WriteFile(filename, b, 0644); err != nil {
		fmt.Errorf("error: %s", err.Error())
		return false, err
	}
	return true , nil
}

func ReadFile(filename string) ([]Message, error){
	var messages []Message
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(b, &messages); err != nil {
		fmt.Errorf("json.Unmarshal error: %s", err.Error())
		return nil, err
	}
	return messages, nil
}