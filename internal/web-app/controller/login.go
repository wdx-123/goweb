package controller

import (
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/AuthService"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
)

// LoginViewHandler 处理页面登入
func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/login.html")
	if err != nil {
		log.Println("出现错误: ", err)
	}
	t.Execute(w, nil)
}
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("申请了LoginUserHandler函数")
	// 只允许POST请求，这个已在路由层保证
	var loginReq model.LoginReq
	// 解析前端数据
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		w.WriteHeader(pkg.StatusCodeMap[400]) // 参数错误
		log.Println("无效请求数据，解析失败")
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "无效请求数据，解析失败")
		return
	}

	// 认证服务调用
	if err := AuthService.LoginDetection(loginReq); err != nil { // 错误比较，分类返回
		switch {
		case errors.Is(err, pkg.ErrUserNotFound):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "用户不存在")
		case errors.Is(err, pkg.ErrPasswordMismatch):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "密码错误")
		default:
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "服务器内部错误")
		}
		return
	}
	pkg.SetSessionCookie(w, loginReq.Username, loginReq.RememberMe)

	// 返回成功响应
	pkg.SendSuccessResponse(w, loginReq.Username+"登入成功")
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/index.html")
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, nil)
}
func UserViewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/user.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
