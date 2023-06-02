package main

import (
	"context"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
)

const pi = 3.14159265358979323846264338327950288419716939937510582097494459

func ordenarFourier(entrada ...complex128) []complex128 {
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

		slicePar = ordenarFourier(slicePar...)
		sliceImpar = ordenarFourier(sliceImpar...)

		sliceSalida = slicePar

		for j := 0; j < len(sliceImpar); j++ {
			sliceSalida = append(sliceSalida, sliceImpar[j])
		}

		return sliceSalida
	}
}

func transformadaFormula(fft ...complex128) []complex128 {
	n := len(fft)
	if n > 1 {

		par := make([]complex128, n/2)
		par = traerPar(fft...)

		impar := make([]complex128, n/2)
		impar = traerImpar(fft...)

		par = transformadaFormula(par...)

		impar = transformadaFormula(impar...)

		w := complex(0, 0)
		z := complex(0, 0)

		for m := 0; m < n/2; m++ {
			aux := float64(m) / float64(n)
			w = complex(float64(coseno(2*pi*float64(aux))), imag(w))
			w = complex(real(w), float64(-1*seno(2*pi*float64(aux))))
			z = complex(real(w)*real(impar[m])-imag(w)*imag(impar[m]), real(w)*imag(impar[m])+imag(w)*real(impar[m]))
			fft[m] = complex(real(par[m])+real(z), imag(fft[m]))
			fft[m] = complex(real(fft[m]), imag(par[m])+imag(z))

			fft[m+(n>>1)] = complex(real(par[m])-real(z), imag(fft[m+n/2]))
			fft[m+(n>>1)] = complex(real(fft[m+n/2]), imag(par[m])-imag(z))

		}
	}

	return fft
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

func serv() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := New(":8080")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Println("server started")

	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("server stoped")
}

func main() {
	serv()
	//var cantidadDatos int
	//fmt.Print("Ingrese un número entero: ")
	//fmt.Scan(&cantidadDatos)
	/*cantidadDatos := 8

	if potenciaDeDos(cantidadDatos) {
		fmt.Println("El número ingresado es una potencia de 2.")
		slice := []complex128{}
		sliceOrdenado := []complex128{}
		fft := []complex128{}

		for k := 0; k < cantidadDatos; k++ {
			aux := complex(float64(k), 0)
			slice = append(slice, aux)
		}

		sliceOrdenado = ordenarFourier(slice...)
		fmt.Println(sliceOrdenado, " slice ordenado")
		fft = transformadaFormula(slice...)
		fmt.Println("\n\n\nDespues de aplicar la fft:\n", fft, "\n\n")

		JSON(fft...)
		//serv()



	} else {
		fmt.Println("El número ingresado NO es una potencia de 2.")

	}*/

}
