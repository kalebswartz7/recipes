package models

import "time"

type Step struct {
	ID		int		`json:"id"`
	Sequence	int		`json:"sequence"`
	Category	string		`json:"category"`
	Description	string		`json:"description"`
	Minutes		int		`json:"minutes"`
	CreatedAt	time.Time	`json:"created_at"`

	// foreign keys
	RecipeId	int		`json:"recipe_id"`
}
