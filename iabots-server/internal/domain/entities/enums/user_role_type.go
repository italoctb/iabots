package enums

import (
	"encoding/json"
	"fmt"
)

// --------------------
// UserRole Enum
// --------------------
type UserRoleType string

const (
	RoleAdmin    UserRoleType = "admin"
	RoleCostumer UserRoleType = "cliente"
)

func (r UserRoleType) IsValid() bool {
	switch r {
	case RoleAdmin, RoleCostumer:
		return true
	default:
		return false
	}
}

func (r UserRoleType) String() string {
	return string(r)
}

func (r UserRoleType) MarshalJSON() ([]byte, error) {
	if !r.IsValid() {
		return nil, fmt.Errorf("invalid UserRole: %s", r)
	}
	return json.Marshal(string(r))
}

func (r *UserRoleType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	converted := UserRoleType(str)
	if !converted.IsValid() {
		return fmt.Errorf("invalid UserRole: %s", str)
	}
	*r = converted
	return nil
}

func AllUserRoles() []UserRoleType {
	return []UserRoleType{RoleAdmin, RoleCostumer}
}
