package query

import "github.com/RossMerr/Caudex.Graph/graph/vertices"

type QueryResult struct {
	Results []*vertices.Vertex
}

func NewQueryResult(v []*vertices.Vertex) *QueryResult {
	qr := &QueryResult{Results: v}
	return qr
}
