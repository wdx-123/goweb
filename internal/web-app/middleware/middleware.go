package middleware

import (
	"GoWeb/internal/web-app/dao"
	"GoWeb/internal/web-app/pkg"
	"encoding/base64"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("进入认证中间件...")
		// 1、从Cookie中获取会话信息
		cookie, err := r.Cookie("user-session")
		if err != nil {
			log.Println("未找到会话cookie: ", err)
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[401], "请先等录")
			return
		}
		u, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			log.Println("cookie-value基于base64编码解析失败: ", err)
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "cookie-value基于base64编码解析失败")
			return
		}
		username := string(u)
		log.Println("恭喜 ", username, " 登入成功")
		// 2、验证用户是否存在
		if _, err = dao.FindUser(username); err != nil {
			log.Println("该用户不存在: ", err)
			return
		}
		log.Println("认证通过，您可以继续访问")
		// 3、验证通过，请继续接下来的请求
		next.ServeHTTP(w, r)
	})
}
