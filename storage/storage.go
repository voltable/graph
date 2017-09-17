package storage

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type Storage interface {
	ForEach() func() enumerables.Iterator
	Fetch() func(string) (*vertices.Vertex, error)
}
