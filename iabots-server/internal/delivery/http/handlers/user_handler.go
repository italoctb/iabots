package handlers

import (
	"iabots-server/internal/domain/usecases/user"
	"iabots-server/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *user.CreateUserUseCase
}

func NewUserHandler(usecase *user.CreateUserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req user.CreateUserParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendError(ctx, err)
		return
	}

	_, err := h.usecase.Execute(req)
	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}
