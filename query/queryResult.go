package query

import "github.com/RossMerr/Caudex.Graph/graph"

type QueryResult struct {
	Results []*graph.Vertex
}

func NewQueryResult(v []*graph.Vertex) *QueryResult {
	qr := &QueryResult{Results: v}
	return qr
}
