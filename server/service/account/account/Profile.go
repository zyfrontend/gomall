package serviceAccount

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "profile",
	})
}
