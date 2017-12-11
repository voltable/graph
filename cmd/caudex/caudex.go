package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

const (
	path = "../../browser/www/static"
)

func main() {
	flag.Parse()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
	http.Handle("/", push())
	logrus.Print(http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", nil))
}

func push() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			// Push is supported. Try pushing rather than
			// waiting for the browser request these static assets.
			if err := pusher.Push(path+"/index.bundle.js", nil); err != nil {
				logrus.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})
}

const indexHTML = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Caudex Browser</title>
  </head>
  <body>
  <script type="text/javascript" src="/static/index.bundle.js"></script></body>
</html>
`
