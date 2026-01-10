package user

import (
	"iabots-server/internal/delivery/http/handlers"
	. "iabots-server/internal/domain/usecases/user"
	"iabots-server/internal/infra/database"
	"iabots-server/internal/infra/database/repositories"
)

type UserModule struct {
	Handler *handlers.UserHandler
}

func NewUserModule(db *database.Database) *UserModule {
	userRepo := repositories.NewUserGormRepository(db.DB)
	usecase := NewCreateUserUseCase(userRepo)
	handler := handlers.NewUserHandler(usecase)

	return &UserModule{
		Handler: handler,
	}
}
