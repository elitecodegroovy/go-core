package main

import (
	js "github.com/elitecodegroovy/go-core/chart7/json"
	"fmt"
	"os"
	"path/filepath"
	"github.com/kardianos/osext"
	"log"
)

var (
	// Initialization of the working directory. Needed to load asset files.
	binaryFilePath = initWorkingDirectory()
)

func initWorkingDirectory() string {
	// Get the absolute path this executable is located in.
	executablePath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Error: Couldn't determine working directory: " + err.Error())
	}
	// Set the working directory to the path the executable is located in.
	os.Chdir(executablePath)
	return ""
}
func loadDBConfig(){
	fmt.Println("执行文件路径：", binaryFilePath)
	//使用绝对路径，也可以使用相对路径
	filename := filepath.Join(binaryFilePath,  "db_config.json" )
	fmt.Println(filename)
	fmt.Printf("db_config: %#v", js.LoadConfiguration(filename))

}
func main(){
	loadDBConfig()
}

//
