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

// json:xxx 用于 JSON序列化/反序列化
// db:xxx 用于 sqlx（数据库操作）
