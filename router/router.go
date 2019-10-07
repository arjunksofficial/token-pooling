package router

import (
	"github.com/arjunksofficial/token-pooling/handler"
	"github.com/arjunksofficial/token-pooling/infrastructure"

	"github.com/arjunksofficial/token-pooling/lru"
	"github.com/go-chi/chi"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux        *chi.Mux
	SQLHandler *infrastructure.SQL
	Cache      *lru.Cache
	CurrentKey *int
}

// SetupHandler define routes for Router
func (r Router) SetupHandler() {
	th := handler.NewTokenHandler(r.SQLHandler, r.Cache, r.CurrentKey)
	r.Mux.Route("/", func(cr chi.Router) {
		cr.Get("/", handler.Welcome)
		cr.Get("/token", th.GetToken)
		cr.Get("/stats", th.Stats)
	})

}
