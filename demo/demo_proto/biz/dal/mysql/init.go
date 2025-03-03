package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// 提取 DSN 生成逻辑到独立函数
func getDSN() string {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	if mysqlUser == "" || mysqlPassword == "" || mysqlHost == "" || mysqlDatabase == "" {
		panic("missing required environment variables: MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, or MYSQL_DATABASE")
	}

	return fmt.Sprintf(conf.GetConf().MySQL.DSN,
		mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
}

// Init 开启数据库连接
func Init() error {
	dsn := getDSN()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}

	// 自动迁移数据库
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		return fmt.Errorf("failed to auto migrate database: %v", err)
	}

	log.Printf("Database initialized successfully with DSN: %s", dsn)
	return nil
}
