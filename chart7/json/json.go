package json

import (
	"os"
	"fmt"
	"encoding/json"
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