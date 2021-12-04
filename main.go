package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strconv"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
	})

	router.Use(
		c.Handler,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	//public routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r,  fizzbuzz(4, 5, 10000, "hello", "world"))
	})

	return router
}

func main() {
	router := Routes()

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatalln("Server error:", err.Error())
	}
}

func fizzbuzz(int1 int, int2 int, limit int, str1 string, str2 string) string {
	var output string

	for i:=1; i<=limit; i++ {
		switch {
		case i % (int1*int2) == 0:
			output += str1+str2
		case i % int1 == 0:
			output += str1
		case i % int2 == 0:
			output += str2
		default:
			output += strconv.Itoa(i)
		}

		if i != limit {
			output += ","
		}
	}

	return output
}