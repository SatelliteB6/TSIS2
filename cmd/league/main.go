package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SatelliteB6/TSIS2/pkg/league"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "Satellite_b6"
	dbname   = "crud"
)

func initDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
	return db
}

func main() {
	db := initDB()
	defer db.Close()

	http.HandleFunc("/champions", func(w http.ResponseWriter, r *http.Request) {
		league.ListChampions(w, r, db)
	})
	http.HandleFunc("/champions/create", func(w http.ResponseWriter, r *http.Request) {
		league.CreateChampion(w, r, db)
	})
	http.HandleFunc("/champions/get", func(w http.ResponseWriter, r *http.Request) {
		league.GetChampion(w, r, db)
	})
	http.HandleFunc("/champions/update", func(w http.ResponseWriter, r *http.Request) {
		league.UpdateChampion(w, r, db)
	})
	http.HandleFunc("/champions/delete", func(w http.ResponseWriter, r *http.Request) {
		league.DeleteChampion(w, r, db)
	})

	port := "8080"
	log.Printf("Server running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
