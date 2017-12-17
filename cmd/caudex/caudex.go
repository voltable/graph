package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/RossMerr/Caudex.Graph/rpc"
	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

const (
	path = "../../browser/www/static"
)

func main() {
	flag.Parse()
	grpcServer := grpc.NewServer()
	rpc.RegisterGraphServer(grpcServer, rpc.Graph{})

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
	mux.Handle("/", push(grpcServer))
	logrus.Print(http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", mux))
}

func push(grpcServer *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		} else if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			if err := pusher.Push("/static/index.bundle.js", nil); err != nil {
				logrus.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/vertex.proto.json", nil); err != nil {
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
