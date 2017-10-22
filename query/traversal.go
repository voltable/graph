package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/storage"
)

var (
	errPathNotDefine = errors.New("Record Not found")
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
func (t *Traversal) Travers(iterator enumerables.Iterator, path Path) (iteratorFrontier IteratorFrontier, err error) {
	var edgePath *EdgePath
	var vertexPath *VertexPath

	if path == nil {
		return nil, errPathNotDefine
	}

	for p := path.Next(); p != nil; p = p.Next() {
		if pv, ok := p.(*PredicateVertexPath); ok {
			if vertexPath == nil {
				vertexPath = NewVertexPath(iterator, t.storage, pv.Variable)
			}

			edgePath = vertexPath.Node(pv.PredicateVertex)
			iteratorFrontier = edgePath.Iterate

		} else if pe, ok := p.(*PredicateEdgePath); ok {
			vertexPath = edgePath.Relationship(pe.PredicateEdge)
			iteratorFrontier = vertexPath.Iterate
		}
	}

	return
}
