package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"AddServiceApp",
		"POST",
		"/v1/application/{appId}/service",
		AddServiceApp,
	},

	Route{
		"CreateApp",
		"POST",
		"/v1/application",
		CreateApp,
	},

	Route{
		"DeleteApp",
		"DELETE",
		"/v1/application/{appId}",
		DeleteApp,
	},

	Route{
		"DeleteAppService",
		"DELETE",
		"/v1/application/{appId}/service/{serviceId}",
		DeleteAppService,
	},

	Route{
		"GetApp",
		"GET",
		"/v1/application/{appId}",
		GetApp,
	},

	Route{
		"GetAppServiceStatus",
		"GET",
		"/v1/application/{appId}/service/{serviceId}",
		GetAppServiceStatus,
	},

	Route{
		"GetAppServices",
		"GET",
		"/v1/application/{appId}/service",
		GetAppServices,
	},

	Route{
		"ListApp",
		"GET",
		"/v1/application",
		ListApp,
	},

	Route{
		"UpdateApp",
		"PUT",
		"/v1/application/{appId}",
		UpdateApp,
	},

	Route{
		"ListService",
		"GET",
		"/v1/service",
		ListService,
	},
}
