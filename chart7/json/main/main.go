package main

import (
	"github.com/elitecodegroovy/go-core/chart7/json"
	"fmt"
	"os"
	"path"
)

func loadDBConfig(pwd string){
	filename := path.Join(pwd,  "db_config.json" )
	fmt.Println(filename)
	fmt.Printf("db_config: %#v", json.LoadConfiguration(filename))

}
func main(){
	pwd, _ := os.Getwd()
	fmt.Println("path:", pwd)

	loadDBConfig(pwd)
}

//https://www.kaihag.com/external-assets-working-directories-and-go/
