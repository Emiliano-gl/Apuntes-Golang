package main

import "fmt"

/*
	En la libreria fmt también viene una función que nos permite
	leer información desde terminal, siendo esta Scanf
*/

func main() {
	fmt.Println("Hola mundo")

	/*
		Para mostrar información en pantalla tenemos las siguientes
		opciones:

		fmt.Print(data)
		fmt.Println(data)

		Siendo la unica diferencia que en fmt.Print toda la información
		se muestra en la misma línea, mientras que fmt.Println agrega
		un salto de línea al final.

		Tenemos una tercera opción la cual es fmt.Printf, esta nos
		permite interpolar información mediante verbos como se
		hacía en C para los diferentes tipo de datos, ejemplo:
	*/

	edad := 20
	fmt.Printf("Mi edad es: %d\n", edad)

	/*
		Como se puede ver, gracias a fmt.Printf se puede interpolar
		la edad en el string al momento de mostrarlo en pantalla, para
		más información buscar en:

		https://golang.org/pkg/fmt/
	*/

	/*
		Para leer datos desde terminal usaremos la función fmt.Scanf
		de la siguiente manera

		fmt.Scanf(verbo + \n, &variable)

		Siendo los verbos los mismos usados en Printf segun el tipo
		de dato a leer, además hay que agregar un \n junto al verbo
		para evitar errores al momento de ingresar datos, ademas
		tendremos que agregar la dirección de la variable donde se
		guardará la información obtenida, Ejemplo:
	*/

	var numero int
	fmt.Scanf("%d\n", &numero)
	fmt.Printf("El número es: %d\n", numero)

	/*
		Nota: Para obtener la derección de una variable es necesario
		usar el operador &, como se puede ver acontinuación

		var edad int
		fmt.Println(&edad)

		Esto mostraria en pantalla la dirección de la variable edad
	*/

}
