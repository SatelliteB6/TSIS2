package league

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connStr := "user=postgres password=Satellite_b6 dbname=crud sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database")
}

type Champion struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Price int    `json:"price"`
}

func ListChampions(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, _ := strconv.Atoi(pageSizeStr)
	filter := r.URL.Query().Get("filter")
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")

	if page == 0 {
		page = 1
	}

	if sortBy == "" {
		sortBy = "id"
	}
	if sortOrder != "desc" {
		sortOrder = "asc"
	}

	query := "SELECT id, name, class, price FROM champions"
	if filter != "" {
		_, err := strconv.Atoi(filter)
		if err == nil {
			query += " WHERE price = " + filter
		} else {
			query += " WHERE name LIKE '%" + filter + "%' OR class LIKE '%" + filter + "%'"
		}
	}
	query += " ORDER BY " + sortBy + " " + sortOrder
	if pageSize > 0 {
		query += " LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa((page-1)*pageSize)
	}

	rows, err := db.Query(query)
	if err != nil {
		handleError(w, err)
		return
	}
	defer rows.Close()

	var champions []Champion
	for rows.Next() {
		var champion Champion
		err := rows.Scan(&champion.ID, &champion.Name, &champion.Class, &champion.Price)
		if err != nil {
			handleError(w, err)
			return
		}
		champions = append(champions, champion)
	}

	writeJSONResponse(w, champions)
}

func CreateChampion(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newChampion Champion
	err := json.NewDecoder(r.Body).Decode(&newChampion)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = db.QueryRow("INSERT INTO champions (name, class, price) VALUES ($1, $2, $3) RETURNING id",
		newChampion.Name, newChampion.Class, newChampion.Price).
		Scan(&newChampion.ID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newChampion)
}

func GetChampion(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Champion ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid champion ID", http.StatusBadRequest)
		return
	}

	var champion Champion
	err = db.QueryRow("SELECT id, name, class, price FROM champions WHERE id = $1", id).
		Scan(&champion.ID, &champion.Name, &champion.Class, &champion.Price)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Error retrieving champion", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, champion)
}

func UpdateChampion(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Champion ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid champion ID", http.StatusBadRequest)
		return
	}

	var updatedChampion Champion
	err = json.NewDecoder(r.Body).Decode(&updatedChampion)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var existingID int
	err = db.QueryRow("SELECT id FROM champions WHERE id = $1", id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Champion not found", http.StatusNotFound)
			return
		} else {
			log.Printf("Error querying database: %v", err)
			http.Error(w, "Error checking champion existence", http.StatusInternalServerError)
			return
		}
	}

	_, err = db.Exec("UPDATE champions SET name = $1, class = $2, price = $3 WHERE id = $4",
		updatedChampion.Name, updatedChampion.Class, updatedChampion.Price, id)
	if err != nil {
		log.Printf("Error updating champion in the database: %v", err)
		http.Error(w, "Error updating champion", http.StatusInternalServerError)
		return
	}

	successMessage := map[string]string{"message": "Champion updated successfully"}
	writeJSONResponse(w, successMessage)
}

func DeleteChampion(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Champion ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid champion ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM champions WHERE id = $1", id)
	if err != nil {
		handleError(w, err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Champion with specified ID not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Champion deleted successfully")
}

func handleError(w http.ResponseWriter, err error) {
	fmt.Println("Error:", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
