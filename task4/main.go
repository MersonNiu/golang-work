package main

import (
	"GoBlogProject/config"
	"GoBlogProject/handlers"
	"GoBlogProject/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	// handlers.AutoSetTable()
	// services.Register("admin", "123456", "admin@blog.com")
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/post/get", handlers.PostGetAllHandler)
	r.GET("/post/get/:id", handlers.PostGetIdHandler)
	r.GET("/comment/:id/get", handlers.CommentGetHandler)
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/me", handlers.MeHandler)
		auth.POST("/post/create", handlers.PostCreateHandler)
		auth.PUT("/post/:id/update", handlers.PostUpdateHandler)
		auth.DELETE("/post/:id/delete", handlers.PostDeleteHandler)
		auth.POST("/comment/:id/create", handlers.CommentCreateHandler)

	}
	r.Run(":8080")
}
