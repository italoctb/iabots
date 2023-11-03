package models

type IabotsPayload struct {
	CustomerID uint          `json:"customer_id"`
	Messages   []RoleMessage `json:"messages"`
}

type RoleMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
