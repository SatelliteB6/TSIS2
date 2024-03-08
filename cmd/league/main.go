// cmd/league/main.go

package main

import (
	"log"
	"net/http"
	"github.com/SatelliteB6/TSIS2/pkg/league"
)

func main() {
	http.HandleFunc("/champions", league.ListChampions)
	http.HandleFunc("/champions/create", league.CreateChampion)
	http.HandleFunc("/champions/get", league.GetChampion)
	http.HandleFunc("/champions/update", league.UpdateChampion)
	http.HandleFunc("/champions/delete", league.DeleteChampion)

	// Start the HTTP server
	port := "8080"
	log.Printf("Server running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
