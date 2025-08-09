package routes

import (
	"github.com/Haizhitao/blog_system/controllers"
	"github.com/Haizhitao/blog_system/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	//注册&登录
	user := r.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
	}

	//文章路由
	posts := r.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)
		authPosts := posts.Group("")
		authPosts.Use(middleware.AuthloginMiddleware())
		{
			authPosts.POST("", controllers.CreatePost)
			ownerPosts := authPosts.Group("")
			{
				ownerPosts.PUT("/:id", controllers.UpdatePost)
				ownerPosts.DELETE("/:id", controllers.DeletePost)
			}
		}
	}

	//评论路由
	comments := r.Group("/comments")
	{
		comments.GET("/post/:id", controllers.GetCommentsByPostId)
		comments.Use(middleware.AuthloginMiddleware())
		comments.POST("", controllers.CreateComment)
	}

}
