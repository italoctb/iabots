package assistantbot

import (
	. "iabots-server/internal/delivery/http/handlers"
	. "iabots-server/internal/domain/usecases/assistant_bot"
	. "iabots-server/internal/infra/database"
	. "iabots-server/internal/infra/database/repositories"
)

type AssistantBotModule struct {
	Handler *AssistantBotHandler
}

func NewAssistantBotModule(db *Database) *AssistantBotModule {
	/// Injection of binds - Assistant Bot Module
	/// --------------------------------------------------
	///  Repositories
	/// --------------------------------------------------
	assistantBotRepo := NewAssistantBotGormRepository(db.DB)
	customerRepo := NewCustomerGormRepository(db.DB)
	modelProviderRepo := NewModelProviderGormRepository(db.DB)

	/// --------------------------------------------------
	///  Use Cases
	/// --------------------------------------------------
	createAssistantBotUseCase := NewCreateAssistantBotUseCase(assistantBotRepo, customerRepo, modelProviderRepo)
	/// --------------------------------------------------
	///  Handlers
	/// --------------------------------------------------
	handler := NewAssistantBotHandler(createAssistantBotUseCase)

	return &AssistantBotModule{
		Handler: handler,
	}
}
