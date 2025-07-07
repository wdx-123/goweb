package pkg

import "net/http"

// StatusCodeMap 完整状态码映射表示例
var StatusCodeMap = map[int]int{
	// 成功类
	200: http.StatusOK,      // 操作成功
	201: http.StatusCreated, // 资源创建成功

	// 客户端错误类
	400: http.StatusBadRequest,         // 请求参数错误
	401: http.StatusUnauthorized,       // 未授权（需要登录）
	402: http.StatusPaymentRequired,    // 付费请求（电商场景）
	403: http.StatusForbidden,          // 禁止访问（权限不足）
	404: http.StatusNotFound,           // 资源不存在
	405: http.StatusMethodNotAllowed,   // 方法不允许
	409: http.StatusConflict,           // 资源冲突（如用户名重复）
	412: http.StatusPreconditionFailed, // 前提条件失败

	// 服务器错误类
	500: http.StatusInternalServerError, // 服务器内部错误
	501: http.StatusNotImplemented,      // 功能未实现
	503: http.StatusServiceUnavailable,  // 服务不可用
}

/*
统一通信标准：
	让后端业务状态与 HTTP 协议标准对齐，符合 RESTful 接口规范
代码解耦：
	业务逻辑可以自定义状态码，通过映射表适配 HTTP 协议，以后换协议时（如 gRPC），
	只需修改映射表，业务代码不用改。
解耦+翻译表+为了解决业务码与标准码的问题
*/
