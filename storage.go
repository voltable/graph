package graph

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type Storage interface {
	ForEach() Iterator
	Fetch(string) (*vertices.Vertex, error)
}
