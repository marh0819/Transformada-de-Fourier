package main

import (
	"math"
)

// magnitud calcula las magnitudes de los valores complejos en la serie fft.
// Recibe una serie de valores complejos (fft) y devuelve una serie de valores de magnitud correspondientes.
// La función itera sobre los valores complejos en fft y calcula la magnitud utilizando la fórmula |z| = sqrt(Re(z)^2 + Im(z)^2).
// Los resultados se almacenan en un slice de valores de magnitud que se devuelve al final.
func magnitud(fft ...complex128) []float64 {
	n := len(fft)
	mag := make([]float64, n)

	for i := 0; i < n; i++ {
		sumaCuadrados := elevar(real(fft[i]), 2) + elevar(imag(fft[i]), 2)
		raiz := math.Sqrt(sumaCuadrados)
		mag[i] = raiz
	}

	return mag
}

// fase calcula las fases de los valores complejos en la serie fft.
// Recibe una serie de valores complejos (fft) y devuelve una serie de valores de fase correspondientes.
// La función itera sobre los valores complejos en fft y calcula la fase utilizando la fórmula arctan(Im(z) / Re(z)).
// Los resultados se almacenan en un slice de valores de fase que se devuelve al final.
func fase(fft ...complex128) []float64 {
	n := len(fft)
	fase := make([]float64, n)

	for i := 0; i < n; i++ {
		division := imag(fft[i]) / real(fft[i])
		arctan := math.Atan(division)
		fase[i] = arctan
	}

	return fase
}
