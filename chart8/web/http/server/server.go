package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
)

var msg = "Go核心技术编程与实践"
func sayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, msg)
}
func main() {
	http.HandleFunc("/go", sayHandler)

	//设置服务器属性
	server := &http.Server{
		Addr: "0.0.0.0:9090",
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
		MaxHeaderBytes: 1 << 12,
	}
	log.Println("程序已启动...")
	server.ListenAndServe()
}