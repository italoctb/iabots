package entities

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserCustomer struct {
	ID         uuid.UUID          `gorm:"primaryKey"`
	UserID     uuid.UUID          `gorm:"type:uuid;not null;index"`
	CustomerID uuid.UUID          `gorm:"type:uuid;not null;index"`
	Role       MembershipRoleType `gorm:"type:varchar(30);not null;default:'owner'"`
	CreatedAt  time.Time
}

type MembershipRoleType string

const (
	MembershipOwner      MembershipRoleType = "owner"
	MembershipConsultant MembershipRoleType = "consultant"
	MembershipMember     MembershipRoleType = "member"
)

func (r MembershipRoleType) IsValid() bool {
	switch r {
	case MembershipOwner, MembershipConsultant, MembershipMember:
		return true
	default:
		return false
	}
}

func (r MembershipRoleType) String() string {
	return string(r)
}

func (r MembershipRoleType) MarshalJSON() ([]byte, error) {
	if !r.IsValid() {
		return nil, fmt.Errorf("invalid MembershipRoleType: %s", r)
	}
	return json.Marshal(string(r))
}

func (r *MembershipRoleType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	converted := MembershipRoleType(str)
	if !converted.IsValid() {
		return fmt.Errorf("invalid MembershipRoleType: %s", str)
	}
	*r = converted
	return nil
}

func AllMembershipRoleTypes() []MembershipRoleType {
	return []MembershipRoleType{MembershipOwner, MembershipConsultant, MembershipMember}
}
