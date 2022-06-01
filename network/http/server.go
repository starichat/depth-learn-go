package main

import (
	"fmt"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	//1. 定义路由处理函数
	http.HandleFunc("/hello", hellohandler)
	//2. 启动服务
	fmt.Println("start server")
	http.ListenAndServe(":9000", nil)
}
