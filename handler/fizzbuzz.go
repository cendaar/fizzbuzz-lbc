package handler

import (
	"github.com/cendaar/fizzbuzz/models"
	"github.com/cendaar/fizzbuzz/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"log"
	"net/http"
)

func fizzbuzzRoutes(router chi.Router) {
	router.Post("/", fizzbuzz)
	router.Get("/", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "helloworld")
}

func fizzbuzz(w http.ResponseWriter, r *http.Request) {
	fizzbuzzModel := &models.Fizzbuzz{}

	if err := render.Bind(r, fizzbuzzModel); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err))
		return
	}

	fizzbuzzService := services.NewFizzbuzzService()
	output := fizzbuzzService.ComputeFizzbuzz(fizzbuzzModel)

	statsService := services.NewStatsService(redisInstance)
	err := statsService.HandleFizzbuzzRequest(fizzbuzzModel)
	if err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		log.Fatalln(err)
	}

	render.PlainText(w, r, output)
}