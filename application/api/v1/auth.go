package v1

import (
	dtos "english_exam_go/domain/dtos/user"
	"english_exam_go/domain/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	as services.IAuthService
}

func CreateAuthController(as services.IAuthService) *AuthController {
	fmt.Println("Create auth controller")
	return &AuthController{as: as}
}

func (ac *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest dtos.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := ac.as.Login(c, loginRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}

func (ac *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Register router"})
	}
}
