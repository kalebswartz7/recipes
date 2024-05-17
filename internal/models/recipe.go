package models

import "time"

type Recipe struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Difficulty  int       `json:"difficulty"`
	Minutes     int       `json:"minutes"`
	CreatedAt   time.Time `json:"created_at"`
}
