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
func (t *Traversal) Travers(i func() Iterator, path Path) []interface{} {
	edgePath := NewEdgePath(i, t.fetch)
	vertexPath := NewVertexPath(i, t.fetch)
	var result interface{}
	results := make([]interface{}, 0)
	var iterated bool
	for p := path.Next(); p != nil; p = p.Next() {
		if pv, ok := p.(*PredicateVertexPath); ok {
			edgePath = vertexPath.Node(pv.PredicateVertex)
			result, iterated = edgePath.Iterate()()

		} else if pe, ok := p.(*PredicateEdgePath); ok {
			vertexPath = edgePath.Relationship(pe.PredicateEdge)
			result, iterated = vertexPath.Iterate()()
		}
		if iterated {
			results = append(results, result)
		}
	}

	return results
}
