package handler

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/status"
)

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	// TODO error handling
	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))
}
