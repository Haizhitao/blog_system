package database

import (
	"github.com/Haizhitao/blog_system/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB() error {
	conf := config.LoadConfig()
	var dialector gorm.Dialector
	dialector = mysql.Open(conf.DBDSN)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("数据库连接失败！")
	}
	DB = db
	return nil
}
