package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler *httptransport.Server
}

type Routes []Route


