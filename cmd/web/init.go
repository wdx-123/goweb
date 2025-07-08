package main

import (
	"log"
	"os"
)

// init 初始化日志调用
func init() {
	// 配置
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 时间+【文件名+行号】

	// 创建日志目录
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("创建目录失败:%v", err)
	}

	// 2、打开日志文件
	logFile, err := os.OpenFile(
		"logs/app.log", // 具体文件路径
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil { // 检查文件打开错误
		log.Fatalf("打开日志文件失败：%v", err)
	}
	// 输出
	log.SetOutput(logFile)
}
