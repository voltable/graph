package handlers

import (
	"net/http"
	"strings"

	"github.com/RossMerr/Caudex.Graph/rpc"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func init() {
	grpcServer := grpc.NewServer()
	rpc.RegisterGraphServer(grpcServer, rpc.Graph{})
}

// Grpc does a ppush over http2 for any grpc client connection
func Grpc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
