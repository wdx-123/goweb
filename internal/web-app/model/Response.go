package model

type Response struct {
	Code int         `json:"code"` // 状态码，0表示成功，非0表示错误
	Data interface{} `json:"data"` // 响应数据，使用接口类型支持任意数据格式
	Msg  string      `json:"msg"`  // 状态消息，描述操作结果
}
