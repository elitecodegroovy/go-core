package main

import (
	"net/http"
	"html/template"
	"log"
)

var t *template.Template
func init(){
	//步骤1
	var err error
	if t, err = template.ParseFiles("templates/simple.html"); err != nil {
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
