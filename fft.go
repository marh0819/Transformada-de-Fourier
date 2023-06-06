package main

import (
	"fmt"
	"net/http"
)

const pi = 3.14159265358979323846264338327950288419716939937510582097494459

// ordenarFourier es una función que ordena una lista de números complejos utilizando el algoritmo de ordenación de Fourier.
// Recibe una lista de números complejos de tipo complex128 como argumento de entrada.
// Devuelve una lista de números complejos ordenados de tipo complex128.
func ordenarFourier(entrada ...complex128) []complex128 {
	// Verifica si la longitud de la entrada es igual a 1, en cuyo caso devuelve la entrada sin cambios.
	if len(entrada) == 1 {
		return entrada
	} else {
		// Crea tres slices para almacenar los números complejos: slicePar para los índices pares, sliceImpar para los índices impares y sliceSalida para el resultado final.
		slicePar := []complex128{}
		sliceImpar := []complex128{}
		sliceSalida := []complex128{}

		// Itera sobre la entrada y separa los números complejos en los slices correspondientes.
		for i := 0; i < len(entrada); i++ {
			if i%2 == 0 {
				slicePar = append(slicePar, entrada[i])
			} else {
				sliceImpar = append(sliceImpar, entrada[i])
			}
		}

		// Llama recursivamente a la función ordenarFourier para ordenar los slices slicePar y sliceImpar.
		slicePar = ordenarFourier(slicePar...)
		sliceImpar = ordenarFourier(sliceImpar...)

		// Combina los slices slicePar y sliceImpar en el slice sliceSalida.
		sliceSalida = slicePar
		for j := 0; j < len(sliceImpar); j++ {
			sliceSalida = append(sliceSalida, sliceImpar[j])
		}

		// Devuelve el slice sliceSalida que contiene los números complejos ordenados.
		return sliceSalida
	}
}

// transformadaFormula es una función que calcula la transformada utilizando la fórmula de transformada de Fourier rápida (FFT).
// Recibe una lista de números complejos de tipo complex128 como argumento fft.
// Devuelve una lista de números complejos transformados de tipo complex128.
func transformadaFormula(fft ...complex128) []complex128 {
	// Obtiene la longitud de la lista fft.
	n := len(fft)

	// Verifica si n es mayor que 1, lo que indica que se debe realizar la transformada.
	if n > 1 {
		// Crea dos slices de tamaño n/2 para almacenar los números complejos pares e impares.
		par := make([]complex128, n/2)
		par = traerPar(fft...)

		impar := make([]complex128, n/2)
		impar = traerImpar(fft...)

		// Llama recursivamente a la función transformadaFormula para calcular la transformada de las sub-listas par e impar.
		par = transformadaFormula(par...)
		impar = transformadaFormula(impar...)

		// Calcula los valores de w y z necesarios para la transformada.
		w := complex(0, 0)
		z := complex(0, 0)

		for m := 0; m < n/2; m++ {
			aux := float64(m) / float64(n)
			w = complex(float64(coseno(2*pi*float64(aux))), imag(w))
			w = complex(real(w), float64(-1*seno(2*pi*float64(aux))))
			z = complex(real(w)*real(impar[m])-imag(w)*imag(impar[m]), real(w)*imag(impar[m])+imag(w)*real(impar[m]))

			// Calcula los valores de la transformada y los asigna a la lista fft.
			fft[m] = complex(real(par[m])+real(z), imag(fft[m]))
			fft[m] = complex(real(fft[m]), imag(par[m])+imag(z))

			fft[m+(n>>1)] = complex(real(par[m])-real(z), imag(fft[m+n/2]))
			fft[m+(n>>1)] = complex(real(fft[m+n/2]), imag(par[m])-imag(z))
		}
	}

	// Devuelve la lista fft con los números complejos transformados.
	return fft
}

// fft es una función que realiza la transformada rápida de Fourier (FFT) en una cantidad específica de datos.
// Recibe un objeto `http.ResponseWriter` y un objeto `http.Request` como parámetros w y r, respectivamente.
func fft(w http.ResponseWriter, r *http.Request) {
	// Definición de la cantidad de datos a utilizar para la transformada.
	cantidadDatos := 8

	// Establece el encabezado de la respuesta HTTP con el tipo de contenido "application/json".
	w.Header().Set("Content-Type", "application/json")

	// Verifica si la cantidad de datos es una potencia de dos utilizando la función potenciaDeDos.
	if potenciaDeDos(cantidadDatos) {
		fmt.Println("El número ingresado es una potencia de 2.")

		// Crea un slice llamado "slice" para almacenar los números complejos.
		slice := []complex128{}

		// Crea un slice llamado "sliceOrdenado" para almacenar los números complejos ordenados.
		sliceOrdenado := []complex128{}

		// Crea un slice llamado "fft" para almacenar los números complejos transformados.
		fft := []complex128{}

		// Genera una secuencia de números complejos y los agrega al slice "slice".
		for k := 0; k < cantidadDatos; k++ {
			aux := complex(float64(k), 0)
			slice = append(slice, aux)
		}

		// Ordena el slice de números complejos utilizando la función ordenarFourier.
		sliceOrdenado = ordenarFourier(slice...)
		fmt.Println(sliceOrdenado, " slice ordenado")

		// Aplica la transformada de Fourier rápida (FFT) al slice utilizando la función transformadaFormula.
		fft = transformadaFormula(slice...)
		fmt.Println("\nDespués de aplicar la fft:\n", fft)

		// Convierte el slice de números complejos transformados a formato JSON utilizando la función JSON.
		data := JSON(fft...)

		// Escribe la respuesta JSON en el objeto http.ResponseWriter.
		fmt.Fprintf(w, "%v", data)
	} else {
		fmt.Println("\nEl número ingresado NO es una potencia de 2.")
	}
}
