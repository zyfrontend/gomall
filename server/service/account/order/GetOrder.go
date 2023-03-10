package serviceOrder

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "profile",
	})
}
