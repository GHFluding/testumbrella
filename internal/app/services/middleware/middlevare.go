package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckAdminMW(c *gin.Context) {
	role := c.GetHeader("User-Role")
	if role == "admin" {
		log.Println("red button user detected")
	}
	c.Next()
}
