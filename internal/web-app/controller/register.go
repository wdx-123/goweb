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
		log.Println("无效请求数据")
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[409], "无效请求数据")
	}
	// 调用函数 service,确定后，调用dao层，否则返回失败
	if err := AuthService.Register(mod); err != nil {
		switch {
		case errors.Is(err, pkg.ErrUserExists):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "用户存在")
		case errors.Is(err, pkg.ErrPasswordMismatch):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "用户密码错误")
		default:
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "注册失败")
		}
		return
	}
	pkg.SendSuccessResponse(w, "注册成功，本程序欢迎您")

}
