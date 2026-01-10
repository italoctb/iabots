package user

import (
	. "iabots-server/internal/domain/entities"
	"iabots-server/internal/domain/repositories"
	"iabots-server/pkg/utils"
	"iabots-server/pkg/validators"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewCreateUserUseCase(userRepo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (usecase *CreateUserUseCase) Execute(input CreateUserParams) (*User, error) {
	if err := validators.ValidateEmail(input.Email); err != nil {
		return nil, utils.Validation(err.Error())
	}

	if err := validators.ValidatePassword(input.Password); err != nil {
		return nil, utils.Validation(err.Error())
	}

	existing, _ := usecase.userRepo.FindByEmail(input.Email)
	if existing != nil {
		return nil, utils.Validation("email já cadastrado")
	}

	user := &User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     RoleCustomer,
	}

	if err := usecase.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
