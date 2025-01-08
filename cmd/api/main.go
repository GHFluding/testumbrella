package main

import (
	"testtask/internal/app/services/handler"
	mw "testtask/internal/app/services/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(mw.CheckAdminMW)
	r.GET("/day", handler.DayAfterHandler())
	r.Run()
}
