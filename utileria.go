package main

import "math"

// elevar es una función que realiza la operación de elevar una base a un exponente dado.
// Recibe una base de tipo float64 y un exponente de tipo int como argumentos.
// Devuelve el resultado de elevar la base al exponente, que es un número de tipo float64.
func elevar(base float64, exponente int) float64 {
	// Verifica si el exponente es mayor que 0.
	if exponente > 0 {
		valorini := base

		// Itera desde 1 hasta el exponente y multiplica la base por sí misma en cada iteración.
		for i := 1; i < exponente; i++ {
			base = base * valorini
		}

		// Devuelve el resultado de la operación de potencia.
		return base
	} else {
		// Si el exponente es 0 o negativo, devuelve 1.
		return 1
	}
}

// factorial es una función que calcula el factorial de un número entero.
// Recibe un número entero x como argumento.
// Devuelve el factorial de x, que es un número entero.
func factorial(x int) int {
	factorial := 1

	// Verifica si x es mayor que 0.
	if x > 0 {
		// Calcula el factorial de x multiplicando todos los números desde 1 hasta x.
		for i := 1; i <= x; i++ {
			factorial *= i
		}

		// Devuelve el factorial calculado.
		return factorial
	} else {
		// Si x es 0 o negativo, devuelve 1.
		return 1
	}
}

// seno es una función que calcula el seno de un número utilizando la serie de Taylor.
// Recibe un número de tipo float64 como argumento x.
// Devuelve el seno de x, que es un número de tipo float64.
func seno(x float64) float64 {
	var seno float64 = 0

	// Calcula el seno utilizando la serie de Taylor con 10 términos.
	for n := 0; n < 10; n++ {
		seno += elevar(-1, n) * (elevar(x, (2*n+1)) / float64(factorial(2*n+1)))
	}

	// Devuelve el valor calculado del seno.
	return seno
}

// coseno es una función que calcula el coseno de un número utilizando la serie de Taylor.
// Recibe un número de tipo float64 como argumento x.
// Devuelve el coseno de x, que es un número de tipo float64.
func coseno(x float64) float64 {
	var coseno float64 = 0

	// Calcula el coseno utilizando la serie de Taylor con 10 términos.
	for n := 0; n < 10; n++ {
		coseno += elevar(-1, n) * (elevar(x, (2*n)) / float64(factorial(2*n)))
	}

	// Devuelve el valor calculado del coseno.
	return coseno
}

// arctan es una función que calcula la arcotangente de un número utilizando la serie de Taylor.
// Recibe un número de tipo float64 como argumento x.
// Devuelve la arcotangente de x, que es un número de tipo float64.
func arctan(x float64) float64 {
	var arctan float64 = 0

	// Calcula la arcotangente utilizando la serie de Taylor con 15 términos.
	for n := 0; n < 15; n++ {
		arctan += elevar(-1, n) * (elevar(x, (2*n+1)) / float64(2*n+1))
	}

	// Devuelve el valor calculado de la arcotangente.
	return arctan
}

// traerPar es una función que extrae los elementos en posiciones pares de un slice de números complejos.
// Recibe un número variable de argumentos de tipo complex128 (arre ...complex128).
// Devuelve un nuevo slice de números complejos que contiene solo los elementos en posiciones pares.
func traerPar(arre ...complex128) []complex128 {
	n := len(arre) / 2
	p := 0
	pares := make([]complex128, n)

	// Itera sobre el slice de números complejos y selecciona solo los elementos en posiciones pares.
	for i := 0; i < len(arre); i++ {
		if i%2 == 0 {
			pares[p] = arre[i]
			p++
		}
	}

	// Devuelve el nuevo slice con los elementos en posiciones pares.
	return pares
}

// traerImpar es una función que extrae los elementos en posiciones impares de un slice de números complejos.
// Recibe un número variable de argumentos de tipo complex128 (arre ...complex128).
// Devuelve un nuevo slice de números complejos que contiene solo los elementos en posiciones impares.
func traerImpar(arre ...complex128) []complex128 {
	n := len(arre) / 2
	p := 0
	impares := make([]complex128, n)

	// Itera sobre el slice de números complejos y selecciona solo los elementos en posiciones impares.
	for i := 0; i < len(arre); i++ {
		if i%2 != 0 {
			impares[p] = arre[i]
			p++
		}
	}

	// Devuelve el nuevo slice con los elementos en posiciones impares.
	return impares
}

// potenciaDeDos es una función que verifica si un número es una potencia de dos.
// Recibe un número entero n como argumento.
// Devuelve un valor booleano indicando si n es una potencia de dos.
func potenciaDeDos(n int) bool {
	if n <= 0 {
		return false
	}

	// Calcula el logaritmo en base 2 de n utilizando la función Log2 del paquete math.
	log2 := math.Log2(float64(n))

	// Verifica si el logaritmo truncado es igual al logaritmo original para determinar si n es una potencia de dos.
	return math.Floor(log2) == log2
}
