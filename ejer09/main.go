package main

import "fmt"

func main() {

	map2()
}

func map1() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["mexico"] = "D.F"
	paises["argentina"] = "buenos aires"

	fmt.Println(paises["mexico"])
	fmt.Println(paises)

}

func map2() {
	campeonato := map[string]int{
		"barcelona":    39,
		"real madrid":  38,
		"chivas":       37,
		"boca juniors": 30}

	fmt.Println(campeonato)

	campeonato["riber plate"] = 25
	campeonato["chivas"] = 55
	delete(campeonato, "real madrid")

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("el equipo %s, tiene el puntaje de %d \n", equipo, puntaje)
	}
	puntaje, existe := campeonato["mineiro"]
	fmt.Printf("el puntaje del equipo es %d, y el equipo existe %t \n", puntaje, existe)
}
