package main

import (
	"fmt"
	"math"
)

const pi = 3.14159265358979323846264338327950288419716939937510582097494459

func llenarSlice(largo int) []int {

	entradaA := []int{}
	for i := 0; i < largo; i++ {
		entradaA = append(entradaA, i)
	}
	return entradaA

}

func imprimeSlice(entrada ...int) {
	for i := 0; i < 0; i++ {
		fmt.Print(entrada[i])
	}

}

func recursivaFourier(entrada ...int) []int {
	if len(entrada) == 1 {
		return entrada
	} else {

		slicePar := []int{}
		sliceImpar := []int{}
		sliceSalida := []int{}
		for i := 0; i < len(entrada); i++ {
			if i%2 == 0 {
				slicePar = append(slicePar, entrada[i])

			} else {
				sliceImpar = append(sliceImpar, entrada[i])

			}
		}

		slicePar = recursivaFourier(slicePar...)
		sliceImpar = recursivaFourier(sliceImpar...)

		sliceSalida = slicePar

		for j := 0; j < len(sliceImpar); j++ {
			sliceSalida = append(sliceSalida, sliceImpar[j])
		}

		return sliceSalida
	}
}

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
	if n > 1 {

		par := make([]complex128, n/2)
		par = traerPar(arre...)

		impar := make([]complex128, n/2)
		impar = traerImpar(arre...)

		par = transformadaFormula(par...)

		impar = transformadaFormula(impar...)

		w := complex(0, 0)
		z := complex(0, 0)

		for m := 0; m < n/2; m++ {
			aux := float64(m) / float64(n)
			w = complex(float64(math.Cos(2*math.Pi*float64(aux))), imag(w))
			w = complex(real(w), float64(-1*math.Sin(2*math.Pi*float64(aux))))
			z = complex(real(w)*real(impar[m])-imag(w)*imag(impar[m]), real(w)*imag(impar[m])+imag(w)*real(impar[m]))
			arre[m] = complex(real(par[m])+real(z), imag(arre[m]))
			arre[m] = complex(real(arre[m]), imag(par[m])+imag(z))

			arre[m+(n>>1)] = complex(real(par[m])-real(z), imag(arre[m+n/2]))
			arre[m+(n>>1)] = complex(real(arre[m+n/2]), imag(par[m])-imag(z))

		}
	}

	return arre
}

func traerPar(arre ...complex128) []complex128 {
	n := len(arre) / 2
	p := 0
	pares := make([]complex128, n)

	for i := 0; i < len(arre); i++ {
		if i%2 == 0 {
			pares[p] = arre[i]
			p++
		}
	}

	return pares
}

func traerImpar(arre ...complex128) []complex128 {
	n := len(arre) / 2
	p := 0
	impares := make([]complex128, n)

	for i := 0; i < len(arre); i++ {
		if i%2 != 0 {
			impares[p] = arre[i]
			p++
		}
	}

	return impares
}

func potenciaDeDos(n int) bool {
	if n <= 0 {
		return false
	}

	log2 := math.Log2(float64(n))
	return math.Floor(log2) == log2
}

func main() {

	a := 8
	//var a int
	//fmt.Print("Ingrese un número entero: ")
	//fmt.Scan(&a)
	if potenciaDeDos(a) {
		fmt.Println("El número ingresado es una potencia de 2.")
		arreglo := []complex128{}
		arreglo2 := []complex128{}
		arreglo3 := []complex128{}

		for k := 0; k < a; k++ {
			aux := complex(float64(k), 0)
			arreglo = append(arreglo, aux)
		}

		arreglo2 = recursivaFourier2(arreglo...)
		fmt.Println(arreglo2, "Arreglo ")
		arreglo3 = transformadaFormula(arreglo...)
		fmt.Println("\n\n\n", arreglo3)

	} else {
		fmt.Println("El número ingresado NO es una potencia de 2.")
	}

}
