package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

// AppError = erro tipado (domínio/validação/etc)
type AppError struct {
	Code    int
	Message string
	Err     error // opcional: erro original (para log)
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Helpers "bonitos" pra padronizar
func BadRequest(msg string) error {
	return &AppError{Code: http.StatusBadRequest, Message: msg}
}

func Validation(msg string) error {
	return &AppError{Code: http.StatusUnprocessableEntity, Message: msg}
}

func Unauthorized(msg string) error {
	if msg == "" {
		msg = "unauthorized"
	}
	return &AppError{Code: http.StatusUnauthorized, Message: msg}
}

func Forbidden(msg string) error {
	if msg == "" {
		msg = "forbidden"
	}
	return &AppError{Code: http.StatusForbidden, Message: msg}
}

func NotFound(msg string) error {
	if msg == "" {
		msg = "not found"
	}
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

// SendError padroniza resposta e loga o erro real
func SendError(c *gin.Context, err error) {
	log.Printf("[ERROR] %s %s -> %v", c.Request.Method, c.Request.URL.Path, err)

	// Se já for AppError, respeita code+message
	var appErr *AppError
	if errors.As(err, &appErr) {
		c.AbortWithStatusJSON(appErr.Code, APIError{
			Code:    appErr.Code,
			Message: appErr.Message,
		})
		return
	}

	// fallback (erro inesperado)
	c.AbortWithStatusJSON(http.StatusInternalServerError, APIError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	})
}
