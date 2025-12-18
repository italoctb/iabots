package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

// Common errors
var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrValidation     = errors.New("validation error")
)

// SendError handles the error response
func SendError(c *gin.Context, err error) {
	apiErr := mapErrorToAPIError(err)
	c.AbortWithStatusJSON(apiErr.Code, apiErr)
}

func mapErrorToAPIError(err error) APIError {
	switch {
	case errors.Is(err, ErrUnauthorized):
		return APIError{Code: http.StatusUnauthorized, Message: ErrUnauthorized.Error()}
	case errors.Is(err, ErrForbidden):
		return APIError{Code: http.StatusForbidden, Message: ErrForbidden.Error()}
	case errors.Is(err, ErrNotFound):
		return APIError{Code: http.StatusNotFound, Message: ErrNotFound.Error()}
	case errors.Is(err, ErrBadRequest):
		return APIError{Code: http.StatusBadRequest, Message: ErrBadRequest.Error()}
	case errors.Is(err, ErrValidation):
		return APIError{Code: http.StatusUnprocessableEntity, Message: ErrValidation.Error()}
	default:
		return APIError{Code: http.StatusInternalServerError, Message: ErrInternalServer.Error()}
	}
}
