package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pateason/todo-server/internal/transport/routes/utils"
	"net/http"
)

type RetrieveTodosRequest struct {
	Date string `form:"date"`
}

type ManageTodoRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags" binding:"required"`
}

type SetTodoStatusRequest struct {
	IsActive bool `json:"isActive" binding:"required"`
}

func addTodoRouteGroup(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	addRetrieveTodosRoute(todoGroup)
	addCreateTodoRoute(todoGroup)
	addUpdateTodoRoute(todoGroup)
	addSetTodoStatusRoute(todoGroup)
	addDeleteTodoRoute(todoGroup)
}

func addRetrieveTodosRoute(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", func(context *gin.Context) {
		var queryParams RetrieveTodosRequest
		utils.ValidateQueryParams(&queryParams, context)

		// @TODO retrieve the records

		context.JSON(http.StatusOK, queryParams)
	})
}

func addCreateTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.PUT("/", func(context *gin.Context) {
		var payload ManageTodoRequest
		utils.ValidateJSONBody(&payload, context)

		// @TODO create the record

		context.JSON(http.StatusOK, payload)
	})
}

func addUpdateTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/todo/:todoId", func(context *gin.Context) {
		var payload ManageTodoRequest
		utils.ValidateJSONBody(&payload, context)

		//todoId := context.Query("todoId")

		// @TODO update the record

		context.JSON(http.StatusOK, payload)
	})
}

func addSetTodoStatusRoute(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/todo/:todoId/status", func(context *gin.Context) {
		var payload SetTodoStatusRequest
		utils.ValidateJSONBody(&payload, context)

		//todoId := context.Query("todoId")

		// @TODO set the record status flag

		context.JSON(http.StatusOK, payload)
	})
}

func addDeleteTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.DELETE("/todo/:todoId", func(context *gin.Context) {
		//todoId := context.Query("todoId")

		// @TODO delete the record

		context.Status(http.StatusNoContent)
	})
}
