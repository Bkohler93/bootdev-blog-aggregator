package bloggo

import (
	"github.com/go-chi/chi/v5"
)

func (cfg apiConfig) apiV1Router() *chi.Mux {
	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handlerGetReadiness)
	v1Router.Get("/err", handlerGetErr)

	v1Router.Post("/users", cfg.handlerPostUsers)
	v1Router.Get("/users", cfg.middlewareAuth(cfg.handlerGetUser))

	v1Router.Post("/feeds", cfg.middlewareAuth(cfg.handlerPostFeed))
	v1Router.Get("/feeds", cfg.handlerGetAllFeeds)

	return v1Router
}
