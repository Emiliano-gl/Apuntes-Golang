package main

import "net/http"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./Archivos de Pruebas/holamundo.html")
	})
	http.ListenAndServe(":3000", nil)
}
