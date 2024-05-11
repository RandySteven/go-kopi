package apps

import (
	"github.com/RandySteven/go-kopi/enums"
	"github.com/gorilla/mux"
	"net/http"
)

type EndpointRouter struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
	method  string
}

func RegisterEndpointRouter(path, method string, handler func(w http.ResponseWriter, r *http.Request)) *EndpointRouter {
	return &EndpointRouter{path: path, handler: handler, method: method}
}

func NewEndpointRouters(h *Handlers) map[enums.RouterPrefix][]EndpointRouter {
	endpointRouters := make(map[enums.RouterPrefix][]EndpointRouter)

	endpointRouters[enums.AuthPrefix] = []EndpointRouter{
		*RegisterEndpointRouter("/register", http.MethodPost, h.UserHandler.RegisterUser),
		*RegisterEndpointRouter("/login", http.MethodPost, h.UserHandler.LoginUser),
	}

	return endpointRouters
}

func (h *Handlers) InitRouter(r *mux.Router) {
	mapRouters := NewEndpointRouters(h)

	authRouter := r.PathPrefix(enums.AuthPrefix.ToString()).Subrouter()
	for _, router := range mapRouters[enums.AuthPrefix] {
		authRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}
}
