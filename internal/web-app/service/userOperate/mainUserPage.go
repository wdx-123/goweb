package userOperate

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"encoding/base64"
	"log"
	"net/http"
)

func GetName(r *http.Request) (user *model.User, err error) {
	// 首先获取本用户
	cookie, err := r.Cookie("user-session")
	if err != nil {
		log.Println("获取cookie错误", err)
		return nil, pkg.ErrUnauthorized
	}
	u, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Println("cookie-value基于base64编码解析失败: ", err)
		return nil, pkg.ErrMissingParam
	}
	username := string(u)
	// 获取用户名字身份
	user, err = dao.FindUser(username)
	if err != nil {
		log.Println("获取用户基本信息错误", err)
		return nil, pkg.ErrUserExists
	}
	log.Println("用户姓名：", username)
	return user, err
}

func GetAllUser() ([]model.SafeUser, error) {
	// 获取用户安全身份的集合
	var safeUser []model.SafeUser
	users, err := dao.GetAllUsers()
	if err != nil {
		log.Println("获取用户集合错误", err)
		return nil, pkg.ErrUserExists
	}
	log.Println("查询人员集合：")
	for _, u := range users {
		safeUser = append(safeUser, u.ToSafeUser())
	}
	return safeUser, err
}
