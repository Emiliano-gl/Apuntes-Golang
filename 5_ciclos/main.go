package main

import "fmt"

func main() {
	/*
		En Go solo existe el ciclo for, pero a difencia de otros
		lenguajes de progrmación, el ciclo for en Go dependiendo de como
		lo uses, funcionará como un ciclo while o un for.

		Los ciclos for se usan de la misma forma que otros lenguajes
		de programación como puede ser Java o C++.

		for varieable; condicional; modificador {
			código a ejecutar
		}

		ejemplo:
	*/

	for i := 0; i < 10; i++ {
		fmt.Println("i vale: ", i)
	}

	/*
		Para hacer un while, basta con solo dejar la condicional
		en el ciclo for, con esto lograremos emular un ciclo
		while, no olvides hacer una forma de salir del ciclo.

		for condicional {
			código a ejecutar
		}

		ejemplo:
	*/

	contador := 0

	for contador < 5 {
		fmt.Println("Contador while: ", contador)
		contador++
	}

	/*
		En Go tenemos una tercera opción que es una variente del
		ciclo while, solo que este es un ciclo infinito, para poder
		hacer esto solo basta con poner el for sin nada más, tenemos
		que tener en cuenta que el ciclo debe tener una forma de salir
		de el.

		for {
			código a ejecutar
		}

		Con esto tendríamos un bucle infinito que puede sernos útil
		dependiendo de la situación.
	*/
}
