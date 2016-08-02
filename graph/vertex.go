package caudex

// Vertex .
type Vertex struct {
	ID    string
	edges []Edge
	Value *interface{}
	label string
}

// Label vertex label type
func (v *Vertex) Label() string {
	return v.label
}

// AddDirectedEdge links two vertex's and returns the edge
func (v *Vertex) AddDirectedEdge(to *Vertex) *Edge {
	e := edge{}
	edge := Edge{from: v, to: to, edge: &e}
	v.edges = append(v.edges, edge)
	return &edge
}

// AddEdge links two vertex's and returns the edge
func (v *Vertex) AddEdge(to *Vertex) (*Edge, *Edge) {
	e := edge{}
	edge := Edge{from: v, to: to, edge: &e}
	v.edges = append(v.edges, edge)

	edge2 := Edge{from: to, to: v, edge: &e}
	to.edges = append(to.edges, edge2)
	return &edge, &edge2
}

// RemoveEdge remove a edge
func (v *Vertex) RemoveEdge(to *Vertex, label string) {
	if to == nil {
		return
	}

	fromEdges := v.edges
	toEdges := to.edges

	for e := range fromEdges {
		if fromEdges[e].to == to && fromEdges[e].edge.Label == label {
			remove(e, &fromEdges)
			break
		}
	}

	for e := range toEdges {
		if toEdges[e].to == to && toEdges[e].edge.Label == label {
			remove(e, &toEdges)
			break
		}
	}
}

func remove(remove int, edges *[]Edge) {
	(*edges)[remove], (*edges)[len(*edges)-1] = (*edges)[len(*edges)-1], (*edges)[remove]
	*edges = (*edges)[:len(*edges)-1]
}
