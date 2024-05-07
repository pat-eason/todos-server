package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pateason/todo-server/internal/database"
	"github.com/pateason/todo-server/internal/router/routes/utils"
	"github.com/pateason/todo-server/internal/services"
	"net/http"
)

type RetrieveTodosRequestModel struct {
	Date string `form:"date"`
}

type ManageTodoRequestModel struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags" binding:"required"`
}

type SetTodoStatusRequestModel struct {
	IsActive bool `json:"isActive" binding:"required"`
}

func addTodoRouteGroup(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	addRetrieveTodosRoute(todoGroup)
	addRetrieveTodoRoute(todoGroup)
	addCreateTodoRoute(todoGroup)
	addUpdateTodoRoute(todoGroup)
	addSetTodoStatusRoute(todoGroup)
	addDeleteTodoRoute(todoGroup)
}

func addRetrieveTodosRoute(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", func(context *gin.Context) {
		var queryParams RetrieveTodosRequestModel
		utils.ValidateQueryParams(&queryParams, context)

		records, err := services.RetrieveTodos(services.RetrieveTodosModel{})
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, records)
	})
}

func addRetrieveTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/:todoId", func(context *gin.Context) {
		todoId := context.Param("todoId")

		record, err := database.RetrieveTodoEntity(todoId)
		if err != nil {
			context.JSON(http.StatusNotFound, err)
			return
		}

		context.JSON(http.StatusOK, record)
	})
}

func addCreateTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.PUT("/", func(context *gin.Context) {
		var payload ManageTodoRequestModel
		utils.ValidateJSONBody(&payload, context)

		record, err := services.CreateTodo(services.CreateTodoModel{
			Title:   payload.Title,
			Content: payload.Content,
		})
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, record)
	})
}

func addUpdateTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/:todoId", func(context *gin.Context) {
		var payload ManageTodoRequestModel
		utils.ValidateJSONBody(&payload, context)

		//todoId := context.Param("todoId")

		// @TODO update the record

		context.JSON(http.StatusOK, payload)
	})
}

func addSetTodoStatusRoute(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/:todoId/status", func(context *gin.Context) {
		var payload SetTodoStatusRequestModel
		utils.ValidateJSONBody(&payload, context)

		//todoId := context.Param("todoId")

		// @TODO set the record status flag

		context.JSON(http.StatusOK, payload)
	})
}

func addDeleteTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.DELETE("/:todoId", func(context *gin.Context) {
		todoId := context.Param("todoId")

		_, err := database.RetrieveTodoEntity(todoId)
		if err != nil {
			context.JSON(http.StatusNotFound, err)
			return
		}

		err = database.DeleteTodoEntity(todoId)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.Status(http.StatusNoContent)
	})
}
