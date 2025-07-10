package userOperate

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"log"
	"regexp"
)

func CreateUser(user model.User) error { // 一户记得设计几个全局错误
	// 校验密码格式是否正确，就简单一点，直指核心
	if user.Password == "" || user.Username == "" {
		log.Println("密码与用户名不能为空")
		return pkg.ErrMissingParam
	}
	// 1、密码的格式，大于6位。--regexy
	reg, err := regexp.Compile(`\d{4,12}`)
	if err != nil {
		log.Println("注册界面-regexp编辑失败")
		return pkg.ErrInternalServer // 服务器内部错误
	}
	resBool := reg.MatchString(user.Password)
	if resBool != true {
		log.Println("注册界面-密码格式错误，不符合(4-12位数字)")
		return pkg.ErrPasswordMismatch // 密码不匹配
	}

	// 校验该用户是否存在
	// 查重的呢
	_, err = dao.FindUser(user.Username)
	if err == nil {
		log.Printf("注册界面-已经存在该用户")
		return pkg.ErrUserExists // 用户已存在
	}

	// 插入 - 成功
	// 插入数据库
	err = dao.CreateUser(model.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	})

	return err

}
