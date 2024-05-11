package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addGetPingRoute(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
