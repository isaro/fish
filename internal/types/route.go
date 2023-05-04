package types

import "net/http"

type Route struct {
	Method      string
	HandlerFunc http.HandlerFunc
	Route       string
	Name        string
}
