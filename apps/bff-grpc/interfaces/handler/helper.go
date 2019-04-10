package handler

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	// TODO error handling
	w.WriteHeader(http.StatusBadGateway)
}
