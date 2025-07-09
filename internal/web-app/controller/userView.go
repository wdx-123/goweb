package controller

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

func UserView(w http.ResponseWriter, r *http.Request) {
	// 能来到这里，说明已经经过了中间件验证，身份无须在验证
	// 首先获取本用户
	cookie, err := r.Cookie("user-session")
	if err != nil {
		log.Println("获取cookie错误", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[401], "获取cookie错误")
		return
	}
	u, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Println("cookie-value基于base64编码解析失败: ", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "cookie-value基于base64编码解析失败")
		return
	}
	username := string(u)
	// 获取用户名字身份
	user, err := dao.FindUser(username)
	if err != nil {
		log.Println("获取用户基本信息错误", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "获取用户基本信息错误")
		return
	}
	// 响应后端的数据
	Data := model.UserProfile{
		Username: user.Username,
		IsAdmin:  user.Role == "0",
	}
	log.Println("姓名：", user.Username, " 管理员：", user.Role == "0")
	// 获取用户安全身份的集合
	var safeUser []model.SafeUser
	users, err := dao.GetAllUsers()
	if err != nil {
		log.Println("获取用户集合错误", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "获取用户集合错误")
		return
	}
	log.Println("查询人员集合：")
	for _, u := range users {
		safeUser = append(safeUser, u.ToSafeUser())
		log.Println("姓名：", u.ToSafeUser().Username, "创建时间：", u.CreatedAt, " 身份：", u.Role)
	}
	// 装入data，发送至响应体中
	Data.Users = safeUser
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
	//	pkg.SendSuccessResponse(w, "调用成功")
}
