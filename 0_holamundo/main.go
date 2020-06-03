/*
	En todo proyecto en Go, el archivo principal debe contener:

	* El nombre del pakage debe ser main como se muestra en el código
	* Debe contener una función main
*/

package main

import "fmt"

/*
	Aquí importamos la libreria fmt que nos permite leer y escribir
	datos desde terminal, siendo mediante la funcion Println con la
	que mostraremos datos en pantalla.

	La forma de mostrar usar la función Println es la siguiente:

	fmt.Println(dato)
*/

func main() {
	fmt.Println("Hola mundo")

	// Nota: es importante ver que en Go no es necesario hacer uso del ;
}
