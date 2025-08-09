package middleware

import (
	"errors"
	"github.com/Haizhitao/blog_system/config"
	"github.com/Haizhitao/blog_system/database"
	"github.com/Haizhitao/blog_system/models"
	"github.com/Haizhitao/blog_system/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func AuthloginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			token, _ = c.Cookie("jwt_token")
		}
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未提供认证令牌",
			})
			return
		}
		// 3. 解析和验证token
		cfg := config.LoadConfig()
		claims, err := util.ParseToken(token, cfg.JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的令牌: " + err.Error(),
			})
			return
		}
		c.Set("username", claims.Username)
		c.Set("userid", claims.UserID)
		c.Next()
	}
}

func AuthOwerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		id, err := strconv.ParseUint(param, 10, 32)
		if err != nil || id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID无效"})
			c.Abort()
			return
		}

		userIdValue, exists := c.Get("userid")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
			c.Abort()
			return
		}

		userId, ok := userIdValue.(uint)
		if !ok || userId == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无效的用户ID"})
			c.Abort()
			return
		}

		var post models.Post
		if err := database.DB.First(&post, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "文章记录不存在"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库异常"})
			}
			c.Abort()
			return
		}

		if post.UserID != userId {
			c.JSON(http.StatusForbidden, gin.H{"error": "您不是此文章的作者"})
			c.Abort()
			return
		}

		c.Next()
	}
}
