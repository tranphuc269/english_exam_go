package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func CreateAuthController() *AuthController {
	fmt.Println("Create auth controller")
	return &AuthController{}
}

func (ac *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Login router"})
	}
}

func (ac *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Register router"})
	}
}
