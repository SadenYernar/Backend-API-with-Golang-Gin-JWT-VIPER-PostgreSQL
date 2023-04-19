package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(ctx *gin.Context) {
	status, allPost, err := h.service.GetAllPostService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(status, allPost)
}
