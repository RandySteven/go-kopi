package routes

import (
	"net/http"

	"github.com/RandySteven/go-kopi/enums"
	api_http "github.com/RandySteven/go-kopi/handlers/https"
	"github.com/RandySteven/go-kopi/middlewares"
	"github.com/gorilla/mux"
)

type (
	HandlerFunc func(w http.ResponseWriter, r *http.Request)

	Router struct {
		methodName  string
		path        string
		handler     HandlerFunc
		method      string
		middlewares []enums.Middleware
	}

	RouterPrefix map[enums.RouterPrefix][]*Router
)

func NewEndpointRouters(api *api_http.HTTPs) RouterPrefix {
	endpointRouters := make(RouterPrefix)

	return endpointRouters
}

func InitRouter(routers RouterPrefix, r *mux.Router) {
	middleware := middlewares.NewMiddlewares()
	clientMiddleware := middlewares.RegisterClientMiddleware(middleware)
	serverMiddleware := middlewares.RegisterServerMiddleware(middleware)

	r.Use(
		serverMiddleware.LoggingMiddleware,
		serverMiddleware.CorsMiddleware,
		serverMiddleware.TimeoutMiddleware,
		serverMiddleware.CheckHealthMiddleware,
		clientMiddleware.AuthenticationMiddleware,
		clientMiddleware.RateLimiterMiddleware,
	)
}
