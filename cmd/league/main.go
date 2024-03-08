package main

import (
	"net/http"
	"github.com/SatelliteB6/TSIS2/pkg/league"
)

func main() {
	// Define HTTP routes
	http.HandleFunc("/champions", league.ListChampions)
	http.HandleFunc("/champions/create", league.CreateChampion)
	http.HandleFunc("/champions/get", league.GetChampion)
	http.HandleFunc("/champions/update", league.UpdateChampion)
	http.HandleFunc("/champions/delete", league.DeleteChampion)

	port := "8080"
	http.ListenAndServe(":"+port, nil)
}
