package main

import (
	"fmt"
)

//la variable calculo es definida como una funcion anonima
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("la suma es: %d \n", calculo(7, 5))
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("la resta es: %d \n", calculo(7, 9))

	Operaciones()

	tabladel := 2
	mitabla := tabla(tabladel)

	for i := 1; i < 11; i++ {
		fmt.Println(mitabla())
	}
}

func Operaciones() {
	resultado := func() int {
		a := 2
		b := 3
		return a * b
	}
	fmt.Printf("la multiplicacion es: %d \n", resultado())
}

//closures..

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	//al llamar a tabla este solo toma el retorno func () int y da por inicializado a secuencia
	return func() int {
		secuencia += 1 //esto es igual a seciencia =secuencia + 1
		return numero * secuencia
	}
}
