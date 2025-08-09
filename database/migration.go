package database

import (
	"github.com/Haizhitao/blog_system/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	return err
}
