package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		switch r.Method {
		case http.MethodGet:
			league.ListChampions(w, r, db)
		case http.MethodPost:
			league.CreateChampion(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/champions/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/champions/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid champion ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			r.URL.RawQuery = "id=" + strconv.Itoa(id)
			league.GetChampion(w, r, db)
		case http.MethodPut:
			r.URL.RawQuery = "id=" + strconv.Itoa(id)
			league.UpdateChampion(w, r, db)
		case http.MethodDelete:
			r.URL.RawQuery = "id=" + strconv.Itoa(id)
			league.DeleteChampion(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := "8080"
	log.Printf("Server running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
