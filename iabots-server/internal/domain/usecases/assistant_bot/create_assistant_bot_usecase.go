package assistantbot

import (
	"errors"
	. "iabots-server/internal/domain/entities"
	. "iabots-server/internal/domain/repositories"
	. "iabots-server/pkg/utils"
	. "iabots-server/pkg/validators"

	"github.com/google/uuid"
)

type CreateAssistantBotParams struct {
	CustomerID      uuid.UUID
	Name            string
	ModelProviderID uuid.UUID

	ContextMessage string
	MaxTokens      int
	FreezeTime     int

	Status          AssistantStatus
	OnlineStartTime *string
	OnlineEndTime   *string
}

type CreateAssistantBotUseCase struct {
	repository              AssistantBotRepository
	customerRepository      CustomerRepository
	modelProviderRepository ModelProviderRepository
}

func NewCreateAssistantBotUseCase(repository AssistantBotRepository, customerRepository CustomerRepository, modelProviderRepository ModelProviderRepository) *CreateAssistantBotUseCase {
	return &CreateAssistantBotUseCase{
		repository:              repository,
		customerRepository:      customerRepository,
		modelProviderRepository: modelProviderRepository,
	}
}

func (usecase *CreateAssistantBotUseCase) Execute(p CreateAssistantBotParams) (*AssistantBot, error) {
	if p.CustomerID == uuid.Nil {
		return nil, BadRequest("customer_id inválido")
	}
	_, err := usecase.customerRepository.FindByID(p.CustomerID)
	if err != nil {
		return nil, NotFound("customer_id não encontrado")
	}
	if p.ModelProviderID == uuid.Nil {
		return nil, BadRequest("model_provider_id inválido")
	}
	_, err = usecase.modelProviderRepository.FindByID(p.ModelProviderID)
	if err != nil {
		return nil, NotFound("model_provider_id não encontrado")
	}
	if p.Name == "" {
		return nil, BadRequest("name é obrigatório")
	}

	status := p.Status
	if status == "" {
		status = AssistantPaused
	}

	if !status.IsValid() {
		return nil, Validation("status inválido. Valores válidos: " + AllBotStatusesToString())
	}

	// se AUTO, faz sentido ter horário preenchido
	if status == AssistantAuto && (p.OnlineStartTime == nil || p.OnlineEndTime == nil) {
		return nil, Validation("para status=auto, informe onlineStartTime e onlineEndTime")
	}

	/// Usando a funçao ValidateHHMM para validar os horários, se fornecidos
	if p.OnlineStartTime != nil {
		if err := ValidateHHMM(*p.OnlineStartTime); err != nil {
			return nil, Validation("onlineStartTime inválido: " + err.Error())
		}
	}

	if p.OnlineEndTime != nil {
		if err := ValidateHHMM(*p.OnlineEndTime); err != nil {
			return nil, Validation("onlineEndTime inválido: " + err.Error())
		}
	}

	// exemplo: regra simples de maxTokens
	if p.MaxTokens < 0 || p.FreezeTime < 0 {
		return nil, Validation("maxTokens/freezeTime não podem ser negativos")
	}

	bot := &AssistantBot{
		ID:              uuid.New(),
		CustomerID:      p.CustomerID,
		Name:            p.Name,
		ModelProviderID: p.ModelProviderID,
		ContextMessage:  p.ContextMessage,
		MaxTokens:       p.MaxTokens,
		FreezeTime:      p.FreezeTime,
		Status:          status,
		OnlineStartTime: p.OnlineStartTime,
		OnlineEndTime:   p.OnlineEndTime,
	}

	if err := usecase.repository.Create(bot); err != nil {
		// se quiser tratar duplicidade por constraint depois
		return nil, errors.New(err.Error())
	}

	return bot, nil
}
