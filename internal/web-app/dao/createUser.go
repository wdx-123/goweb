package dao

import "GoWeb/internal/web-app/model"

func CreateUser(user model.User) error {
	query := `
		insert into user (username,password,role) 
		values (?,?,?)
	`
	DB.Exec(query, user.Username, user.Password, user.Role)
	return nil
}
