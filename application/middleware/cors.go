package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetCors() gin.HandlerFunc {
	f := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"X-Requested-With",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})

	return f
}
