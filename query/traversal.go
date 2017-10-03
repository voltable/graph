package query

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
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
func (t *Traversal) Travers(i enumerables.Iterator, path Path) IteratorFrontier {

	iterated := false
	var result interface{}

	var edgePath *EdgePath
	var vertexPath *VertexPath

	return func() (item *Frontier, ok bool) {
		for p := path.Next(); p != nil; p = p.Next() {

			if pv, ok := p.(*PredicateVertexPath); ok {
				if vertexPath == nil {
					vertexPath = NewVertexPath(i, t.storage, pv.Variable)
				}

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
