package main

import (
	"GoWeb/internal/web-app/controller"
	"GoWeb/internal/web-app/dao"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("项目开始喽！")

	// 连接数据库
	err := dao.InitDB("root:1234@tcp(localhost:3306)/goweb?parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("数据库打开失败：", err)
	}
	defer dao.DB.Close()

	// 注册路由：
	http.HandleFunc("/login", controller.LoginViewHandler)
	http.HandleFunc("/register", controller.RegisterViewHandler)
	http.HandleFunc("/api/register", controller.RegisterUserHandler)
	http.HandleFunc("/api/login", controller.LoginUserHandler) // 登入
	http.HandleFunc("/dashboard_login", controller.Dashboard)
	http.HandleFunc("/user", controller.Test2)
	// 静态资源加载
	fs := http.FileServer(http.Dir("D:\\workspace_go\\GoWeb\\web\\static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Printf("http://localhost:8000/login")

	// 开始监听
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("监听出现错误：", err)
	}
}
