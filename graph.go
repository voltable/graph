package graph

import "github.com/RossMerr/Caudex.Graph/vertices"

//Graph structure for saving graph data
type Graph interface {
	Create(...*vertices.Vertex) error
	Delete(...*vertices.Vertex) error
	Find(string) (*vertices.Vertex, error)
	Update(...*vertices.Vertex) error
	Close()
	Query() []*vertices.Vertex
}
