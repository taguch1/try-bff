package handler

import (
	"net/http"
)

// Health handler
type Health interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type healthImpl struct {
}

//NewHealth  handler constructor
func NewHealth() Health {
	return &healthImpl{}
}

func (h *healthImpl) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
