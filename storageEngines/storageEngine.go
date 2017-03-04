package storageEngines

import "github.com/RossMerr/Caudex.Graph/graph"

//StorageEngine structure for saving graph data
type StorageEngine interface {
	Create(...*graph.Vertex) error
	Delete(...*graph.Vertex) error
	Find(string) (*graph.Vertex, error)
	Update(...*graph.Vertex) error
	Close()
}
