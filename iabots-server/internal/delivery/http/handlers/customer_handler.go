package handlers

import (
	"net/http"

	"iabots-server/internal/delivery/http/middlewares"
	. "iabots-server/internal/domain/usecases/customer"
	"iabots-server/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerHandler struct {
	createCustomerUseCase *CreateCustomerUseCase
}

func NewCustomerHandler(createCustomerUseCase *CreateCustomerUseCase) *CustomerHandler {
	return &CustomerHandler{createCustomerUseCase: createCustomerUseCase}
}

type createCustomerRequest struct {
	Name     string `json:"name"`
	Whatsapp string `json:"whatsapp"`
}

func (h *CustomerHandler) CreateCustomer(ctx *gin.Context) {
	var req createCustomerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendError(ctx, err)
		return
	}

	raw, ok := ctx.Get(middlewares.ContextUserIDKey)
	if !ok {
		utils.SendError(ctx, utils.Unauthorized("usuário não autenticado"))
		return
	}

	userID, ok := raw.(uuid.UUID)
	if !ok {
		utils.SendError(ctx, utils.Unauthorized("usuário não autenticado"))
		return
	}

	_, err := h.createCustomerUseCase.Execute(CreateCustomerParams{
		Name:     req.Name,
		Whatsapp: req.Whatsapp,
		UserID:   userID,
	})
	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}
