package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Emiliano")
	fmt.Println("Que aburridamente lento")
}

func miNombreLento(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
