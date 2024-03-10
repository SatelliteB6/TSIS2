package model

import (
	"database/sql"
	"log"
)

type Champion struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Price int    `json:"price"`
}

func CreateChampion(db *sql.DB, champion Champion) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO champions(name, class, price) VALUES($1, $2, $3) RETURNING id",
		champion.Name, champion.Class, champion.Price).Scan(&id)

	if err != nil {
		log.Println("Error creating champion:", err)
		return 0, err
	}

	return id, nil
}

func GetChampion(db *sql.DB, championID int) (*Champion, error) {
	var champion Champion
	err := db.QueryRow("SELECT id, name, class, price FROM champions WHERE id = $1", championID).
		Scan(&champion.ID, &champion.Name, &champion.Class, &champion.Price)

	if err != nil {
		log.Println("Error getting champion:", err)
		return nil, err
	}

	return &champion, nil
}

func ListChampions(db *sql.DB) ([]Champion, error) {
	rows, err := db.Query("SELECT id, name, class, price FROM champions")
	if err != nil {
		log.Println("Error listing champions:", err)
		return nil, err
	}
	defer rows.Close()

	var champions []Champion
	for rows.Next() {
		var champion Champion
		err := rows.Scan(&champion.ID, &champion.Name, &champion.Class, &champion.Price)
		if err != nil {
			log.Println("Error scanning champion:", err)
			return nil, err
		}
		champions = append(champions, champion)
	}

	return champions, nil
}

func UpdateChampion(db *sql.DB, champion Champion) error {
	_, err := db.Exec("UPDATE champions SET name=$1, class=$2, price=$3 WHERE id=$4",
		champion.Name, champion.Class, champion.Price, champion.ID)

	if err != nil {
		log.Println("Error updating champion:", err)
		return err
	}

	return nil
}

func DeleteChampion(db *sql.DB, championID int) error {
	_, err := db.Exec("DELETE FROM champions WHERE id=$1", championID)
	if err != nil {
		log.Println("Error deleting champion:", err)
		return err
	}

	return nil
}
