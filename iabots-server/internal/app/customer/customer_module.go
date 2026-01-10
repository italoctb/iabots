package customer

import (
	. "iabots-server/internal/delivery/http/handlers"
	. "iabots-server/internal/domain/usecases/customer"
	. "iabots-server/internal/infra/database"
	. "iabots-server/internal/infra/database/repositories"
)

type CustomerModule struct {
	Handler *CustomerHandler
}

func NewCustomerModule(db *Database) *CustomerModule {
	/// Injection of binds - Customer Module
	/// --------------------------------------------------
	///  Repositories
	/// --------------------------------------------------
	customerRepo := NewCustomerGormRepository(db.DB)
	userCustomerRepo := NewUserCustomerGormRepository(db.DB)

	/// --------------------------------------------------
	///  Use Cases
	/// --------------------------------------------------
	createCustomerUseCase := NewCreateCustomerUseCase(customerRepo, userCustomerRepo)
	/// --------------------------------------------------
	///  Handlers
	/// --------------------------------------------------
	handler := NewCustomerHandler(createCustomerUseCase)

	return &CustomerModule{
		Handler: handler,
	}
}
