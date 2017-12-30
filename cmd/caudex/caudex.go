package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/RossMerr/Caudex.Graph/cmd/caudex/handlers"
	"github.com/Sirupsen/logrus"
)

var httpAddr = flag.String("https", ":8080", "Listen address")

const (
	static = "../../browser/www/static"
	json   = "../../browser/www/json"
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static))))
	mux.Handle("/json/", http.StripPrefix("/json/", http.FileServer(http.Dir(json))))

	mux.Handle("/", handlers.NotFound(handlers.Grpc(handlers.Push(Index()))))
	logrus.Print("Caudex ", http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", mux))
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
  <script rel="preload" type="text/javascript" src="/static/index.bundle.js"></script></body>
</html>
`
