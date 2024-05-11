package apps

import (
	"github.com/RandySteven/go-kopi/middlewares"
	"github.com/gorilla/mux"
)

func RegisterMiddleware(r *mux.Router) {
	middlewares.CorsMiddleware(r)
}
