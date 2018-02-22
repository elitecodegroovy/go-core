package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)
type APIHandler func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, error)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := "欢迎！"
	code := 200
	response := []byte(fmt.Sprintf(`{"code":%d, "message":"%s"}`,code, data))
	handleResp(w, code, response)
}

func handleResp(w http.ResponseWriter,  code int, response []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type", "httpclient; version=1.0")
	w.WriteHeader(code)
	w.Write(response)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := fmt.Sprintf( "你好， %s！", ps.ByName("uid"))
	code := 200
	response := []byte(fmt.Sprintf(`{"code":%d, "message":"%s"}`,code, data))
	handleResp(w, code, response)
}

func LogNotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data := "404 无效请求"
		code := 404
		response := []byte(fmt.Sprintf(`{"code":%d, "message":"%s"}`, code, data))
		handleResp(w, code, response)
	})
}

func LogPanicHandler() func(w http.ResponseWriter, req *http.Request, p interface{}) {
	return func(w http.ResponseWriter, req *http.Request, p interface{}) {
		log.Output(2, fmt.Sprintf("请求panic操作：%d %s %s (%s)",
			0, req.Method, req.URL.RequestURI(), req.RemoteAddr))
	}
}

func LogMethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Output(2, fmt.Sprintf("请求拒绝：%d %s %s (%s)",
			0, req.Method, req.URL.RequestURI(), req.RemoteAddr))
		data := "403 请求拒绝"
		code := 403
		response := []byte(fmt.Sprintf(`{"code":%d, "message":"%s"}`, code, data))
		handleResp(w, code, response)
	})
}

func main() {
	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	router.NotFound = LogNotFoundHandler()
	router.PanicHandler =  LogPanicHandler()
	router.MethodNotAllowed = LogMethodNotAllowedHandler()

	router.Handle("GET", "/", Index)
	router.Handle("GET", "/user/:uid", Hello)

	log.Fatal(http.ListenAndServe(":9090", router))
}