package model

type Champion struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Price int    `json:"price"`
}
