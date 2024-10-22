package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r *mux.Router
}

func New(router *mux.Router) *api {
	return &api{r: router}
}

func (api *api) Handle() {
	api.r.HandleFunc("/api/hello", hello).Methods(http.MethodGet)
	//api.r.HandleFunc("/api/goodbye"), googoodbye)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
