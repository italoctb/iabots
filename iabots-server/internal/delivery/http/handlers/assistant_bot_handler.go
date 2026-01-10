package handlers

import (
	"net/http"

	. "iabots-server/internal/domain/entities"
	. "iabots-server/internal/domain/usecases/assistant_bot"

	. "iabots-server/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AssistantBotHandler struct {
	createUsecase *CreateAssistantBotUseCase
}

func NewAssistantBotHandler(createUsecase *CreateAssistantBotUseCase) *AssistantBotHandler {
	return &AssistantBotHandler{createUsecase: createUsecase}
}

func (h *AssistantBotHandler) CreateBot(ctx *gin.Context) {
	customerIDStr := ctx.Param("customerId")
	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		SendError(ctx, BadRequest("customerId inválido"))
		return
	}

	var req struct {
		Name            string `json:"name"`
		ModelProviderID string `json:"modelProviderId"`

		ContextMessage string `json:"contextMessage"`
		MaxTokens      int    `json:"maxTokens"`
		FreezeTime     int    `json:"freezeTime"`

		Status          string  `json:"status"`
		OnlineStartTime *string `json:"onlineStartTime"`
		OnlineEndTime   *string `json:"onlineEndTime"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		SendError(ctx, BadRequest("body inválido"))
		return
	}

	mpID, err := uuid.Parse(req.ModelProviderID)
	if err != nil {
		SendError(ctx, BadRequest("modelProviderId inválido"))
		return
	}

	bot, err := h.createUsecase.Execute(CreateAssistantBotParams{
		CustomerID:      customerID,
		Name:            req.Name,
		ModelProviderID: mpID,
		ContextMessage:  req.ContextMessage,
		MaxTokens:       req.MaxTokens,
		FreezeTime:      req.FreezeTime,
		Status:          AssistantStatus(req.Status),
		OnlineStartTime: req.OnlineStartTime,
		OnlineEndTime:   req.OnlineEndTime,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, bot)
}
