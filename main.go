package main

import (
	"fmt"
	"net/http"
)

// index es una función que maneja las solicitudes al punto de entrada "/" de un servidor HTTP.
// Recibe un objeto ResponseWriter (w) y un objeto Request (r) como argumentos.
// La función verifica el método de la solicitud y responde de acuerdo a ello.
func index(w http.ResponseWriter, r *http.Request) {
	// Verifica si el método de la solicitud no es GET.
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

	// Si el método de la solicitud es GET, se puede agregar aquí el código para manejar la respuesta.
	// Este código se ejecutará cuando se realice una solicitud GET a "/".
}

// servidor es una función que configura y ejecuta un servidor HTTP.
// Utiliza el paquete http para definir los manejadores de las diferentes rutas del servidor.
// El servidor responde a las solicitudes HTTP con los archivos index.html, consum.js y style.css, y ejecuta la función fft en la ruta /transformada.
// Escucha en el puerto 8080 y utiliza el enrutador (mux) para manejar las solicitudes entrantes.
// La función inicia el servidor y entra en un bucle infinito para atender las solicitudes entrantes.
// Si ocurre un error al iniciar el servidor, se produce un pánico.
func servidor() {
	// Crea un nuevo enrutador HTTP.
	mux := http.NewServeMux()

	// Manejador para /grafica que sirve el archivo index.html.
	mux.HandleFunc("/grafica", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./consumo/index.html")
	})

	// Manejador para /consum.js que sirve el archivo consum.js.
	mux.HandleFunc("/consum.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./consumo/consum.js")
	})

	// Manejador para /style.css que sirve el archivo style.css.
	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./consumo/style.css")
	})

	// Manejador para /transformada que ejecuta la función fft.
	mux.HandleFunc("/transformada", fft)

	// Manejador para la raíz que redirige a /grafica.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/grafica", http.StatusTemporaryRedirect)
	})

	// Crea un nuevo servidor HTTP con la configuración definida.
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Inicia el servidor y entra en un bucle infinito para atender las solicitudes entrantes.
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// main es la función de entrada del programa.
// Llama a la función servidor para iniciar el servidor HTTP.
func main() {
	servidor()
}
