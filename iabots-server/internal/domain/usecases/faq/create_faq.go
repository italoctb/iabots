package faq

import (
	"strings"

	. "iabots-server/internal/domain/entities"
	. "iabots-server/internal/domain/repositories"
	. "iabots-server/pkg/utils"

	"github.com/google/uuid"
)

type CreateFaqParams struct {
	BotID    uuid.UUID `json:"botId"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
}

type CreateFaqUseCase struct {
	repository    FaqRepository
	botRepository AssistantBotRepository
}

func NewCreateFaqUseCase(repository FaqRepository, botRepository AssistantBotRepository) *CreateFaqUseCase {
	return &CreateFaqUseCase{repository: repository, botRepository: botRepository}
}

func (usecase *CreateFaqUseCase) Execute(params CreateFaqParams) (*Faq, error) {
	q := strings.TrimSpace(params.Question)
	a := strings.TrimSpace(params.Answer)

	if params.BotID == uuid.Nil {
		return nil, BadRequest("bot_id inválido")
	}
	_, err := usecase.botRepository.FindByID(params.BotID)
	if err != nil {
		return nil, NotFound("bot_id não encontrado")
	}

	if q == "" || len(q) < 3 {
		return nil, Validation("question inválida")
	}
	if a == "" || len(a) < 3 {
		return nil, Validation("answer inválida")
	}

	faq := &Faq{
		ID:       uuid.New(),
		BotID:    params.BotID,
		Question: q,
		Answer:   a,
		IsActive: true,
	}

	if err := usecase.repository.Create(faq); err != nil {
		return nil, err
	}

	return faq, nil
}
