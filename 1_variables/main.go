package main

import "fmt"

func main() {
	/*
		Las variables en Go son de tipado estático, por ende
		hay que declarar de que tipo es cada variable, y no
		pueden cambiar de tipo

		Nota: Las variables en Go deben nombrarse siguiendo
		las reglas del CamelCase
	*/

	/*
		Tipos de datos:

		Números enteros: 					byte, int, int8, int16, int32, int64
		Números enteros solo positivos: 	uint, uint8, uint16, uint32, uint64
		Números decimales: 					float32, float64
		Números complejos: 					complex64, complex128
		Booleanos: 							bool
		Texto: 								string
	*/

	// Declaración de variables y su tipo

	var numero int    // Valor inicial de cualquier tipo de número -> 0
	var nombre string // Valor inicial de un string -> ""
	var bandera bool  // Valor inicial de un booleano -> false

	// Declaración dinamica de variable a partir de su valor inicial

	edad := 20

	/*
		El operador := asigna el tipo de la variable automaticamente
		a partir del valor asignado al inicializarla, siendo en este caso
		una variable de tipo int, ya que se le asigna el tipo int debido al
		valor 20 asignado al ser inicializada la variable edad.
	*/

	// otro ejemplo del uso del operador :=

	apellido := "Garcia"

	// Imprimir las variables en la terminal

	fmt.Println(numero)
	fmt.Println(nombre)
	fmt.Println(bandera)
	fmt.Println(edad)
	fmt.Println(apellido)

	/*
		Nota: Un programa no se puede compilar si se crea una variable y/o importa
		un paquete y no se usa
	*/
}
