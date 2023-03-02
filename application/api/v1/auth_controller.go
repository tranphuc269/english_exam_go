package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/user"
	"english_exam_go/domain/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	as services.IAuthService
}

func CreateAuthController(as services.IAuthService) *AuthController {
	return &AuthController{as: as}
}

func (ac *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest dtos.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			exception.Handle(err, c)
			return
		}
		authRes, err := ac.as.Login(c, loginRequest)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(*authRes, c)
	}
}

func (ac *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerRequest dtos.RegisterAccountRequest
		if err := c.ShouldBind(&registerRequest); err != nil {
			exception.Handle(err, c)
			return
		}
		authRes, err := ac.as.Register(c, registerRequest)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(*authRes, c)
	}
}

func (ac *AuthController) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		stringToken := c.GetHeader("Authorization")
		http_utils.SuccessHandle(stringToken, c)
	}
}
