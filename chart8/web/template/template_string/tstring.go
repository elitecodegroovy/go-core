package main

import (
	"net/http"
	"html/template"
	"log"
)
//定义一个命名为tmpl模板
var t = template.New("tmpl")
var tmpl = `<html>
<head>
    <title>Simple Template Program</title>
</head>
<body>
<h1>渲染一个简单页面</h1>
{{ . }}
</body>
</html>
`
func init(){
	//步骤1
	var err error
	if t, err = t.Parse(tmpl); err != nil {
		log.Fatalf("启动初始化加载失败，error :%s" , err.Error())
	}
	log.Output(2, "启动初始化成功")
}
func simpleTemplateHandler(w http.ResponseWriter, r *http.Request) {
	//步骤2
	t.Execute(w, "Simple Template Program!")
}

func main() {
	server := http.Server{
		Addr: ":9090",
	}
	http.HandleFunc("/template", simpleTemplateHandler)
	log.Output(2, "程序已经启动")
	server.ListenAndServe()
}
