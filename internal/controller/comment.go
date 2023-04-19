package controller

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/dto"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createComment(ctx *gin.Context) {
	var commentData model.Comment

	if err := ctx.ShouldBindJSON(&commentData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	status, comment, err := h.service.CreateCommentsInService(commentData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	responce := dto.CommentDTO(comment)

	ctx.JSON(status, responce)
}

func (h *Handler) reactionToComment(ctx *gin.Context) {
	var reaction model.LikeComment

	if err := ctx.ShouldBindJSON(&reaction); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	likeStatus, err := h.service.LikeCommentService(reaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(likeStatus, reaction)
}
