package client

import (
	"github.com/Sirupsen/logrus"
	context "golang.org/x/net/context"
)

type Graph struct {
}

func (g Graph) Query(context.Context, *QueryRequest) (*QueryReply, error) {
	logrus.Print("hit")

	return nil, nil
}
