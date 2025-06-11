package main

import (
	"GoWeb/internal/web-app/controller"
	"GoWeb/internal/web-app/dao"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("项目开始喽")

	// 连接数据库
	err := dao.InitDB("root:1234@tcp(localhost:3306)/goweb")
	if err != nil {
		fmt.Println("数据库打开失败：", err)
	}
	defer dao.DB.Close()

	// 注册路由：
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)

	// 静态资源加载
	fs := http.FileServer(http.Dir("D:\\workspace_go\\GoWeb\\web\\static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Printf("http://localhost:8888/login")
	// 开始监听
	http.ListenAndServe(":8888", nil)
}
