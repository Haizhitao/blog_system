package controllers

import (
	"github.com/Haizhitao/blog_system/config"
	"github.com/Haizhitao/blog_system/database"
	"github.com/Haizhitao/blog_system/models"
	"github.com/Haizhitao/blog_system/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入有效的用户名、密码及邮箱"})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "加密密码失败"})
		return
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	username := data["username"]
	password := data["password"]
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不能为空"})
		return
	}
	var storeUser models.User
	if err := database.DB.Where("username = ?", username).First(&storeUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	cfg := config.LoadConfig()
	tokenString, err := util.GenerateToken(storeUser.ID, storeUser.Username, cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
