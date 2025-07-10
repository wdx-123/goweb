package userOperate

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/model"
	"GoWeb/internal/web-app/pkg"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func UpdateUser(idStr string, r *http.Request) error {
	// 判断
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("参数错误，无法转化成id")
		return pkg.ErrInvalidRequest
	}
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("解析失败")
		return pkg.ErrMissingParam
	}
	user.ID = id
	err = dao.UpdateUser(user)
	if err != nil {
		log.Println("数据库更新失败")
		return pkg.ErrDatabaseError
	}
	return nil
}
