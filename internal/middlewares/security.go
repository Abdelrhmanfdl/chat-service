package middlewares

import (
	"net/http"
	"strings"

	jwtUtils "chat-chat-go/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "authorization header not provided"})
			ctx.Abort()
			return
		}

		claims, err := jwtUtils.ValidateJWT(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.Subject)
		ctx.Next()
	}
}
