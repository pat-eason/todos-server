package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ValidationErrorResponse struct {
	ValidationError string `json:"validationError"`
}

func generateResponseFromError(err error) ValidationErrorResponse {
	return ValidationErrorResponse{ValidationError: err.Error()}
}

func ValidateJSONBody(payload any, context *gin.Context) {
	if err := context.ShouldBindJSON(payload); err != nil {
		context.JSON(http.StatusUnprocessableEntity, generateResponseFromError(err))
	}
}

func ValidateQueryParams(payload any, context *gin.Context) {
	if err := context.ShouldBind(payload); err != nil {
		context.JSON(http.StatusUnprocessableEntity, generateResponseFromError(err))
	}
}
