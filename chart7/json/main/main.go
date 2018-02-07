package main

import (
	js "github.com/elitecodegroovy/go-core/chart7/json"
	"fmt"
	"os"
	"path/filepath"
	"github.com/kardianos/osext"
	"log"
	"flag"
	"time"
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


func writeOneStructure(){
	msg := js.Message{
		Id: 1,
		CreatedTime:time.Now().UnixNano(),
		Msg: "The State of Go1.10",

	}
	filename := filepath.Join(binaryFilePath ,"oneMsg.json")
	if b, err := js.WriteFile(msg, filename); err != nil {
		fmt.Errorf("WriteJson2File error : %s", err.Error())
	}else {
		fmt.Println("写入文件" , filename, ",结果：", b)
	}
}

func wirteMultipleStructures(){
	msgs := []js.Message{
		{
			Id: 1,
			CreatedTime:time.Now().UnixNano(),
			Msg: "The State of Go1.10",
		},
		{
			Id: 2,
			CreatedTime:time.Now().UnixNano(),
			Msg: "Go's defer statement",
		},
		{
			Id: 3,
			CreatedTime:time.Now().UnixNano(),
			Msg: "Realtime redis channels browser",
		},
	}
	filename := filepath.Join(binaryFilePath , "multipleMsgs.json")
	if b, err := js.WriteFile(msgs, filename); err != nil {
		fmt.Errorf("WriteJson2File error : %s", err.Error())
	}else {
		fmt.Println("写入文件" , filename, ",结果：", b)
	}
}

func writeJson2File(){
	writeOneStructure()
	wirteMultipleStructures()
}

func readJsonFromFile(){
	const jsonFileName = "multipleMsgs.json"
	filename := filepath.Join(binaryFilePath , jsonFileName)
	data , err := js.ReadFile(filename)
	if err != nil {
		fmt.Errorf("ReadFile error: %s", err.Error())
	}
	fmt.Printf("读取文件%s,内容: %#v", filename, data )
}

func main(){
	writeJson2File()
	readJsonFromFile()
	//loadDBConfig()
}


func loadDBConfig(){
	fmt.Println("执行文件路径：", binaryFilePath)
	//使用绝对路径，也可以使用相对路径
	filename := filepath.Join(binaryFilePath,  "db_config.json" )
	fmt.Println(filename)
	fmt.Printf("配置文件信息: %#v", js.LoadConfiguration(filename))
}
