package rpc

import (
	"github.com/Sirupsen/logrus"
	context "golang.org/x/net/context"
)

type Graph struct {
}

func (g Graph) Query(context.Context, *QueryRequest) (*QueryReply, error) {
	logrus.Print("hit")
	v := &Vertex{Id: "test"}
	m := make(map[string]*Vertex)
	m[v.GetId()] = v
	reply := &QueryReply{Properties: m, Text: "hi ross"}
	return reply, nil
}
