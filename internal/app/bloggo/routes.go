package bloggo

import (
	"net/http"

	"github.com/bkohler93/bootdev-blog-aggregator/internal/helpers"
	"github.com/go-chi/chi/v5"
)

func (cfg apiConfig) apiV1Router() *chi.Mux {
	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", getReadiness)
	v1Router.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})

	v1Router.Post("/users", cfg.postUsers)
	v1Router.Get("/users", cfg.getUser)

	return v1Router
}
