package handler

import (
	"github.com/cendaar/fizzbuzz/models"
	"github.com/cendaar/fizzbuzz/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func fizzbuzzRoutes(router chi.Router) {
	router.Post("/", fizzbuzz)
	router.Get("/", helloworld)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "helloworld")
}

func fizzbuzz(w http.ResponseWriter, r *http.Request) {
	fizzbuzzModel := &models.Fizzbuzz{}
	if err := render.Bind(r, fizzbuzzModel); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	fizzbuzzService := services.NewFizzbuzzService(redisInstance)
	output := fizzbuzzService.ComputeFizzbuzz(fizzbuzzModel)

	render.PlainText(w, r, output)
}