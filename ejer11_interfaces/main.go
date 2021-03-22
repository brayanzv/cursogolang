package main

import (
	"fmt"
)

//se define los metodos o struct que vamos a usar para implementar esta interface
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	esCarnivoro()
}

//genero humano

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "hombre"
	} else {
		return "mujer"
	}
}

func HumanosRespirando(hu humano) {
	hu.respirar()
	hu.comer()
	fmt.Printf("soy un/a %s, ya estoy respirando y \n", hu.sexo())
}

func main() {
	pedro := new(hombre)
	pedro.esHombre = true
	HumanosRespirando(pedro)

}
