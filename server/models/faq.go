package models

type Faq struct {
	ID         uint
	CustomerId int
	Question   string
	Answer     string
	Embedding  []uint8
}
