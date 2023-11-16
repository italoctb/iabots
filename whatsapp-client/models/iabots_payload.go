package models

type IabotsPayload struct {
	CustomerID int           `json:"customer_id"`
	Messages   []RoleMessage `json:"messages"`
}

type RoleMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
