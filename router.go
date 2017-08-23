package main

import (
	"github.com/gorilla/mux"
)

//APINewRouter implements mux.NewRouter

func APINewRouter(routes Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	return router
}

func AddRoutes(router *mux.Router, routes Routes) {
	for _, route := range routes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler).
			Name(route.Name)
	}
}

//func AddV1Routes(r *mux.Router,routes Routes) *mux.Router {
//	path_prefit := "v1";
//
//	return AddRoutes(r, routes)
//}
