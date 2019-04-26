package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/voltable/graph/cmd/caudex/handlers"
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
	<meta name="viewport" content="width=device-width, minimum-scale=1.0, initial-scale=1.0, user-scalable=yes">
	<title>Caudex Browser</title>
	<link rel="stylesheet" media="screen" href="/static/style.css" />

	<!-- Ensure Web Animations polyfill is loaded since neon-animation 2.0 doesn't import it -->
	<script src="static/web-animations-next-lite.min.js"></script>

	<script src="/static/webcomponents-hi-sd-ce.js" nomodule></script>
	<script src="/static/index.bundle.js"></script>

  </head>
  <body>
  	<cg-app></cg-app>
  </body>
</html>
`
