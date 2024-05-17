package models

import "time"

type StepIngredient struct {
	StepId       int       `json:"step_id"`
	IngredientId int       `json:"ingredient_id"`
	QuantityId   int       `json:"quantity_id"`
	CreatedAt    time.Time `json:"created_at"`
}
