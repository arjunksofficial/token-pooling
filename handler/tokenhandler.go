package handler

import (
	"net/http"

	"github.com/arjunksofficial/token-pooling/infrastructure"
	"github.com/arjunksofficial/token-pooling/lru"
	"github.com/arjunksofficial/token-pooling/usecase"
)

//TokenHandler .
type TokenHandler struct {
	UseCase usecase.UseCaseInterface
}

// Welcome gets welcome page
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

// GetToken gets token
func (th *TokenHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(th.UseCase.GetToken()))
}

// Stats Gets token usage statistics
func (th *TokenHandler) Stats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(th.UseCase.Stats()))
}

// NewTokenHandler ..
func NewTokenHandler(db *infrastructure.SQL, cache *lru.Cache, currentkey *int) *TokenHandler {
	return &TokenHandler{
		UseCase: usecase.NewUsecase(db, cache, currentkey),
	}
}
