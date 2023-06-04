package http_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessHandle(data interface{}, c *gin.Context) {
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Thành công", "data": data})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Thành công"})
	}
	return
}

func SuccessHandlePaginate(data interface{}, total int, c *gin.Context) {
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Thành công", "data": data, "total": total})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Thành công"})
	}
	return
}

func CreatedHandle(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Thành công"})
	return
}

func NoContentHandle(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
	return
}
