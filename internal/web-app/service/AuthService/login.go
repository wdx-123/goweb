package AuthService

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"log"
)

func LoginDetection(loginReq model.LoginReq) error { // 登入界面吗？
	// 获取数据库内信息
	user, err := dao.FindUser(loginReq.Username)

	// 判断是否存在该用户
	if err != nil { // 不存在
		log.Println("登入界面：不存在该用户")
		return err
	}

	// 判断用户密码是否正确
	if user.Password != loginReq.Password {
		log.Println("密码错误")
		return pkg.ErrPasswordMismatch
	}
	return nil
}
