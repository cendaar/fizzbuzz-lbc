package handler

import (
	"fmt"
	"github.com/baqtiste/fizzbuzz/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func fizzbuzzRoutes(router chi.Router) {
	router.Get("/", fizzbuzz)
}

func fizzbuzz(w http.ResponseWriter, r *http.Request) {
	fizzbuzz := &models.Fizzbuzz{}
	if err := render.Bind(r, fizzbuzz); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	fmt.Println(fizzbuzz)
}