package apps

import (
	"github.com/RandySteven/go-kopi/enums"
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
