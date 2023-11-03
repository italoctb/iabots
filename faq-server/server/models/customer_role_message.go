package models

type CustomerRoleMessage struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	CustomerID uint   `json:"customer_id"`
	Role       string `json:"role"`
	Message    string `json:"message"`
}
