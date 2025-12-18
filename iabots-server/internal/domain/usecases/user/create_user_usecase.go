package user

import (
	"errors"
	"iabots-server/internal/domain/entities"
	"iabots-server/internal/domain/repositories"
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

func (uc *CreateUserUseCase) Execute(input CreateUserParams) (*entities.User, error) {
	if err := validators.ValidateEmail(input.Email); err != nil {
		return nil, err
	}

	if err := validators.ValidatePassword(input.Password); err != nil {
		return nil, err
	}

	existing, _ := uc.userRepo.FindByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	user := &entities.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     entities.RoleClient,
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
