package main

import (
	"flag"
	"github.com/kardianos/osext"
	"log"
	"os"
	myxml "github.com/elitecodegroovy/go-core/chart7/xml"
	"fmt"
	"path/filepath"
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

var xmlFileName = "companystaffs.xml"
func write2XMLFile(){
	var company  myxml.Company

	// add two staff details
	company.Staffs = append(company.Staffs, myxml.Staff{ID: 103, FirstName: "Li", LastName: "ShiMing", UserName: "李世明"})
	company.Staffs = append(company.Staffs, myxml.Staff{ID: 108, FirstName: "Liu", LastName: "BoJue", UserName: "刘伯爵"})

	filename := filepath.Join(binaryFilePath, xmlFileName)
	if result , err := myxml.Write2XMLFile(company, filename); err != nil {
		fmt.Printf("error : %s", err.Error())
	}else {
		fmt.Printf(" 写入xml：%t", result)
	}
}

func readFromXMLFile(){
	r , err :=myxml.ReadFromXMLFile(filepath.Join(binaryFilePath, xmlFileName))
	if err != nil {
		fmt.Errorf("error : %s", err.Error())
	}else {
		fmt.Printf("成功读取xml文件数据:%#v", r)
	}
}

func main(){
	write2XMLFile()
	readFromXMLFile()
}