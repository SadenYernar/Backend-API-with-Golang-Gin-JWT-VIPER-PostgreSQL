package middleware

import (
	"project/internal/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		tokenString := ctx.GetHeader("Authorization")
		if err != nil || len(tokenString) == 0 {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		username, err := token.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("username", username)
		ctx.Next()
	}
}
