package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	addMiddleware(router)
	addRoutes(router)
	router.Run()
}

func addMiddleware(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
}

func addRoutes(router *gin.Engine) {
	addGetPingRoute(router)
	addTodoRouteGroup(router)
}
