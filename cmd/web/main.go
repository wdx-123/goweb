package main

import (
	"GoWeb/internal/web-app/controller"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("项目开始喽")
	// 注册路由：
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)

	// 静态资源加载
	fs := http.FileServer(http.Dir("D:\\workspace_go\\GoWeb\\web\\static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// 开始监听
	http.ListenAndServe(":8888", nil)
}
