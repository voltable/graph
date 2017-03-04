package traversal

import "github.com/RossMerr/Caudex.Graph/graph/vertices"

type Traversal interface {
	Query(root *vertices.Vertex, fn func(*vertices.Vertex) bool) []*vertices.Vertex
}
