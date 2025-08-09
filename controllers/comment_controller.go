package controllers

import (
	"github.com/Haizhitao/blog_system/database"
	"github.com/Haizhitao/blog_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentsByPostId(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	var post models.Post
	if err := database.DB.Model(&models.Post{}).Preload("Comments").First(&post, postId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	type simpleComment struct {
		UserID  uint
		PostID  uint
		Content string
	}
	var comments []simpleComment
	for _, comment := range post.Comments {
		comments = append(comments, simpleComment{
			UserID:  comment.UserID,
			PostID:  comment.PostID,
			Content: comment.Content,
		})
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据不能为空"})
		return
	}
	if postId, err := strconv.Atoi(data["post_id"].(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post_id无效"})
		return
	} else {
		comment.PostID = uint(postId)
	}
	if content, ok := data["content"].(string); ok {
		comment.Content = content
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content格式错误"})
		return
	}
	if userId, exist := c.Get("userid"); exist {
		comment.UserID = userId.(uint)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "认证异常，未获取到UserID"})
		return
	}
	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "评论成功"})
}
