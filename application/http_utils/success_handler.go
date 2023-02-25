package http_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessHandle(data interface{}, c *gin.Context) {
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
	return
}

func CreatedHandle(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
	return
}

func NoContentHandle(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
	return
}
