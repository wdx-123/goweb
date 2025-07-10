package CUserOperate

import (
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/userOperate"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("创建用户列表")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("解析请求失败:", err)
		pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "解析请求失败")
		return
	}
	err = userOperate.CreateUser(user)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrUserExists):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "用户存在")
		case errors.Is(err, pkg.ErrPasswordMismatch):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "用户密码错误")
		case errors.Is(err, pkg.ErrMissingParam):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "用户名 或 密码 不能为空")
		default:
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "注册失败")
		}
		return
	}
	// 简单验证
	pkg.SendSuccessResponse(w, "用户"+user.Username+"已添加成功")
}
