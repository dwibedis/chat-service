package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Router struct {
	router *mux.Router
}

func NewRouter(router *mux.Router) *Router {
	return &Router{router:router}
}

func (r *Router) RegisterNewHandler(path string, f func(http.ResponseWriter, *http.Request), method string) {
	r.router.HandleFunc(path, f).Methods(method)
}

func (r *Router) ListenAndServe(portNumber int) error {
	port := ":"+strconv.Itoa(portNumber)
	err := http.ListenAndServe(port, r.router)
	if err != nil {
		return err
	}
	return nil
}
