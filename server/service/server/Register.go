package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 输入
type RegisterRequest struct {
}

// 输出
type RegisterResponse struct {
	message string
	data    string
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register",
	})
}
