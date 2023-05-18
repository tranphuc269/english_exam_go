package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/user"
	"english_exam_go/domain/services"
	auth_utils "english_exam_go/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
			c.JSON(http.StatusInternalServerError,
				gin.H{"code": 401, "message": "Tài khoản hoặc mật khẩu không chính xác", "status": http.StatusUnauthorized})
			return
		}
		authRes, err := ac.as.Login(c, loginRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"code": 401, "message": "Tài khoản hoặc mật khẩu không chính xác", "status": http.StatusUnauthorized})
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
		http_utils.SuccessHandle(authRes, c)
	}
}

func (ac *AuthController) GetTeachers() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := ac.as.Teachers(c)
		http_utils.SuccessHandle(response, c)
	}
}

func (ac *AuthController) GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := ac.as.Students(c)
		http_utils.SuccessHandle(response, c)
	}
}

func (ac *AuthController) Me() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		claim, _ := auth_utils.ParseToken(tokenString)
		//if err != nil {
		//	domain_exception.Handle(err, c)
		//}

		response, _ := ac.as.Me(c, claim.Email)
		//if err != nil {
		//	domain_exception.Handle(err, c)
		//}
		http_utils.SuccessHandle(response, c)
	}
}

func (ac *AuthController) UserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		id, err := strconv.Atoi(params)
		if err != nil {
			fmt.Printf("err %s\n", err)
			exception.Handle(err, c)
			return
		}

		response, _ := ac.as.FindById(c, id)
		http_utils.SuccessHandle(response, c)
	}
}

func (ac *AuthController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateRequest dtos.UpdateAccountRequest
		if err := c.ShouldBind(&updateRequest); err != nil {
			exception.Handle(err, c)
			return
		}
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		claim, _ := auth_utils.ParseToken(tokenString)
		//if err != nil {
		//	domain_exception.Handle(err, c)
		//}

		_ = ac.as.Update(c, updateRequest, claim.Email)
		//if err != nil {
		//	domain_exception.Handle(err, c)
		//}
		http_utils.SuccessHandle(true, c)
	}
}
