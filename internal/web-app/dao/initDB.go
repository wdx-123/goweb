package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// InitDB 初始化数据库
func InitDB(dsn string) error {
	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// 测试链接
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("数据库ping失败，%v", err)
	}
	fmt.Println("数据库连接成功")
	return nil
}
