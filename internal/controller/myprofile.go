package controller

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAllData struct {
	Info model.User
	Post []model.Post
}

func (h *Handler) myProfile(ctx *gin.Context) {
	username := ctx.GetString("username")

	userInfo, err := h.service.GetUserInfoServiceByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	status, userPost, err := h.service.GetUserPostsService(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	userData := &UserAllData{
		Info: userInfo,
		Post: userPost,
	}

	ctx.JSON(status, userData)
}
