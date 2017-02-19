package graphs

import "errors"

var (
	errEdgeNotFound = errors.New("Edge Not found")
)

// VertexOperation a CRUD operation to perform over a vertex
type VertexOperation struct {
	vertex *Vertex
}

// AddDirectedEdge links two vertex's and returns the edge
func (op *VertexOperation) AddDirectedEdge(to *Vertex) (*Edge, error) {
	edge := Edge{id: to.ID, isDirected: Directed}
	op.vertex.edges[edge.id] = edge
	return &edge, nil
}

// AddEdge links two vertex's and returns the edge
func (op *VertexOperation) AddEdge(to *Vertex) (*Edge, *Edge, error) {
	edge := Edge{id: to.ID, isDirected: Undirected}
	op.vertex.edges[edge.id] = edge

	edge2 := Edge{id: op.vertex.ID, isDirected: Undirected}
	to.edges[edge2.id] = edge2
	return &edge, &edge2, nil
}

// RemoveEdgeByLabel remove a edge
func (op *VertexOperation) RemoveEdgeByLabel(to *Vertex, label string) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := op.vertex.removeRelationshipOnLabel(label)

	if isDirected == Undirected {
		to.removeRelationshipOnLabel(label)
	}

	return nil
}

// RemoveEdge remove a edge
func (op *VertexOperation) RemoveEdge(to *Vertex) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := op.vertex.removeRelationshipsOnVertex(to)

	if isDirected == Undirected {
		to.removeRelationshipsOnVertex(op.vertex)
	}

	return nil
}
