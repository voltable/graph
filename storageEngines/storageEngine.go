package storageEngines

import "github.com/RossMerr/Caudex.Graph/graph/vertices"

//StorageEngine structure for saving graph data
type StorageEngine interface {
	Create(...*vertices.Vertex) error
	Delete(...*vertices.Vertex) error
	Find(string) (*vertices.Vertex, error)
	Update(...*vertices.Vertex) error
	Close()
}
