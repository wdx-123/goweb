package controller

import (
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/AuthService"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
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
	// 只允许POST请求
	if r.Method != http.MethodPost {
		w.WriteHeader(pkg.StatusCodeMap[405]) // 方法不允许
		log.Println("方法不允许")
		SendErrorResponse(w, pkg.StatusCodeMap[405], "方法不允许")
		return
	}
	var loginReq model.LoginReq

	// 解析前端数据
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		w.WriteHeader(pkg.StatusCodeMap[400]) // 参数错误
		log.Println("400：参数错误")
		SendErrorResponse(w, pkg.StatusCodeMap[400], "400：参数错误")
		return
	}

	// 调用这里的函数
	if err := AuthService.LoginDetection(loginReq); err != nil {
		w.WriteHeader(pkg.StatusCodeMap[400])
		log.Println("用户不存在或密码不匹配")
		SendErrorResponse(w, pkg.StatusCodeMap[400], "400:用户不存在或密码不匹配")
		return
	}

	// 设置cookie，登入成功由于后期身份认证
	Cookie := &http.Cookie{
		Name:     "user-session",    // cookie名称
		Value:    loginReq.Username, // 存储名称
		Path:     "/",               // 全站通用
		HttpOnly: true,
		Secure:   false,
	}

	// 勾选"记住我",设置7天
	if loginReq.RememberMe {
		Cookie.Expires = time.Now().Add(7 * 24 * time.Hour) // 7天后过期
		Cookie.MaxAge = 7 * 24 * 3600
	}
	http.SetCookie(w, Cookie)
	// 返回成功响应
	SendSuccessResponse(w, "登入成功")
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
func Test2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/user.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
