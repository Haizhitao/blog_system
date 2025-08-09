package main

import (
	"github.com/Haizhitao/blog_system/database"
	"github.com/Haizhitao/blog_system/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败：%v", err)
	}

	err = database.AutoMigrate(database.DB)
	if err != nil {
		log.Fatalf("数据表创建失败：%v", err)
	}

	r := gin.Default()
	//r.Use(middleware.Logger())
	routes.InitRoutes(r)

	err = r.Run(":8080")
	if err != nil {
		log.Printf("服务启动失败：%v", err)
	}
}
