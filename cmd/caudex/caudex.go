package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler.ServeHTTP(gzw, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("../../browser/www"))
	http.Handle("/", Gzip(fs))

	logrus.Print("Listening...")
	http.ListenAndServe(":3000", nil)
}
