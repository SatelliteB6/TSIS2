package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Champion struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Price int    `json:"price"`
}

var champions = []Champion{
	{ID: 1, Name: "Aatrox", Class: "Fighter", Price: 4800},
	{ID: 2, Name: "Ahri", Class: "Mage", Price: 3150},
	{ID: 3, Name: "Akali", Class: "Assassin", Price: 3150},
	{ID: 4, Name: "Alistar", Class: "Tank", Price: 1350},
	{ID: 5, Name: "Amumu", Class: "Tank", Price: 450},
}

func ListChampions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(champions)
}

func CreateChampion(w http.ResponseWriter, r *http.Request) {
	var newChampion Champion
	err := json.NewDecoder(r.Body).Decode(&newChampion)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newChampion.ID = len(champions) + 1

	champions = append(champions, newChampion)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newChampion)
}

func GetChampion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid champion ID", http.StatusBadRequest)
		return
	}

	for _, champion := range champions {
		if champion.ID == id {
			json.NewEncoder(w).Encode(champion)
			return
		}
	}

	http.Error(w, "Champion not found", http.StatusNotFound)
}

func UpdateChampion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
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

	for i, champion := range champions {
		if champion.ID == id {
			champions[i] = updatedChampion
			json.NewEncoder(w).Encode(updatedChampion)
			return
		}
	}

	http.Error(w, "Champion not found", http.StatusNotFound)
}

func DeleteChampion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid champion ID", http.StatusBadRequest)
		return
	}

	for i, champion := range champions {
		if champion.ID == id {
			champions = append(champions[:i], champions[i+1:]...)
			w.Write([]byte(`{"message": "Champion deleted"}`))
			return
		}
	}

	http.Error(w, "Champion not found", http.StatusNotFound)
}
