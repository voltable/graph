package caudex

// Vertex .
type Vertex struct {
	ID    string
	edges []Edge
	Value *interface{}
	Label string
}

// AddDirectedEdge links two vertex's and returns the edge
func AddDirectedEdge(from *Vertex, to *Vertex) *Edge {
	e := edge{}
	edge := Edge{from: from, to: to, edge: &e}
	from.edges = append(from.edges, edge)
	return &edge
}

// AddEdge links two vertex's and returns the edge
func AddEdge(from *Vertex, to *Vertex) *Edge {
	e := edge{}
	edge := Edge{from: from, to: to, edge: &e}
	from.edges = append(from.edges, edge)

	edge2 := Edge{from: to, to: from, edge: &e}
	to.edges = append(to.edges, edge2)
	return &edge
}

// RemoveEdge remove a edge
func RemoveEdge(from *Vertex, to *Vertex) {
	fromEdges := from.edges
	toEdges := to.edges

	for e := range fromEdges {
		if fromEdges[e].to == to {
			remove(e, &fromEdges)
			break
		}
	}

	for e := range toEdges {
		if toEdges[e].to == to {
			remove(e, &toEdges)
			break
		}
	}
}

func remove(remove int, edges *[]Edge) {
	(*edges)[remove], (*edges)[len(*edges)-1] = (*edges)[len(*edges)-1], (*edges)[remove]
	*edges = (*edges)[:len(*edges)-1]
}
