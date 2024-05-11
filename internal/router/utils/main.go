package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidationErrorResponse struct {
	ValidationError string `json:"validationError"`
}

func generateResponseFromError(err error) ValidationErrorResponse {
	return ValidationErrorResponse{ValidationError: err.Error()}
}

func ValidateJSONBody(payload any, context *gin.Context) error {
	if err := context.ShouldBindJSON(payload); err != nil {
		context.JSON(http.StatusUnprocessableEntity, generateResponseFromError(err))
		return errors.New("invalid request")
	}
	return nil
}

func ValidateQueryParams(payload any, context *gin.Context) error {
	if err := context.ShouldBind(payload); err != nil {
		context.JSON(http.StatusUnprocessableEntity, generateResponseFromError(err))
		return errors.New("invalid request")
	}
	return nil
}
