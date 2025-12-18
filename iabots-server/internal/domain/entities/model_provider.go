package entities

import "github.com/google/uuid"

type ModelProvider struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	Name            string    // ex: "OpenAI", "Cohere", "Mistral", "Google"
	ModelName       string    // ex: "gpt-4", "command-r", "gemini-pro"
	TokenCostPrompt float64   // preço por 1k tokens enviados (prompt)
	TokenCostAnswer float64   // preço por 1k tokens de resposta (completion)
	Active          bool
}
