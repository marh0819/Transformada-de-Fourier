package main

import (
	"encoding/json"
	"fmt"
)

type MagYFase struct {
	Mag  float64 `json:"mag"`
	Fase float64 `json:"fase"`
}

// Creamos Json para la magnitud y la fase.
func JSON(fft ...complex128) {

	magnitudes := magnitud(fft...)
	fases := fase2(fft...)

	data := make([]MagYFase, 0)

	for i := 0; i < len(magnitudes); i++ {
		item := MagYFase{
			Mag:  magnitudes[i],
			Fase: fases[i],
		}

		data = append(data, item)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println("Json: ", string(jsonData))
}
