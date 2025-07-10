package CUserOperate

import (
	"GoWeb/internal/web-app/pkg"
	"GoWeb/internal/web-app/service/userOperate"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	err := userOperate.DeleteUser(idStr)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrDatabaseError):
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[501], "删除失败")
		default:
			pkg.SendErrorResponse(w, pkg.StatusCodeMap[404], "传递参数无效")
		}
		return
	}
	pkg.SendSuccessResponse(w, idStr+"已成功被删除")
}
