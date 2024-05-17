package models

import "time"

type Quantity struct {
	ID        int       `json:"id"`
	Category  string    `json:"category"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
