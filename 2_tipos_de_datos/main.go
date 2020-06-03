package main

import (
	"fmt"
	"strconv"
)

/*
	La libreria strconv nos permitirá convertir diferentes
	tipo de datos en otros
*/

func main() {
	/*
		Al querer concatenar un entero con un número obtendremos un error
		esto debido a que no se puede concatener un entero con una cadena,
		un ejemplo de esto es el siguiente:

		edad := 20
		fmt.Println("Mi edad es: " + edad)

		Este código generaría el error antes mencionado debido a la diferencia de
		tipos, siendo uno string y el otro de tipo int, para solucionar esto
		usamos la libreria strconv, a continuación se muestra un ejemplo
	*/

	edad := 20
	edadStr := strconv.Itoa(edad)

	fmt.Println("Mi edad es: " + edadStr)

	/*
		Como se puede ver, con la función Itoa se puede convertir un número
		entero en un string, siendo de esta manera como se soluciona el problema
		de los tipos de datos.

		otra forma de hacerlo es en la misma línea:

		fmt.Println("Mi edad es: " + strconv.Itoa(edad))
	*/

	/*
		¿Qué pasa si queremos convertir un string a un entero?

		para eso usamos la función Atoi que viene en strconv,
		siendo esta la función inversa de Itoa
	*/

	numeroString := "33"
	numero, _ := strconv.Atoi(numeroString)

	fmt.Println(numero + 7)

	/*
		Con el código anterior hemos convertido un string en un
		número entero, siendo así que ahora se le pude sumar el
		número 7, algo importante a notar es que la función Atoi
		regresa más de un valor, siendo el segundo valor el error
		en caso que no se pueda convertir el string en entero.

		Nota: para ignorar una variable retornada, es necesario poner
		el operador _ en la posición de la variable a ignorar, como se
		mostró en el código anterior
	*/
}
