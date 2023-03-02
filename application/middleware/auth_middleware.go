package middleware

import (
	utils "english_exam_go/utils/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := utils.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
