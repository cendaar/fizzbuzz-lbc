package handler

import (
	"github.com/cendaar/fizzbuzz/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func statsRoutes(router chi.Router) {
	router.Get("/", fizzbuzzLeaderboard)
}

func fizzbuzzLeaderboard(w http.ResponseWriter, r *http.Request) {
	statsService := services.NewStatsService(redisInstance)
	output, err := statsService.GetFormattedRequestLeader()
	if err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
	}

	_, _ = io.WriteString(w, output)
}