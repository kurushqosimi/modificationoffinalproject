package middlewares

import "net/http"

func SetContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/xml; charset=utf-8")
			h.ServeHTTP(w, r)
		},
	)
}
