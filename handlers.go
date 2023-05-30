package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

}

func getCountries(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
	fmt.Fprintf(w, "%v", countries)
}

func addCountries(w http.ResponseWriter, r *http.Request) {

	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "%v", err)
		return
	}
	countries = append(countries, country)
	fmt.Fprintf(w, "Country was added")
}
