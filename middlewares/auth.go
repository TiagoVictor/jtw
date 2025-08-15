package middlewares

import (
	"jwt-authentication-golang/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenSting := context.GetHeader("Authorization")
		if tokenSting == "" {
			context.JSON(401, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}

		err := auth.ValidateJWT(tokenSting)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
