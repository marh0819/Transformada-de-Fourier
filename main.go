package main

import (
	"fmt"
	"math"
)

func recursivaFourier2(entrada ...complex128) []complex128 {
	if len(entrada) == 1 {
		return entrada
	} else {

		slicePar := []complex128{}
		sliceImpar := []complex128{}
		sliceSalida := []complex128{}
		for i := 0; i < len(entrada); i++ {
			if i%2 == 0 {

				slicePar = append(slicePar, entrada[i])

			} else {

				sliceImpar = append(sliceImpar, entrada[i])

			}
		}

		slicePar = recursivaFourier2(slicePar...)
		sliceImpar = recursivaFourier2(sliceImpar...)

		sliceSalida = slicePar

		for j := 0; j < len(sliceImpar); j++ {
			sliceSalida = append(sliceSalida, sliceImpar[j])
		}

		return sliceSalida
	}
}

func transformadaFormula(arre ...complex128) []complex128 {
	n := len(arre)
	w := complex(0, 0)
	z := complex(0, 0)
	salidaA := make([]complex128, n)

	for m := 0; m < len(arre)/2; m += 2 {
		aux := m / n
		w = complex(float64(math.Cos(2*math.Pi*float64(aux))), imag(w))
		w = complex(real(w), float64(-1*math.Sin(2*math.Pi*float64(aux))))
		z = complex(real(w)*real(arre[m+1])-imag(w)*imag(arre[m+1]), imag(z))
		z = complex(real(z), real(w)*imag(arre[m+1])+imag(w)*real(arre[m+1]))

		salidaA[m] = complex(real(arre[m])+real(z), imag(salidaA[m]))
		salidaA[m] = complex(real(salidaA[m]), imag(arre[m])+imag(z))

		fmt.Println("", real(arre[m])-real(z))
		salidaA[m+(n/2)] = complex(real(arre[m])-real(z), imag(salidaA[m+n/2]))
		salidaA[m+(n/2)] = complex(real(salidaA[m+n/2]), imag(arre[m])-imag(z))

	}
	return salidaA
}

func main() {

	a := 8 // aqui van los multiplos de llenado malparido miguel
	arreglo := []complex128{}
	arreglo2 := []complex128{}
	arreglo3 := []complex128{}

	//llenado del slice principal (sin modificaciones)
	for k := 0; k < a; k++ {
		aux := complex(float64(k), 0)
		arreglo = append(arreglo, aux)
	}

	//operaciones del slice 
    //ordenamiento recursivo
	arreglo2 = recursivaFourier2(arreglo...)
	fmt.Println(arreglo2, "Arreglo ")


    //aplicacion de la formulal de la transformada
	arreglo3 = transformadaFormula(arreglo2[0], arreglo2[1],)
	fmt.Println(arreglo3)
}
