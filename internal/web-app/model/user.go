package model

import "time"

// User 创建结构体
type User struct {
	ID        int       `json:"id"   db:"id"`
	Username  string    `json:"username"  db:"username"`
	Password  string    `json:"password"  db:"password"`
	Role      string    `json:"role"  db:"role"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
}

// SafeUser 用户安全模型，不会暴露密码
type SafeUser struct {
	ID        int       `json:"id"   db:"id"`
	Username  string    `json:"username"  db:"username"`
	Role      string    `json:"role"  db:"role"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
}

// json:xxx 用于 JSON序列化/反序列化
// db:xxx 用于 sqlx（数据库操作）

func (u *User) ToSafeUser() SafeUser {
	return SafeUser{
		ID:        u.ID,
		Username:  u.Username,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
