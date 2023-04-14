package controller

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/middleware"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Route() *gin.Engine {
	router := gin.Default()

	router.GET("/", h.Index)
	router.GET("/myprofile", h.myProfile)

	post := router.Group("/post").Use(middleware.AuthMiddleware())
	{
		post.POST("/create-post", h.createPost)
		post.POST("/like", h.likePost)
	}

	comment := router.Group("/comment")
	{
		comment.POST("/create", h.createComment)
		comment.POST("/like", h.likeComment)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn) // авториизация
		auth.POST("/sign-up", h.signUp) // регистрация
		auth.POST("/logout", h.logout)
	}

	return router
}
