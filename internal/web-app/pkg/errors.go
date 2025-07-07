package pkg

import "errors"

// 用户相关错误（400-409）
var (
	ErrInvalidUsername  = errors.New("无效的用户名") // 400
	ErrInvalidPassword  = errors.New("无效的密码")  // 400
	ErrUserNotFound     = errors.New("用户不存在")  // 404
	ErrUserExists       = errors.New("用户名已存在") // 409
	ErrPasswordMismatch = errors.New("密码不匹配")  // 401
)

// 认证授权相关错误（401-403）
var (
	ErrUnauthorized     = errors.New("未授权，请先登录") // 401
	ErrPermissionDenied = errors.New("权限不足")     // 403
	ErrInvalidToken     = errors.New("无效的令牌")    // 401
	ErrTokenExpired     = errors.New("令牌已过期")    // 401
)

// 请求相关错误（400-405）
var (
	ErrInvalidRequest   = errors.New("无效的请求")    // 400
	ErrMissingParam     = errors.New("缺少必要参数")   // 400
	ErrMethodNotAllowed = errors.New("不支持的请求方法") // 405
)

// 资源相关错误（404-409）
var (
	ErrResourceNotFound = errors.New("资源不存在") // 404
	ErrResourceConflict = errors.New("资源冲突")  // 409
	ErrResourceLocked   = errors.New("资源被锁定") // 423
)

// 服务器内部错误（500-503）
var (
	ErrInternalServer     = errors.New("服务器内部错误") // 500
	ErrDatabaseError      = errors.New("数据库错误")   // 500
	ErrServiceUnavailable = errors.New("服务暂时不可用") // 503
)
