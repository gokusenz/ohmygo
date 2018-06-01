package http

import (
	"net/http"

	"github.com/pallat/ohmygo/myapp"
)

type Handler struct {
	UserService myapp.UserService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

func NewServeMux() *http.ServeMux {
	return http.NewServeMux()
}
