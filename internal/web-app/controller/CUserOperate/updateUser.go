package CUserOperate

import (
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/userOperate"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	err := userOperate.UpdateUser(idStr, r)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrInvalidRequest):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "无效请求-id")
		case errors.Is(err, pkg.ErrMissingParam):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[400], "缺少必要参数")
		case errors.Is(err, pkg.ErrDatabaseError):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "数据库更新失败")
		default:
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[500], "更新失败")
			return
		}
	}
	pkg.SendSuccessResponse(w, "更新成功")
}
