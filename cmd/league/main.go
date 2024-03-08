package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/champions", ListChampions)
	http.HandleFunc("/champions/create", CreateChampion)
	http.HandleFunc("/champions/get", GetChampion)
	http.HandleFunc("/champions/update", UpdateChampion)
	http.HandleFunc("/champions/delete", DeleteChampion)

	port := ":8080"
	log.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}