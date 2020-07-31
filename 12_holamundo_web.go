package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Nueva petición")
	io.WriteString(res, "Hola Mundo!")
}
