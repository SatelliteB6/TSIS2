package model

type Match struct {
	ID           int    `json:"id"`
	Date         string `json:"date"`
	Duration     string `json:"duration"`
	WinnerTeamID int    `json:"winner_team_id"`
}
