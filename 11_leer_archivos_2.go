package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Open("./Archivos de Pruebas/lorem_ipsum.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(archivo)

	var i int

	for scanner.Scan() {
		linea := scanner.Text()
		i++
		fmt.Print(i, "- ")
		fmt.Println(linea)
	}

	archivo.Close()
}
