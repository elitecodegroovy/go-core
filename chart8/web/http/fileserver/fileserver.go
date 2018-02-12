package main

import (
"net/http"
"fmt"
)
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/", fs)
	fmt.Println("启动应用程序")
	http.ListenAndServe(":9090", mux)
}