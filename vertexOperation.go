package graphs

type VertexOperation struct {
	Vertex *Vertex
}

// AddDirectedEdge links two vertex's and returns the edge
func (op *VertexOperation) AddDirectedEdge(to *Vertex) *Edge {
	edge := Edge{from: op.Vertex, to: to, isDirected: Directed}
	op.Vertex.edges[edge.id] = edge
	return &edge
}

// AddEdge links two vertex's and returns the edge
func (op *VertexOperation) AddEdge(to *Vertex) (*Edge, *Edge) {
	edge := Edge{from: op.Vertex, to: to, isDirected: Undirected}
	op.Vertex.edges[edge.id] = edge

	edge2 := Edge{from: to, to: op.Vertex, isDirected: Undirected}
	to.edges[edge2.id] = edge2
	return &edge, &edge2
}

// RemoveEdge remove a edge
func (op *VertexOperation) RemoveEdge(to *Vertex, label string) {
	if to == nil {
		return
	}

	var isDirected = op.Vertex.remove(label)

	if isDirected == Undirected {
		to.remove(label)
	}
}
