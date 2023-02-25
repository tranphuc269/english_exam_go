package v1

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
	Me(c *gin.Context)
}
