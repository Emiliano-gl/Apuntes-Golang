package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileData, err := ioutil.ReadFile("./Archivos de Pruebas/archivo_de_prueba.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(fileData))
}
