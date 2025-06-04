package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

// LoginHandler 处理页面登入
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/login.html")
	if err != nil {
		fmt.Println("出现错误: ", err)
	}

	t.Execute(w, nil)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/web-app/view/register.html")
	if err != nil {
		fmt.Println("出现错误: ", err)
	}

	t.Execute(w, nil)
}
