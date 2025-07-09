package model

// UserProfile 用户档案
type UserProfile struct {
	Username string     `json:"username" db:"username"`
	IsAdmin  bool       `json:"isAdmin" db:"isAdmin"`
	Users    []SafeUser `json:"users" db:"users"`
}
