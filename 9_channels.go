package main

import "fmt"

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	msg := <-channel
	fmt.Println("Recibí del canal: " + msg)

	msg = <-channel
	fmt.Println("Recibí del canal: " + msg)
}
