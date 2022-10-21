package http

import (
	"net/http"
)

func ExampleMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {

		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
