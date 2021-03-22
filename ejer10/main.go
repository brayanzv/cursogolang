package main

import (
	"fmt"
	"time"

	us "./user"
)

/*type usuario struct {
	id     int
	name   string
	date   time.Time
	status bool
}*/

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUser(1, "Pablo tilota", time.Now(), true)
	fmt.Println(u.Usuario)
}

/*func User1() {
	user := new(usuario)

	user.id = 10
	user.name = "Brayan"

	fmt.Println(user)
}*/
