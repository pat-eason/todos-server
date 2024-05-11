package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pateason/todo-server/internal/router/utils"
	"github.com/pateason/todo-server/internal/services"
)

type RetrieveTodosRequestModel struct {
	Date *string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

type ManageTodoRequestModel struct {
	Title string `json:"title" binding:"required,max=2"`
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
	routeGroup.GET("", func(context *gin.Context) {
		var queryParams RetrieveTodosRequestModel
		err := utils.ValidateQueryParams(&queryParams, context)
		if err != nil {
			return
		}

		var queryDate *time.Time
		if queryParams.Date != nil {
			convertedDate, err := time.Parse("2006-01-02", *queryParams.Date)
			if err != nil {
				context.JSON(http.StatusInternalServerError, err)
				return
			}
			queryDate = &convertedDate
		}

		records, err := services.RetrieveTodos(services.RetrieveTodosModel{Date: queryDate})
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, records)
	})
}

func addRetrieveTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.GET(":todoId", func(context *gin.Context) {
		todoId := context.Param("todoId")

		record, err := services.RetrieveTodo(todoId)
		if err != nil {
			context.JSON(http.StatusNotFound, err)
			return
		}

		context.JSON(http.StatusOK, record)
	})
}

func addCreateTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.PUT("", func(context *gin.Context) {
		var payload ManageTodoRequestModel
		err := utils.ValidateJSONBody(&payload, context)
		if err != nil {
			return
		}

		record, err := services.CreateTodo(services.CreateTodoModel{
			Title: payload.Title,
		})
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, record)
	})
}

func addUpdateTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.POST(":todoId", func(context *gin.Context) {
		var payload ManageTodoRequestModel
		err := utils.ValidateJSONBody(&payload, context)
		if err != nil {
			return
		}

		//todoId := context.Param("todoId")

		// @TODO update the record

		context.JSON(http.StatusOK, payload)
	})
}

func addSetTodoStatusRoute(routeGroup *gin.RouterGroup) {
	routeGroup.POST(":todoId/status", func(context *gin.Context) {
		var payload SetTodoStatusRequestModel
		err := utils.ValidateJSONBody(&payload, context)
		if err != nil {
			return
		}

		todoId := context.Param("todoId")
		_, err = services.RetrieveTodo(todoId)
		if err != nil {
			context.JSON(http.StatusNotFound, err)
			return
		}

		record, err := services.UpdateTodoStatus(todoId, payload.IsActive)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
		}

		context.JSON(http.StatusOK, record)
	})
}

func addDeleteTodoRoute(routeGroup *gin.RouterGroup) {
	routeGroup.DELETE(":todoId", func(context *gin.Context) {
		todoId := context.Param("todoId")

		_, err := services.RetrieveTodo(todoId)
		if err != nil {
			context.JSON(http.StatusNotFound, err)
			return
		}

		err = services.DeleteTodo(todoId)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.Status(http.StatusNoContent)
	})
}
