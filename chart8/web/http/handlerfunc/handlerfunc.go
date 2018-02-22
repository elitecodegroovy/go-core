package main

import (
	"net/http"
	"fmt"
	"log"
)

var msg = "Go核心技术编程与实践"
func sayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, msg)
}
func main() {
	mux := http.NewServeMux()

	// 将sayHandler转化 为HandlerFunc类型
	hf := http.HandlerFunc(sayHandler)
	mux.Handle("/go", hf)

	//It is the same as the following statements
	//mux.HandleFunc("/go", sayHandler)

	log.Println("程序已启动...")
	http.ListenAndServe(":9090", mux)
}