package dao

import (
	"GoWeb/internal/web-app/model"
	"log"
	"time"
)

func CreateUser(user model.User) error {
	query := `
		insert into user (username,password,role,created_at) 
		values (?,?,?,?)
	`
	_, err := DB.Exec(query, user.Username, user.Password, user.Role, time.Now())
	if err != nil {
		log.Println("dao中createUser包调用失败：", err)
	}
	return err
}
