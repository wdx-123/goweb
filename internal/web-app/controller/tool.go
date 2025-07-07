package controller

import (
	"GoWeb/internal/web-app/pkg"
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, code int, msg string) { // 寄送错误信息
	w.WriteHeader(code)                               // 正确的话，默认200可以不写，但错误必须更改
	json.NewEncoder(w).Encode(map[string]interface{}{ // map天然适配转化为json
		"code":  code,
		"error": msg,
	})
}

func SendSuccessResponse(w http.ResponseWriter, msg string) { // 寄送正确信息
	w.WriteHeader(pkg.StatusCodeMap[200])
	json.NewEncoder(w).Encode(map[string]interface{}{ // map天然适配转化为json
		"code":  0,
		"error": msg,
	})
}
