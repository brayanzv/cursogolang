package main

import "fmt"

func main() {

	num := 5
	switch /*num:=5*/ num {
	case 1:
		fmt.Println("el numero es 1")
	case 2:
		fmt.Println("el numero es 2")
	case 3:
		fmt.Println("el numero es 3")
	case 4:
		fmt.Println("el numero es 4")
	case 5:
		fmt.Println("el numero es 5")
	default:
		fmt.Println("el numero es mayor a 5")
	}
}
