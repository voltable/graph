package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/RossMerr/Caudex.Graph/rpc"
	"github.com/Sirupsen/logrus"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	serverAddr         = flag.String("server_addr", ":8080", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption

	creds, err := credentials.NewClientTLSFromFile("cert.pem", "localhost")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		logrus.Panic(err)
	}
	defer conn.Close()
	client := rpc.NewGraphClient(conn)

	request := &rpc.QueryRequest{}
	request.Text = "hi"
	reply, _ := client.Query(context.Background(), request)

	data, _ := proto.Marshal(reply)

	fmt.Printf("reply %+v\n", data)

}
