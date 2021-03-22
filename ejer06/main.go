package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	fmt.Println(dos(5))
	//fmt.Println(tres(1))
	numero, estado := tres(1)
	fmt.Println(numero, estado)
	fmt.Println(Calculo(3, 4, 5, 6, 7, 8, 9))
	fmt.Println(Calculo(5, 9, 6))
	fmt.Println(Calculo(333, 3, 3, 3, 3, 3))
}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (z int) {
	z = numero * 3
	return
}

func tres(numero int) (int, bool) {
	if numero == 1 {
		return numero * 9, true
	} else {
		return numero * 0, false
	}
}
func Calculo(numero ...int) int {
	total := 0
	total2 := 0
	for _, num := range numero {
		total = total + num
	}

	for item, num := range numero {
		total2 = total2 + num

		fmt.Printf("\n item \n %d", item)
	}
	return total2
}
