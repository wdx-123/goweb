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
		SendErrorResponse(w, pkg.StatusCodeMap[400], "无效请求数据，解析失败")
		return
	}

	// 认证服务调用
	if err := AuthService.LoginDetection(loginReq); err != nil { // 错误比较，分类返回
		switch {
		case errors.Is(err, pkg.ErrUserNotFound):
			SendErrorResponse(w, pkg.StatusCodeMap[400], "用户不存在")
		case errors.Is(err, pkg.ErrPasswordMismatch):
			SendErrorResponse(w, pkg.StatusCodeMap[404], "密码错误")
		default:
			SendErrorResponse(w, pkg.StatusCodeMap[500], "服务器内部错误")
		}
		return
	}
	setSessionCookie(w, loginReq.Username, loginReq.RememberMe)

	// 返回成功响应
	SendSuccessResponse(w, loginReq.Username+"登入成功")
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/index.html")
	if err != nil {
		log.Println(err)
	}
	// 检查Cookie（验证是否登录）
	//cookie, err := r.Cookie("user_session")
	//if err != nil {
	//	// 没有Cookie：未登录，跳回登录页
	//	http.Redirect(w, r, "/static/index.html", http.StatusFound)
	//	return
	//}
	//
	//// 有Cookie：显示登录成功页面
	//w.Write([]byte(`
	//    <h1>欢迎回来，` + cookie.Value + `！</h1>
	//    <a href="/logout">退出登录</a>
	//`))
	t.Execute(w, nil)
}
func UserViewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/user.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
