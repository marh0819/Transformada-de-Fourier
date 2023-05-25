package main

func elevar(base float64, exponente int) float64 {

	if exponente > 0 {
		valorini := base

		for i := 1; i < exponente; i++ {
			base = base * valorini
		}
		return base

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

	for n := 0; n < 10; n++ {
		seno += elevar(-1, n) * (elevar(x, (2*n+1)) / float64(factorial(2*n+1)))
	}

	return seno
}

func coseno(x float64) float64 {

	var coseno float64 = 0

	for n := 0; n < 10; n++ {
		coseno += elevar(-1, n) * (elevar(x, (2*n)) / float64(factorial(2*n)))
	}

	return coseno
}

func arctan(x float64) float64 {
	var arctan float64 = 0

	for n := 0; n < 15; n++ {
		arctan += elevar(-1, n) * (elevar(x, (2*n+1)) / float64(2*n+1))
	}

	return arctan
}
