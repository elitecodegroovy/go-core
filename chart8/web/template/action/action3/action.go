package main

import (
	"fmt"
	"html/template"
	"os"
)

var action3 = `
{{ define "templateA" }} 
模板A包含模板B
{{ block "b" .}} 模板B {{ end }}
模板A结束
{{ end }}
`

func main() {
	var err error

	t := template.New("action3")

	t, err = t.Parse(action3)

	if err != nil {
		fmt.Printf("解析失败: %s", err)
	}
	err = t.ExecuteTemplate(os.Stdout, "template", nil)
	if err != nil {
		fmt.Printf("执行失败: %s", err)
	}
}
