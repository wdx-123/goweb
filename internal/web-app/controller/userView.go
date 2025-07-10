package controller

import (
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/userOperate"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
)

func UserView(w http.ResponseWriter, r *http.Request) {
	// 能来到这里，说明已经经过了中间件验证，身份无须在验证
	user, err := userOperate.GetName(r) // 获取用户身份
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrUnauthorized):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[401], "获取cookie错误")
		case errors.Is(err, pkg.ErrMissingParam):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "cookie-value基于base64编码解析失败")
		case errors.Is(err, pkg.ErrUserExists):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "获取用户基本信息错误")
		default:
			log.Println("未知错误")
		}
		return
	}
	// 响应后端的数据
	Data := model.UserProfile{
		Username: user.Username,
		IsAdmin:  user.Role == "0",
	}
	// 装入data，发送至响应体中
	Data.Users, err = userOperate.GetAllUser()
	if err != nil {
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "获取用户集合错误")
		return
	}
	// 后端响应数据已经准备完毕
	t, err := template.ParseFiles("internal/web-app/view/user.html")
	if err != nil {
		log.Println("模板解析失败", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "获取用户集合错误")
		return
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("模板渲染失败", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "获取用户集合错误")
		return
	}
}
func UserViewUpdate(w http.ResponseWriter, r *http.Request) {
	// 能来到这里，说明已经经过了中间件验证，身份无须在验证
	user, err := userOperate.GetName(r) // 获取用户身份
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrUnauthorized):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[401], "获取cookie错误")
		case errors.Is(err, pkg.ErrMissingParam):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "cookie-value基于base64编码解析失败")
		case errors.Is(err, pkg.ErrUserExists):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "获取用户基本信息错误")
		default:
			log.Println("未知错误")
		}
		return
	}
	// 响应后端的数据
	Data := model.UserProfile{
		Username: user.Username,
		IsAdmin:  user.Role == "0",
	}
	// 装入data，发送至响应体中
	Data.Users, err = userOperate.GetAllUser()
	err = json.NewEncoder(w).Encode(Data)
	if err != nil {
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "获取用户集合错误")
		return
	}

}
