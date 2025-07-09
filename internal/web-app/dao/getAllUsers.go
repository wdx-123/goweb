package dao

import (
	"GoWeb/internal/web-app/model"
	"log"
)

func GetAllUsers() ([]model.User, error) {
	// 上锁
	DBmutex.Lock()
	defer DBmutex.Unlock()
	// 查询
	query := "select id,username,role,created_at from user"
	rows, err := DB.Query(query)
	if err != nil {
		log.Println("查询出错:", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Username, &user.Role, &user.CreatedAt)
		if err != nil {
			log.Println("行打印出错")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
