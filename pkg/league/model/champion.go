package model

type Champion struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Difficulty int    `json:"difficulty"`
}
