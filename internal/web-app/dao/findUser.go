package dao

import (
	"GoWeb/internal/web-app/model"
	"database/sql"
	"log"
)

func FindUser(name string) (*model.User, error) {
	DBmutex.Lock()
	defer DBmutex.Unlock()
	var user = &model.User{}
	// 如果没找到，就会报错
	query := `
		select id,username,password,role,created_at
		from user
		where username = ?
		`
	// 返回查找到的好东西
	err := DB.QueryRow(query, name).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("并未找到相应的信息：", err)
		} else {
			log.Println("FindUser包出错", err)
		}
		return user, err
	}
	return user, nil
}
