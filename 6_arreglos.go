package main

import "fmt"

func main() {
	/*
		Para declarar un arreglo en Go, es muy parecido a
		declarar una variable, pero agregando unos [] entes
		del tipo de dato, siendo allí donde va el tamaño
		del arreglo.

		var nombre_del_arreglo [tamaño] tipo_de_dato

		Nota: Los arreglos en Go son estáticos, por ende, no
		es posible cambiar el tamaño o tipo de datos que acepta
		el arreglo.

		ejemplo:
	*/

	var numeros [10]int
	fmt.Println(numeros)

	/*
		Todos los arreglos declarados de la anterior forma
		toman todos sus valores el valor inicial dependiendo
		de su tipo de dato. Por ejemplo un arreglo de 5
		elementos de tipo int se verá así:

		[0,0,0,0,0]

		Para accedes a algun elemento específico del arreglo
		basta con hacerlo con el index del elemento, como se
		muestra a continuación
	*/

	numeros[4] = 5
	fmt.Println(numeros)

	/*
		Para declarar un arreglo con valores definidos haremos
		uso del operador :=

		Nota: Si se meten menos elemtos de los elemtos totales
		los elementos restantes toparan sus valores por defecto
	*/

	edades := [3]int{1, 2, 3}
	fmt.Println(edades)

	// Para saber la longitud de un arreglo, usaremos la función len
	fmt.Println("Números tiene:", len(numeros), "elementos")

	/*
		Para crear un arreglo multi dimensional, es necesario
		agregar un [] más por cada dimensión deseada, por ejemplo
		una matriz 2x2 se haría de la siguiente manera:
	*/

	var matriz [2][2]int
	fmt.Println(matriz)

	/*
		Al igual que un arreglo normal, se puede acceder a los
		datos de la matriz con ayuda del index, ejemplo:
	*/

	matriz[0][0] = 1
	fmt.Println(matriz)
}
