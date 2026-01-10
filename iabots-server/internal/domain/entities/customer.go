package entities

import "github.com/google/uuid"

type Customer struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	CompanyName string
	Whatsapp    string `gorm:"unique"`
}

func NewCustomer(companyName string, whatsapp string) *Customer {
	return &Customer{
		ID:          uuid.New(),
		CompanyName: companyName,
		Whatsapp:    whatsapp,
	}
}
