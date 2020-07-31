package main

import "fmt"

// Una struct funciona parecido a una clase

type User struct {
	edad             int
	nombre, apellido string
}

func (user *User) nombreCompleto() string {
	return user.nombre + " " + user.apellido
}

func main() {
	usuario := new(User)

	usuario.nombre = "Emiliano"
	usuario.apellido = "Guerrero"

	fmt.Println(usuario.nombreCompleto())
}
