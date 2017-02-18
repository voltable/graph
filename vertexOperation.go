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
func (op *VertexOperation) AddDirectedEdge(to *Vertex) *Edge {
	edge := Edge{from: op.vertex, to: to, isDirected: Directed}
	op.vertex.edges[edge.id] = edge
	return &edge
}

// AddEdge links two vertex's and returns the edge
func (op *VertexOperation) AddEdge(to *Vertex) (*Edge, *Edge) {
	edge := Edge{from: op.vertex, to: to, isDirected: Undirected}
	op.vertex.edges[edge.id] = edge

	edge2 := Edge{from: to, to: op.vertex, isDirected: Undirected}
	to.edges[edge2.id] = edge2
	return &edge, &edge2
}

// RemoveEdge remove a edge
func (op *VertexOperation) RemoveEdge(to *Vertex, label string) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := op.vertex.removeRelationshipOnLabel(label)
	if isDirected == Undirected {
		to.removeRelationshipOnLabel(label)
	}

	return nil
}
