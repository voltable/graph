package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type QueryResult struct {
	Results []*vertices.Vertex
	Path    Path
}

func NewQueryResult(v []*vertices.Vertex) *QueryResult {
	qr := &QueryResult{Results: v}
	return qr
}

// Path represents a walk through a property graph and consists of a sequence of alternating nodes and relationships.
type Path struct {
}
