package entities

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type AssistantBot struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	CustomerID      uuid.UUID `gorm:"type:uuid;not null;index"`
	ModelProviderID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name            string
	ContextMessage  string
	MaxTokens       int
	FreezeTime      int
	Status          AssistantStatus
	OnlineStartTime *string // formato: "15:04" (ex: 08:30)
	OnlineEndTime   *string // formato: "15:04" (ex: 18:45)
}

type AssistantStatus string

const (
	AssistantActive  AssistantStatus = "active"
	AssistantPaused  AssistantStatus = "paused"
	AssistantOffline AssistantStatus = "offline"
	AssistantAuto    AssistantStatus = "auto" // avalia horário online
)

func (s AssistantStatus) IsValid() bool {
	switch s {
	case AssistantActive, AssistantPaused, AssistantOffline, AssistantAuto:
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
	return []AssistantStatus{AssistantActive, AssistantPaused, AssistantAuto, AssistantOffline}
}

func AllBotStatusesToString() string {
	// retorna todos os status como uma string separados por vírgula
	statuses := AllBotStatuses()
	strs := make([]string, len(statuses))
	for i, s := range statuses {
		strs[i] = s.String()
	}
	return fmt.Sprintf("%s", strs)

}
