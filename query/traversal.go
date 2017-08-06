package query

import "github.com/RossMerr/Caudex.Graph/vertices"

// Traversal decides how to excute the query
type Traversal struct {
	fetch func(string) (*vertices.Vertex, error)
}

// NewTraversal create a Traversal object used to run the query over the graph
func NewTraversal(fetch func(string) (*vertices.Vertex, error)) *Traversal {
	return &Traversal{fetch: fetch}
}

// Travers run's the query over the graph
func (t *Traversal) Travers(i func() Iterator, q *Query) error {
	edgePath := NewEdgePath(i, t.fetch)
	vertexPath := NewVertexPath(i, t.fetch)
	var iterate func() Iterator
	p := q.path.Next()
	for p != nil {
		if pv, ok := p.(*PredicateVertexPath); ok {
			edgePath = vertexPath.Node(pv.PredicateVertex)
			iterate = vertexPath.Iterate
		} else if pe, ok := p.(*PredicateEdgePath); ok {
			vertexPath = edgePath.Relationship(pe.PredicateEdge)
			iterate = edgePath.Iterate
		} else {
			break
		}

		p = p.Next()
	}

	iterate()
	return nil
}
