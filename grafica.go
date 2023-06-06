package main

import (
	"fmt"
	"net/http"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("consumo")).ServeHTTP(w, r)
}

func grafica() {
	mux := http.NewServeMux()

	// Configura el manejador para servir archivos estáticos
	mux.Handle("/", http.FileServer(http.Dir("./consumo")))

	// Especifica el puerto en el que deseas ejecutar el servidor
	port := "8080"

	fmt.Println("Servidor web escuchando en el puerto:", port)

	// Inicia el servidor web en el puerto especificado
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
