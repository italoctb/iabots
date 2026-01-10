package handlers

import (
	"net/http"

	. "iabots-server/internal/delivery/http/middlewares"
	. "iabots-server/internal/domain/usecases/faq"
	. "iabots-server/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FaqHandler struct {
	createFaqUseCase *CreateFaqUseCase
}

func NewFaqHandler(createFaqUseCase *CreateFaqUseCase) *FaqHandler {
	return &FaqHandler{createFaqUseCase: createFaqUseCase}
}

func (h *FaqHandler) CreateFaq(ctx *gin.Context) {
	// identity (por enquanto só garante que existe)
	_, ok := ctx.Get(ContextUserIDKey)
	if !ok {
		SendError(ctx, Unauthorized("usuário não autenticado"))
		return
	}

	botIDStr := ctx.Param("botId")

	botID, err := uuid.Parse(botIDStr)
	if err != nil || botID == uuid.Nil {
		SendError(ctx, BadRequest("bot_id inválido"))
		return
	}

	var body struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		SendError(ctx, BadRequest("corpo da requisição inválido"))
		return
	}

	created, err := h.createFaqUseCase.Execute(CreateFaqParams{
		BotID:    botID,
		Question: body.Question,
		Answer:   body.Answer,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}
