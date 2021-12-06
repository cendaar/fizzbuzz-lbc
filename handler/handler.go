package handler

import (
	"github.com/cendaar/fizzbuzz/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

var redisInstance *db.RedisInstance

func NewHandler(ri *db.RedisInstance) http.Handler {
	redisInstance = ri
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	router.Route("/", fizzbuzzRoutes)

	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(ErrMethodNotAllowed.StatusCode)
	_ = render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(ErrNotFound.StatusCode)
	_ = render.Render(w, r, ErrNotFound)
}