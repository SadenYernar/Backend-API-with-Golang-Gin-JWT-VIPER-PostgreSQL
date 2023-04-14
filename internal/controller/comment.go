package controller

import (
	"net/http"
	"project/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createComment(ctx *gin.Context) {
	var commentData model.Comment

	if err := ctx.ShouldBindJSON(&commentData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	//	username := ctx.GetString("username")
	
}

func (h *Handler) likeComment(ctx *gin.Context) {}