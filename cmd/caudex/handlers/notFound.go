package handlers

import "net/http"

// Grpc does a ppush over http2 for any grpc client connection
func NotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
