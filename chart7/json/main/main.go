package main

import (
	js "github.com/elitecodegroovy/go-core/chart7/json"
	"fmt"
	"os"
	"path/filepath"
	"github.com/kardianos/osext"
	"log"
	"flag"
)

var (
	// Initialization of the working directory. Needed to load asset files.
	binaryFilePath = initWorkingDirectory()
)

//设置工作目录
func initWorkingDirectory() string {
	var customPath string
	// Check if a custom path has been provided by the user.
	flag.StringVar(&customPath, "custom-path", "",
		"Specify a custom path to the asset files. This needs to be an absolute path.")
	flag.Parse()
	// Get the absolute path this executable is located in.
	executablePath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Error: Couldn't determine working directory: " + err.Error())
	}
	// Set the working directory to the path the executable is located in.
	os.Chdir(executablePath)
	// Return the user-specified path. Empty string if no path was provided.
	return customPath
}
func loadDBConfig(){
	fmt.Println("执行文件路径：", binaryFilePath)
	//使用绝对路径，也可以使用相对路径
	filename := filepath.Join(binaryFilePath,  "db_config.json" )
	fmt.Println(filename)
	fmt.Printf("配置文件信息: %#v", js.LoadConfiguration(filename))

}
func main(){
	loadDBConfig()
}