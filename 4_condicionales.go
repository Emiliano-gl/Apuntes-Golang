package main

import "fmt"

func main() {
	/*
		Las condicionales nos permiten manejar el flujo de los
		programas, en Go la tenemos la siguiente sintaxis:

		if condición {
			código a ejecutar
		}

		Además del if contamos con la sentencias:

		* else if
		* else

		Es importante que al usar el else if y else, se
		debe poner al terminar la {} de cierre, ya que
		ponerlo en la línea siguiente generará un error

		Ejemplo:
	*/

	var edad int
	fmt.Print("Ingrese tu edad: ")
	fmt.Scanf("%d\n", &edad)

	if edad <= 0 {
		fmt.Println("No se puede tener una edad así")
	} else if edad < 18 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Eres mayor de edad")
	}

	/*
		Los operadores lógicos con los que contamos son:

		== 	igual a
		!= 	diferente a
		>	mayor que
		<	menor que
		>=	mayor o igual que
		<=	menor o igual que
		&&	and
		||	or
	*/
}
