package models

import "time"

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// protein, carb, fruit, vegetable etc
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
