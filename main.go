package main

import (
	"log"
	"net/http"

	"github.com/arjunksofficial/token-pooling/infrastructure"
	"github.com/arjunksofficial/token-pooling/usecase"

	"github.com/arjunksofficial/token-pooling/lru"
	"github.com/arjunksofficial/token-pooling/router"
	"github.com/go-chi/chi"
)

func main() {
	r := newRouter()
	r.SetupHandler()
	log.Println("Server is starting......")
	http.ListenAndServe(":8080", r.Mux)
}

func newRouter() router.Router {
	mux := chi.NewRouter()
	masterdb, err := infrastructure.NewSQL()
	if err != nil {
		log.Panic(err)
	}
	tokens := []usecase.Token{}
	currentkey := 0
	routes := router.Router{
		SQLHandler: masterdb,
		Mux:        mux,
		Cache:      lru.New(12),
		CurrentKey: &currentkey,
	}
	// Populate LRU Cache
	err = masterdb.Master.Table("tokens").Order("count desc").Find(&tokens).Error
	for _, v := range tokens {
		routes.Cache.Add(v.TokenID, v.TokenData)
	}
	return routes
}
