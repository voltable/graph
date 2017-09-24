package query

import (
	"github.com/RossMerr/Caudex.Graph/storage"
)

// Traversal decides how to excute the query
type Traversal struct {
	storage storage.Storage
}

// NewTraversal create a Traversal object used to run the query over the graph
func NewTraversal(i storage.Storage) *Traversal {
	return &Traversal{storage: i}
}

// Travers run's the query over the graph and returns a new resulting Iterator
func (t *Traversal) Travers(i IteratorFrontier, path Path) IteratorFrontier {
	edgePath := NewEdgePath(i, t.storage.Fetch())
	vertexPath := NewVertexPath(i, t.storage.Fetch())
	iterated := false
	var result interface{}

	return func() (item *Frontier, ok bool) {
		for p := path.Next(); p != nil; p = p.Next() {
			if pv, ok := p.(*PredicateVertexPath); ok {
				edgePath = vertexPath.Node(pv.PredicateVertex)
				result, iterated = edgePath.Iterate()

			} else if pe, ok := p.(*PredicateEdgePath); ok {
				vertexPath = edgePath.Relationship(pe.PredicateEdge)
				result, iterated = vertexPath.Iterate()
			}
			if iterated {
				if v, is := result.(*Frontier); is {
					return v, true
				}
			}
		}
		return
	}
}
