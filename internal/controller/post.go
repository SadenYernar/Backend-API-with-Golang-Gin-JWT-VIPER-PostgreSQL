package controller

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/dto"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(ctx *gin.Context) {
	var postData model.Post

	if err := ctx.ShouldBindJSON(&postData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// username := ctx.GetString("username")

	// userUUID, err := h.service.GetUserInfoServiceByUsername(username)

	// postData.Uuid = userUUID.Uuid

	status, post, err := h.service.CreatePostService(postData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	responce := dto.PostDTO(post)

	ctx.JSON(status, responce)
}

func (h *Handler) reactionToPost(ctx *gin.Context) {
	var reaction model.LikePost

	if err := ctx.ShouldBindJSON(&reaction); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	likeStatus, err := h.service.LikePostService(reaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(likeStatus, reaction)
}
