package graphs

import "bitbucket.org/rossmerr/caudex/graphs/internal"

type VertexOperation struct {
	Vertex *Vertex
	//edges       chan EdgeEnvelop
	//transaction []EdgeEnvelop
}

// AddDirectedEdge links two vertex's and returns the edge
func (op *VertexOperation) AddDirectedEdge(to *Vertex) *Edge {
	e := edge{}
	edge := Edge{from: op.Vertex, to: to, edge: &e, isDirected: internal.Directed}
	op.Vertex.edges = append(op.Vertex.edges, edge)
	return &edge
}

// AddEdge links two vertex's and returns the edge
func (op *VertexOperation) AddEdge(to *Vertex) (*Edge, *Edge) {
	e := edge{}
	edge := Edge{from: op.Vertex, to: to, edge: &e, isDirected: internal.Undirected}
	op.Vertex.edges = append(op.Vertex.edges, edge)

	edge2 := Edge{from: to, to: op.Vertex, edge: &e, isDirected: internal.Undirected}
	to.edges = append(to.edges, edge2)
	return &edge, &edge2
}

// RemoveEdge remove a edge
func (op *VertexOperation) RemoveEdge(to *Vertex, label string) {
	if to == nil {
		return
	}

	var isDirected = op.Vertex.remove(label)

	if isDirected == internal.Undirected {
		to.remove(label)
	}
}
