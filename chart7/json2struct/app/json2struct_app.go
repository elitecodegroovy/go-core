// 从JSON文档映射为对应的结构
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	. "github.com/elitecodegroovy/go-core/chart7/json2struct"
)

var (
	name        = flag.String("name", "oo", "结构的名称")
	pkg         = flag.String("pkg", "main", "指定代码的包名称")
	fileName   = flag.String("filename", "", "包含json数据的文件名称（默认情况是STDIN）")
	outputFilename  = flag.String("o", "", "输出文件名称 (默认情况是STDOUT)")
	format      = flag.String("fmt", "json", "格式化输入数据(json 或者 yaml，默认是json)")
	tags        = flag.String("tags", "fmt", "逗号分离的一系列输入到结构上的标记，默认情况与fmt相同")
	forceFloats = flag.Bool("forcefloats", false, " 强制 float64 类型 为integral值（默认情况为false）")
	subStruct   = flag.Bool("subStruct", false, "创建子结构 (默认情况是false)")
)

func main() {
	flag.Parse()

	tagList := make([]string, 0)
	if tags == nil || *tags == "" || *tags == "fmt" {
		tagList = append(tagList, *format)
	} else {
		tagList = strings.Split(*tags, ",")
	}

	if isInteractive() && *fileName == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "Expects input on stdin")
		os.Exit(1)
	}

	var reader io.Reader
	reader = os.Stdin
	if *fileName != "" {
		f, err := os.Open(*fileName)
		if err != nil {
			log.Fatalf("reading input file: %s", err)
		}
		defer f.Close()
		reader = f
	}

	var convertFloats bool
	var parser Parser
	switch *format {
	case "json":
		parser = ParseJson
		convertFloats = true
	case "yaml":
		parser = ParseYaml
	}

	//生产结构定义
	if output, err := GenerateJSON(reader, parser, *name, *pkg, tagList, *subStruct, convertFloats); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	} else {
		//将结构定义写入文件
		if *outputFilename != "" {
			err := ioutil.WriteFile(*outputFilename, output, 0644)
			if err != nil {
				log.Fatalf("writing output: %s", err)
			}
		} else {
			fmt.Print(string(output))
		}
	}
}

//  如果os.Stdin 是交互式的，返回true
func isInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
}
