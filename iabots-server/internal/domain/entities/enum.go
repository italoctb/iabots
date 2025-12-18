package entities

import (
	"encoding/json"
	"fmt"
)

// --------------------
// AssistantStatus Enum
// --------------------
type AssistantStatus string

const (
	AssistantActive   AssistantStatus = "active"
	AssistantPaused   AssistantStatus = "paused"
	AssistantDisabled AssistantStatus = "disabled"
)

func (s AssistantStatus) IsValid() bool {
	switch s {
	case AssistantActive, AssistantPaused, AssistantDisabled:
		return true
	default:
		return false
	}
}

func (s AssistantStatus) String() string {
	return string(s)
}

func (s AssistantStatus) MarshalJSON() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("invalid BotStatus: %s", s)
	}
	return json.Marshal(string(s))
}

func (s *AssistantStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	converted := AssistantStatus(str)
	if !converted.IsValid() {
		return fmt.Errorf("invalid BotStatus: %s", str)
	}
	*s = converted
	return nil
}

func AllBotStatuses() []AssistantStatus {
	return []AssistantStatus{AssistantActive, AssistantPaused, AssistantDisabled}
}

// --------------------
// WithdrawalStatus Enum
// --------------------
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

// --------------------
// UserRole Enum
// --------------------
type UserRole string

const (
	RoleAdmin      UserRole = "admin"
	RoleConsultant UserRole = "consultor"
	RoleClient     UserRole = "cliente"
)

func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleConsultant, RoleClient:
		return true
	default:
		return false
	}
}

func (r UserRole) String() string {
	return string(r)
}

func (r UserRole) MarshalJSON() ([]byte, error) {
	if !r.IsValid() {
		return nil, fmt.Errorf("invalid UserRole: %s", r)
	}
	return json.Marshal(string(r))
}

func (r *UserRole) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	converted := UserRole(str)
	if !converted.IsValid() {
		return fmt.Errorf("invalid UserRole: %s", str)
	}
	*r = converted
	return nil
}

func AllUserRoles() []UserRole {
	return []UserRole{RoleAdmin, RoleConsultant, RoleClient}
}
