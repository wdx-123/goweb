package model

type RegisterReq struct {
	Username        string `json:"username"`        // 用户名，JSON字段名为username
	Password        string `json:"password"`        // 密码，JSON字段名为password
	ConfirmPassword string `json:"confirmPassword"` // 确认密码，JSON字段名为confirmPassword
}
