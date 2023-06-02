package main

import (
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

/*
func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Llamar a la función JSON y pasar el resultado directamente a json.NewEncoder
	mag, fase := JSON()
	json.NewEncoder(w).Encode(mag)
	json.NewEncoder(w).Encode(fase)
	fmt.Fprintf(w, "%v", mag)
	fmt.Fprintf(w, "%v", fase)
}*/

func getCountries(w http.ResponseWriter, r *http.Request) {
	cantidadDatos := 8
	w.Header().Set("Content-Type", "application/json")
	if potenciaDeDos(cantidadDatos) {
		fmt.Println("El número ingresado es una potencia de 2.")
		slice := []complex128{}
		sliceOrdenado := []complex128{}
		fft := []complex128{}

		for k := 0; k < cantidadDatos; k++ {
			aux := complex(float64(k), 0)
			slice = append(slice, aux)
		}

		sliceOrdenado = ordenarFourier(slice...)
		fmt.Println(sliceOrdenado, " slice ordenado")
		fft = transformadaFormula(slice...)
		fmt.Println("\n\n\nDespues de aplicar la fft:\n", fft, "\n\n")

		data := JSON(fft...)

		//json.NewEncoder(w).Encode(data)
		fmt.Fprintf(w, "%v", data)

	} else {
		fmt.Println("El número ingresado NO es una potencia de 2.")

	}

}

/*
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
}*/
