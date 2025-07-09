package pkg

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SendErrorResponse(w http.ResponseWriter, code int, msg string) { // 寄送错误信息
	w.WriteHeader(code)                               // 正确的话，默认200可以不写，但错误必须更改
	json.NewEncoder(w).Encode(map[string]interface{}{ // map天然适配转化为json
		"code":  code,
		"error": msg,
	})
}

func SendSuccessResponse(w http.ResponseWriter, msg string) { // 寄送正确信息
	w.WriteHeader(StatusCodeMap[200])
	json.NewEncoder(w).Encode(map[string]interface{}{ // map天然适配转化为json
		"code":  0,
		"error": msg,
	})
}
func SetSessionCookie(w http.ResponseWriter, username string, remeberMe bool) {
	log.Println("用户：", username, "申请登入")
	username = base64.URLEncoding.EncodeToString([]byte(username)) // cookie-value仅支持ascii储存
	// 设置cookie，登入成功用于 后期身份认证
	Cookie := &http.Cookie{
		Name:     "user-session", // cookie名称
		Value:    username,       // 存储名称
		Path:     "/",            // 全站通用
		HttpOnly: true,
		Secure:   false,
	}
	// 勾选"记住我",设置7天
	if remeberMe {
		Cookie.Expires = time.Now().Add(7 * 24 * time.Hour) // 7天后过期
		Cookie.MaxAge = 7 * 24 * 3600
	}
	// 若不同国Expires进行设计，浏览器关闭之后，cookie就会消失
	http.SetCookie(w, Cookie)
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "user-Cookie",
		Value:  "",
		Path:   "/", // 确保能找到之前生效的路径
		MaxAge: -1,  // 立即注销
	})
	SendSuccessResponse(w, "登出成功")
}
