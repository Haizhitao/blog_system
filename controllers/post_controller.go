package controllers

import (
	"github.com/Haizhitao/blog_system/database"
	"github.com/Haizhitao/blog_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if title, ok := data["title"].(string); ok {
		post.Title = title
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title格式错误"})
		return
	}
	if content, ok := data["content"].(string); ok {
		post.Content = content
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content格式错误"})
		return
	}
	if userId, exist := c.Get("userid"); exist {
		post.UserID = userId.(uint)
		//database.DB.Model(&models.User{}).First(&post.User, "id = ?", post.UserID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "认证异常，未获取到UserID"})
		return
	}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "创建成功"})
}

func GetPosts(c *gin.Context) {
	type simplePost struct {
		gorm.Model
		ID      uint
		Title   string
		Content string
		UserID  uint
	}
	var posts []simplePost
	if err := database.DB.Model(&models.Post{}).Select("id", "created_at", "updated_at", "title", "content", "user_id").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	type simplePost struct {
		gorm.Model
		ID      uint
		Title   string
		Content string
		UserID  uint
	}
	var post simplePost
	postID := c.Param("id")
	postIDInt, _ := strconv.Atoi(postID)
	if err := database.DB.Model(&models.Post{}).Select("id", "created_at", "updated_at", "title", "content", "user_id").First(&post, postIDInt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, _ := strconv.Atoi(postID)
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "更新数据不能为空"})
		return
	}
	updateData := make(map[string]interface{})
	if title, ok := data["title"].(string); ok {
		updateData["title"] = title
	} else if data["title"] != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title格式错误"})
		return
	}

	if content, ok := data["content"].(string); ok {
		updateData["content"] = content
	} else if data["content"] != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content格式错误"})
		return
	}

	if len(updateData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有提供有效的更新字段"})
		return
	}

	result := database.DB.Model(&models.Post{}).Where("id = ?", postIDInt).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在或没有变化"})
		return
	}

	var updatedPost models.Post
	if err := database.DB.First(&updatedPost, postIDInt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取更新后的文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

func DeletePost(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	result := database.DB.Where("id = ?", id).Delete(&models.Post{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
