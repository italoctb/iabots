package app

import (
	"iabots-server/internal/app/user"
	"iabots-server/internal/infra/database"
)

// AppModules centraliza todos os módulos do sistema
type AppModules struct {
	User *user.UserModule
}

// NewAppModules inicializa e retorna todos os módulos
func NewAppModules(db *database.Database) *AppModules {
	return &AppModules{
		User: user.NewUserModule(db),
	}
}
