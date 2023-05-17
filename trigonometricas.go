package main

func elevar(valor float64, hasta int) float64 {

	if hasta > 0 {
		valorini := valor

		for i := 1; i < hasta; i++ {
			valor = valor * valorini
		}
		return valor

	} else {
		return 1
	}

}

func factorial(x int) int {

	factorial := 1
	if x > 0 {
		for i := 1; i <= x; i++ {
			factorial *= i
		}

		return factorial
	} else {
		return 1
	}

}

func seno(x float64) float64 {

	var seno float64 = 0

	for n := 0; n < 33; n++ {
		seno += elevar(-1, n) * (elevar(x, (2*n+1)) / float64(factorial(2*n+1)))
	}

	return seno
}

func coseno(x float64) float64 {

	var coseno float64 = 0

	for n := 0; n < 20; n++ {
		coseno += elevar(-1, n) * (elevar(x, (2*n)) / float64(factorial(2*n)))
	}

	return coseno
}
