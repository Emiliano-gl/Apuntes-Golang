package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string `json:"title"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		cursos := Cursos{
			Curso{
				Title:        "El curso definitivo de GO",
				NumeroVideos: 35,
			},
			Curso{
				Title:        "El curso Dart desde 0",
				NumeroVideos: 56,
			},
			Curso{
				Title:        "El curso NodeJs avanzado",
				NumeroVideos: 47,
			},
		}

		json.NewEncoder(res).Encode(cursos)
	})
	http.ListenAndServe(":3000", nil)
}
