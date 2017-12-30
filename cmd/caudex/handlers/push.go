package handlers

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Push a http.Handler for grpc and http2 push
func Push(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pusher, ok := w.(http.Pusher)
		if ok {
			if err := pusher.Push("/static/index.bundle.js", nil); err != nil {
				logrus.Printf("Failed to push: %v", err)
			}
			// if err := pusher.Push("/static/vertex.proto.json", nil); err != nil {
			// 	logrus.Printf("Failed to push: %v", err)
			// }
		}

		next.ServeHTTP(w, r)
	})
}
