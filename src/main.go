package main

import (
	"fmt"
	"math"
)

func main() {
	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("suma:", result)

	// Suma
	result = y - x
	fmt.Println("resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	// Division
	result = y / x
	fmt.Println("Division", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo", result)

	// Incremental
	x++
	fmt.Println("Incrementar", x)

	// Decremental
	x--
	fmt.Println("Decrementar", x)

	//Reto
	// Area del trapecio y circulo
	altura := 4
	base1 := 10
	base2 := 4
	radio := 2
	resultTrapecio := (base1 + base2) * altura / 2
	fmt.Println("area trapecio", resultTrapecio)
	resultRadio := math.Pi * math.Pow(float64(radio), 2)
	fmt.Println("area circulo", resultRadio)
}
