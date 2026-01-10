package entities

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ConsultantWithdrawalRequest struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	ConsultantID string
	AmountBits   float64
	Amount       float64 // valor em moeda local
	Currency     string  // Ex: "BRL"
	Status       WithdrawalStatus
	PixKey       string
	RequestedAt  time.Time
	CompletedAt  *time.Time
}

type WithdrawalStatus string

const (
	WithdrawalPending  WithdrawalStatus = "pending"
	WithdrawalApproved WithdrawalStatus = "approved"
	WithdrawalRejected WithdrawalStatus = "rejected"
)

func (s WithdrawalStatus) IsValid() bool {
	switch s {
	case WithdrawalPending, WithdrawalApproved, WithdrawalRejected:
		return true
	default:
		return false
	}
}

func (s WithdrawalStatus) String() string {
	return string(s)
}

func (s WithdrawalStatus) MarshalJSON() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("invalid WithdrawalStatus: %s", s)
	}
	return json.Marshal(string(s))
}

func (s *WithdrawalStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	converted := WithdrawalStatus(str)
	if !converted.IsValid() {
		return fmt.Errorf("invalid WithdrawalStatus: %s", str)
	}
	*s = converted
	return nil
}

func AllWithdrawalStatuses() []WithdrawalStatus {
	return []WithdrawalStatus{
		WithdrawalPending,
		WithdrawalApproved,
		WithdrawalRejected,
	}
}
