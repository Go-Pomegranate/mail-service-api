// package routers defines routing - redirecting requests to correct handler
package routers

import (
	"github.com/Go_Pomegranate/controllers"
	"github.com/Go_Pomegranate/logger"
	"github.com/Go_Pomegranate/repositories"
	"github.com/gorilla/mux"
	"net/http"
)

//every Route has 4 field:
// Name, type of request e.g POST or GET, Pattern - default slash "/" and
// handlerFunc which is handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var controller = &controllers.Controller{repositories.MailRepository{}}

// slice of routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
}

// init mux router with appropriate handlers, adding the logger which monitor web request traffic
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		// starting listener of logger for every single route
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
