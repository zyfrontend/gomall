package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 输入
type LoginRequest struct {
}

// 输出
type LoginResponse struct {
	message string
	data    string
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}
