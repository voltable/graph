package graphs

import "bitbucket.org/rossmerr/caudex/graphs/internal"

type VertexOperation struct {
	v *Vertex
}

// AddDirectedEdge links two vertex's and returns the edge
func (op *VertexOperation) AddDirectedEdge(to *Vertex) *Edge {
	e := edge{}
	edge := Edge{from: op.v, to: to, edge: &e, isDirected: internal.Directed}
	op.v.edges = append(op.v.edges, edge)
	return &edge
}

// AddEdge links two vertex's and returns the edge
func (op *VertexOperation) AddEdge(to *Vertex) (*Edge, *Edge) {
	e := edge{}
	edge := Edge{from: op.v, to: to, edge: &e, isDirected: internal.Undirected}
	op.v.edges = append(op.v.edges, edge)

	edge2 := Edge{from: to, to: op.v, edge: &e, isDirected: internal.Undirected}
	to.edges = append(to.edges, edge2)
	return &edge, &edge2
}

// RemoveEdge remove a edge
func (op *VertexOperation) RemoveEdge(to *Vertex, label string) {
	if to == nil {
		return
	}

	var isDirected = op.v.remove(label)

	if isDirected == internal.Undirected {
		to.remove(label)
	}
}
