package main

import (
	"html/template"
	"os"
	"fmt"
)

var action2 =
	`
	{{ define "templateA" }} 
	模板A包含模板B
	{{ template "templateB" .}}
	模板A结束
	{{ end }}
	{{define "templateB"}} 
	执行模板
	{{end}}
	`
func main() {
	var err error

	t := template.New("action2")
	t, err = t.Parse(action2)

	if err != nil {
		fmt.Printf("解析失败: %s", err)
	}
	err = t.ExecuteTemplate(os.Stdout, "templateA", nil)
	if err != nil {
		fmt.Printf("执行失败: %s", err)
	}
}