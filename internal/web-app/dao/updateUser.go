package dao

import (
	"GoWeb/internal/web-app/model"
	"log"
)

func UpdateUser(user model.User) (err error) {
	DBmutex.Lock()
	defer DBmutex.Unlock()

	query := "update user set username = ?,role = ? where id = ?"
	_, err = DB.Exec(query, user.Username, user.Role, user.ID)
	if err != nil {
		log.Println("用户信息更新失败：", err)
		return err
	}
	return nil
}
