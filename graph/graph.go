package graph

import (
	"github.com/RossMerr/Caudex.Graph/graph/vertices"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/storageEngines"
)

// Graph structure consisting of vertices and edges
type Graph struct {
	storageEngines.StorageEngine
}

func (g *Graph) Query(root *vertices.Vertex) *query.Query {
	//todo need to setup channel from DFS or BFS
	c := make(chan *vertices.Vertex)
	return &query.Query{
		Iterate: func() query.Iterator {
			return func() (item *vertices.Vertex, ok bool) {
				v, ok := <-c
				return v, ok
			}
		},
	}

	return nil
}
