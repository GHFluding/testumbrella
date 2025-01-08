package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DayAfterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		day := int64(time.Since(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)
		c.String(http.StatusOK, fmt.Sprintf("days left: %d", day))
	}
}
