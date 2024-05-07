package routes

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.Engine) {
	addGetPingRoute(router)
	addTodoRouteGroup(router)
}
