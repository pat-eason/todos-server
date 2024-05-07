package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/pateason/todo-server/internal/transport/routes"
)

func StartRouter() {
	router := gin.Default()
	routes.AddRoutes(router)
	router.Run()
}

// @todo create auth middleware

// @todo create validation middleware
