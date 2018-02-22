package main

import (
	"net/http"
	"fmt"
	"log"
)

func sayHandler(msg string ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, msg)
	})
}
func main() {
	mux := http.NewServeMux()
	msg := "Go核心技术编程与实践"

	mux.Handle("/go", sayHandler(msg))

	log.Println("程序已启动...")
	http.ListenAndServe(":9090", mux)
}

