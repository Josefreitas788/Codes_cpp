package main

import "fmt"

func exponenciacao(base float64, expoente int) float64 {
	if expoente == 0 {
		return 1
	}
	if expoente == 1 {
		return base
	}

	if expoente%2 == 0 {
		return exponenciacao(base*base, expoente/2)
	} else {
		return base * exponenciacao(base*base, (expoente-1)/2)
	}

}

func main() {
	fmt.Println(exponenciacao(2, 6))
}
