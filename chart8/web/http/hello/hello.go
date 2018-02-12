package main

import (
	"net/http"
	"fmt"
	"log"
)

type HelloHandler struct {
	msg 	string
}


func (m *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	mux := http.NewServeMux()
	hh := &HelloHandler{"欢迎进入Go核心编程"}
	mux.Handle("/", hh)

	log.Println("程序启动...")
	http.ListenAndServe(":9090", mux)
}