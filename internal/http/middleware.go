package http

import (
	"net/http"

	"github.com/google/jsonapi"
)

var (
	headerAccept      = http.CanonicalHeaderKey("accept")
	headerContentType = http.CanonicalHeaderKey("content-type")
)

func withAccept(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerAccept, jsonapi.MediaType)
		next.ServeHTTP(w, r)
	}
}

func withContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentType, jsonapi.MediaType)
		next.ServeHTTP(w, r)
	}
}
