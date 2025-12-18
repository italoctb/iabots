package entities

import (
	"time"

	"github.com/google/uuid"
)

type CreditTransactionType string

const (
	Purchase   CreditTransactionType = "purchase"
	Debit      CreditTransactionType = "debit"
	Bonus      CreditTransactionType = "bonus"
	Withdrawal CreditTransactionType = "withdrawal"
)

type CreditTransaction struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	CustomerID string
	AmountBits float64
	Amount     float64 // valor real (em R$)
	Currency   string  // Ex: "BRL"
	Type       CreditTransactionType
	CreatedAt  time.Time
}
