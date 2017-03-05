package traversal

import "github.com/RossMerr/Caudex.Graph/vertices"

type Traversal interface {
	Query(*vertices.Vertex, func(*vertices.Vertex) bool, func(int) bool) []*vertices.Vertex
}
