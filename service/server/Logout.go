package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 输入
type LogoutRequest struct {
}

// 输出
type LogoutResponse struct {
	message string
	data    string
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}
