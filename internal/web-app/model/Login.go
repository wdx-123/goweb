package model

type LoginReq struct {
	Username   string `json:"username" db:"username"` // 用户名
	Password   string `json:"password"`               // 密码
	RememberMe bool   `json:"rememberMe"`             // 记住我，那个按钮
}
