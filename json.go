package main

import (
	"encoding/json"
	"fmt"
)

// MagYFase es una estructura que representa una magnitud y una fase.
// Tiene dos campos, Mag y Fase, que representan la magnitud y la fase respectivamente.
// Estos campos se utilizan para almacenar información sobre los valores de una transformada de Fourier.
type MagYFase struct {
	Mag  float64 `json:"mag"`  // Campo que representa la magnitud.
	Fase float64 `json:"fase"` // Campo que representa la fase.
}

// JSON es una función que convierte una serie de valores de una transformada de Fourier a formato JSON.
// Recibe una serie de valores complejos (fft) y devuelve una cadena de texto en formato JSON que representa los valores de magnitud y fase.
// La función calcula las magnitudes y fases de los valores complejos utilizando las funciones magnitud() y fase2().
// Luego, crea una estructura MagYFase para cada par de magnitud y fase, y los agrega a un slice de estructuras.
// Finalmente, utiliza el paquete json para codificar el slice de estructuras a formato JSON y lo devuelve como una cadena de texto.
// Si ocurre un error durante la codificación a JSON, se imprime un mensaje de error y se devuelve una cadena vacía.
func JSON(fft ...complex128) string {
	// Calcula las magnitudes y fases de los valores complejos.
	magnitudes := magnitud(fft...)
	fases := fase(fft...)

	// Crea un slice para almacenar las estructuras MagYFase.
	data := make([]MagYFase, 0)

	// Crea una estructura MagYFase para cada par de magnitud y fase y las agrega al slice.
	for i := 0; i < len(magnitudes); i++ {
		item := MagYFase{
			Mag:  magnitudes[i],
			Fase: fases[i],
		}

		data = append(data, item)
	}

	// Codifica el slice de estructuras a formato JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return ""
	}

	// Imprime el JSON generado (opcional).
	fmt.Println("Json: ", string(jsonData))

	// Devuelve el JSON como una cadena de texto.
	return string(jsonData)
}
