package main

import (
	"fmt"
	"math"
)

func comprobacion(lar int) bool {
	x := false
	switch lar {

	case 1:
		x = true
		break
	case 2:
		x = true
		break
	case 4:
		x = true
		break
	case 8:
		x = true
		break
	case 16:
		x = true
		break
	case 32:
		x = true
		break
	case 64:
		x = true
		break
	case 128:
		x = true
		break
	case 256:
		x = true
		break
	case 512:
		x = true
		break
	default:
		x = false
		break
	}

	return x
}

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
	//var entradaA = []int {}
	//entradaA := make([]int , 0 )

	//entradaA := []int{}                           // slice de entrada
	///salidaA := []int{}                            // slice de salida
	//ingreso := 7                                 //valor de ingreo (solo prueba)

	//if comprobacion(ingreso){
	//entradaA = llenarSlice(ingreso)               // lena el slice con el valor de ingreso
	//imprimeSlice(entradaA...)                 // imprime los valores del slice
	//	salidaA = recursivaFourier(entradaA...)   //llama a la recuriva de furier ordenadora

	//	fmt.Println(entradaA)
	//	fmt.Println(salidaA)
	//}else {

	//println("existe un error, este algoritmo solo puede ser ingresado con potencias de 2")
	//}

	//a:= complex(0 , 1)
	//b:= complex(1 , 7)
	//c:= complex(2 , 6)
	//d:= complex(3 , 3)
	//e:= complex(4 , 1)
	////f:= complex(5 , 7)
	//g:= complex(6 , 6)
	//h:= complex(7 , 3)

	a := 8 // aqui van los multiplos de llenado malparido miguel
	arreglo := []complex128{}
	arreglo2 := []complex128{}
	arreglo3 := []complex128{}

	for k := 0; k < a; k++ {
		aux := complex(float64(k), 0)
		arreglo = append(arreglo, aux)
	}

	arreglo2 = recursivaFourier2(arreglo...)
	//fmt.Println(a)
	///fmt.Println(imag(a))
	//fmt.Println(arreglo)

	//for i:=0; i < len(arreglo); i++{
	//		fmt.Println(imag(arreglo[i]))
	//	}

	// a = complex(3, imag(a))
	fmt.Println(arreglo2, "Arreglo ")
	arreglo3 = transformadaFormula(arreglo2[0], arreglo2[1])
	fmt.Println(arreglo3)
}
