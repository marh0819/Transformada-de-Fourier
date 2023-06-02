package main

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/transformada", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			getCountries(w, r)
			/*
				case http.MethodPost:
					addCountries(w, r)
			*/
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed")
			return
		}
	})
}
