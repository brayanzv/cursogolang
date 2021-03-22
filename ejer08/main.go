package main

import (
	"fmt"
)

var tabla [10]int

func main() {
	tabla[1] = 1
	tabla[5] = 15
	var tabla2 = [10]int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11}
	tabla3 := [10]int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11}

	fmt.Println(tabla2)
	fmt.Println(tabla3)

	//manera de recorrer una tabla 'len' es para saber el tamaño del array

	for i := 0; i < len(tabla3); i++ {
		fmt.Println(tabla3[i])
	}
	slices()
	variante4()
}

func matriz() {
	var matriz [5][7]int // se inicializa en 1
	matriz[4][3] = 4     //la array o matriz se inica en 0
	fmt.Println(matriz[4][3])

}

func slices() {
	matriz := []int{2, 3, 4, 4}
	fmt.Println(matriz)
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4] // me permite particionar una array mas grande en una mas pequeña delimitada
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo es %d,  capacidad es %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Println(len(nums), cap(nums))
	}
}
