package app

import (
	. "iabots-server/internal/app/assistant_bot"
	. "iabots-server/internal/app/customer"
	. "iabots-server/internal/app/faq"
	. "iabots-server/internal/app/user"
	. "iabots-server/internal/infra/database"
)

// AppModules centraliza todos os módulos do sistema
type AppModules struct {
	User         *UserModule
	Customer     *CustomerModule
	AssistantBot *AssistantBotModule
	Faq          *FaqModule
}

// NewAppModules inicializa e retorna todos os módulos
func NewAppModules(db *Database) *AppModules {
	return &AppModules{
		User:         NewUserModule(db),
		Customer:     NewCustomerModule(db),
		AssistantBot: NewAssistantBotModule(db),
		Faq:          NewFaqModule(db),
	}
}
