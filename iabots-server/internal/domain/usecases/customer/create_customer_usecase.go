package customer

import (
	. "iabots-server/internal/domain/entities"
	"iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
)

type CreateCustomerParams struct {
	Name     string
	Whatsapp string
	UserID   uuid.UUID // vindo do identity middleware
}

type CreateCustomerUseCase struct {
	customerRepo     repositories.CustomerRepository
	userCustomerRepo repositories.UserCustomerRepository
}

func NewCreateCustomerUseCase(
	customerRepo repositories.CustomerRepository,
	userCustomerRepo repositories.UserCustomerRepository,
) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		customerRepo:     customerRepo,
		userCustomerRepo: userCustomerRepo,
	}
}

func (uc *CreateCustomerUseCase) Execute(params CreateCustomerParams) (*Customer, error) {
	customer := &Customer{
		ID:          uuid.New(),
		CompanyName: params.Name,
		Whatsapp:    params.Whatsapp,
	}

	if err := uc.customerRepo.Create(customer); err != nil {
		return nil, err
	}

	link := &UserCustomer{
		UserID:     params.UserID,
		CustomerID: customer.ID,
		Role:       MembershipOwner,
	}

	if err := uc.userCustomerRepo.Link(link); err != nil {
		return nil, err
	}

	return customer, nil
}
