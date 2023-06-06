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

func servidor() {
	mux := http.NewServeMux()

	// Manejador para /grafica que sirve el archivo index.html
	mux.HandleFunc("/grafica", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./consumo/index.html")
	})

	// Manejador para /consum.js que sirve el archivo consum.js
	mux.HandleFunc("/consum.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./consumo/consum.js")
	})

	// Manejador para /style.css que sirve el archivo style.css
	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./consumo/style.css")
	})

	// Manejador para /transformada que ejecuta la función
	mux.HandleFunc("/transformada", fft)

	// Manejador para la raíz que redirige a /grafica
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/grafica", http.StatusTemporaryRedirect)
	})

	// Crea el servidor HTTP y asigna el enrutador (mux) a él
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Inicia el servidor
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func main() {
	servidor()
}
