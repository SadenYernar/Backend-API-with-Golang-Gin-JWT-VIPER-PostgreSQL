package controller

import (
	"fmt"
	"net/http"
	"project/internal/dto"
	"project/internal/model"

	"github.com/gin-gonic/gin"
)

// авторизация
func (h *Handler) signIn(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := h.service.AuthorizationUserService(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cookie := http.Cookie{
		Name:  "session_name",
		Value: token,
	}
	http.SetCookie(ctx.Writer, &cookie)
}

// регистрация
func (h *Handler) signUp(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	status, userAllData, err := h.service.CreateUserService(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	responce := dto.UserDTO(userAllData)

	ctx.JSON(status, responce)
}

func (h *Handler) logout(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cookie := http.Cookie{
		Name:   "session_name",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(ctx.Writer, &cookie)
}
