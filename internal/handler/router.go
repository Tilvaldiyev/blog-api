package handler

import (
	_ "github.com/Tilvaldiyev/blog-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	user.POST("/register", h.createUser)
	user.POST("/login", h.loginUser)

	apiV1.Use(h.authMiddleware())
	apiV1.GET("/user/posts", h.userPosts)

	return router
}
