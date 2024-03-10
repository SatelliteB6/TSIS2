package model

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Players []int  `json:"players"`
}
