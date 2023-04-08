package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaración de variables enteras

	// base := 12
	// var altura int = 14
	// var area int

	// fmt.Println("base", base)
	// fmt.Println("altura", altura)
	// fmt.Print("area", area)

	//Zero value

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	// Suma

	x := 50
	y := 10
	result := x + y
	fmt.Println("Suma", result)

	//Resta

	result = x - y
	fmt.Println("Resta", result)

	//Multiplicación

	result = x * y
	fmt.Println("Multiplicación", result)

	//Multiplicación

	result = x / y
	fmt.Println("División", result)

	// Modulo/Residuo

	result = x % y
	fmt.Println("Módulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Incremental
	x--
	fmt.Println("Incremental", x)
}
