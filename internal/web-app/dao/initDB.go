package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

var DB *sqlx.DB
var DBmutex sync.Mutex

// InitDB 初始化数据库
// sqlx对操作数据库，进行功能增强
func InitDB(dsn string) error {
	DBmutex.Lock()
	defer DBmutex.Unlock()
	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// 测试链接
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("数据库ping失败，%v", err)
	}

	log.Println("数据库连接成功")
	return nil
}
