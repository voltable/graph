package graphs

import (
	"errors"

	"github.com/hashicorp/go-uuid"
)

var (
	errEdgeNotFound = errors.New("Edge Not found")
)

// VertexOperation a CRUD operation to perform over a vertex
type VertexOperation struct {
	vertex *Vertex
}

// AddDirectedEdge links two vertex's and returns the edge
func (op *VertexOperation) AddDirectedEdge(to *Vertex) (*Edge, error) {
	var id string
	var err error
	if id, err = uuid.GenerateUUID(); err != nil {
		return nil, err
	}

	edge := Edge{id: id, from: op.vertex, to: to, isDirected: Directed}
	op.vertex.edges[edge.id] = edge
	return &edge, nil
}

// AddEdge links two vertex's and returns the edge
func (op *VertexOperation) AddEdge(to *Vertex) (*Edge, *Edge, error) {
	var id1, id2 string
	var err error

	if id1, err = uuid.GenerateUUID(); err != nil {
		return nil, nil, err
	}
	edge := Edge{id: id1, from: op.vertex, to: to, isDirected: Undirected}
	op.vertex.edges[edge.id] = edge

	if id2, err = uuid.GenerateUUID(); err != nil {
		return nil, nil, err
	}
	edge2 := Edge{id: id2, from: to, to: op.vertex, isDirected: Undirected}
	to.edges[edge2.id] = edge2
	return &edge, &edge2, nil
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
