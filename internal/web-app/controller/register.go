package controller

import (
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/AuthService"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// RegisterViewHandler 调用注册页面
func RegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/register.html")
	if err != nil {
		log.Println("出现错误: ", err)
	}
	t.Execute(w, nil)
}

// RegisterUserHandler 注册具体用户
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// 规定回消息的 格式
	w.Header().Set("Content-Type", "application/json")
	// pkg.StatusCodeMap
	// 获取 密码
	var mod model.RegisterReq
	if err := json.NewDecoder(r.Body).Decode(&mod); err != nil {
		// 具体调用
		SendErrorResponse(w, pkg.StatusCodeMap[409], "后端json解析失败")
	}
	// 调用函数 service,确定后，调用dao层，否则返回失败
	err := AuthService.Register(mod)
	if err != nil {
		log.Println(err.Error())
		SendErrorResponse(w, pkg.StatusCodeMap[501], err.Error())
	} else {
		SendSuccessResponse(w, "注册成功，本程序欢迎您")
	}
}
