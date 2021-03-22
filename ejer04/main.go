package main

import (
	"bufio"
	"fmt"
	"os"
)

var num1 int
var num2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("ingresa el primer numero ")
	fmt.Scanln(&num1) //& es un puntero

	fmt.Println("ingrese el segundo numero ")
	fmt.Scanln(&num2)

	fmt.Println("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin) //ingresamos datos mediante teclado.

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = num1 + num2
	fmt.Println(leyenda, resultado)
}
