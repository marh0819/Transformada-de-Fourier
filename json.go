package main

import (
	"encoding/json"
	"fmt"
)

type MagYFase struct {
	Mag  float64 `json:"mag"`
	Fase float64 `json:"fase"`
}

/*
// Creamos Json para la magnitud y la fase.
func JSON(fft ...complex128) []byte {

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
		return nil

	}

	fmt.Println("Json: ", string(jsonData))
	return jsonData

}*/

// Para poder revisar que los valores que queremos si son los correctos en el resultado del json debemos usar esta funcion
func JSON(fft ...complex128) string {

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
		return ""
	}

	fmt.Println("Json: ", string(jsonData))
	return string(jsonData)

}
