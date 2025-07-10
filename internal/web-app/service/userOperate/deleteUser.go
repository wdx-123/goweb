package userOperate

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/pkg"
	"log"
	"strconv"
)

func DeleteUser(idStr string) error {
	// 判断
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("参数错误，无法转化成id")
		return err
	}
	// 调用dao
	err = dao.DeleteUserByID(id)
	// 返回
	if err != nil {
		log.Println("删除失败：", err)
		return pkg.ErrDatabaseError
	}
	return nil
}
