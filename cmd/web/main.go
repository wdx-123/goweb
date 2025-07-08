package main

import (
	"GoWeb/internal/web-app/controller"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("项目开始！")

	// 使用gorilla/mux创建路由器
	r := mux.NewRouter()
	// ================ 页面路由 ================
	r.HandleFunc("/login", controller.LoginViewHandler).Methods("GET")       // 登入界面
	r.HandleFunc("/register", controller.RegisterViewHandler).Methods("GET") // 注册界面
	r.HandleFunc("/dashboard_login", controller.Dashboard).Methods("GET")    // 登入后的主界面
	r.HandleFunc("/user", controller.UserViewHandler).Methods("GET")         // 用户界面

	// ================ API路由 ================
	apiRouter := r.PathPrefix("/api").Subrouter()                                   // 根据前缀切割出了一个子路由
	apiRouter.HandleFunc("/users", controller.RegisterUserHandler).Methods("POST")  // 创建用户
	apiRouter.HandleFunc("/users/{id}", controller.LoginUserHandler).Methods("GET") // 获取单个用户，带实现
	// 会话资源（登录/登出）--我推测，需要依靠中间件
	apiRouter.HandleFunc("/sessions", controller.LoginUserHandler).Methods("POST") // 创建会话（登录）
	apiRouter.HandleFunc("/sessions", controller.LogoutHandler).Methods("DELETE")  // 删除（登出）

	// 静态资源加载
	fs := http.FileServer(http.Dir("D:\\workspace_go\\GoWeb\\web\\static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	// 通过第三方路由实现静态资源加载
	// 一旦，有/static/路由的到来，就会触发监听静态资源的这个路由。

	fmt.Printf("http://localhost:8000/login")

	// 开始监听
	err = http.ListenAndServe(":8000", r) // 在这里监听我新建的第三方路由，gorilla
	if err != nil {
		log.Println("监听出现错误：", err)
	}
}
