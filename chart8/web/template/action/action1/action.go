package main

import (
	"fmt"
	"html/template"
	"os"
)

var action1 = `
{{ define "templateA" }} 模板a的内容 {{.First}} ，尺寸{{.Second}} {{end }}
{{define "templateB"}} 模板b的内容 {{end}}
`

func main() {
	var err error
	tmplData := &struct{
		First string
		Second int } {
		"模板1", 100,
	}
	t := template.New("Action1")
	t, err = t.Parse(action1)
	if err != nil {
		fmt.Printf("解析失败: %s", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "templateA", tmplData)
	if err != nil {
		fmt.Printf("执行模板A失败: %s", err)
	}

	fmt.Println()
	err = t.ExecuteTemplate(os.Stdout, "templateB", tmplData)
	if err != nil {
		fmt.Printf("执行模板B失败: %s", err)
	}
}