package graphs

type QueryResult struct {
	Results []*Vertex
}

func NewQueryResult(v []*Vertex) *QueryResult {
	qr := &QueryResult{Results: v}
	return qr
}
