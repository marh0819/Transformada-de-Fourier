package main

import (
	"math"
)

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

func fase(fft ...complex128) []float64 {
	n := len(fft)
	fase := make([]float64, n)

	for i := 0; i < n; i++ {
		division := imag(fft[i]) / real(fft[i])
		arctan := arctan(division)
		fase[i] = arctan

	}

	return fase
}

func fase2(fft ...complex128) []float64 {
	n := len(fft)
	fase := make([]float64, n)

	for i := 0; i < n; i++ {
		division := imag(fft[i]) / real(fft[i])
		arctan := math.Atan(division)
		fase[i] = arctan

	}

	return fase
}
