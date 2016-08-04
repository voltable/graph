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
	edge := Edge{from: v, to: to, edge: &e, isDirected: Directed}
	v.edges = append(v.edges, edge)
	return &edge
}

// AddEdge links two vertex's and returns the edge
func (v *Vertex) AddEdge(to *Vertex) (*Edge, *Edge) {
	e := edge{}
	edge := Edge{from: v, to: to, edge: &e, isDirected: Undirected}
	v.edges = append(v.edges, edge)

	edge2 := Edge{from: to, to: v, edge: &e, isDirected: Undirected}
	to.edges = append(to.edges, edge2)
	return &edge, &edge2
}

// RemoveEdge remove a edge
func (v *Vertex) RemoveEdge(to *Vertex, label string) {
	if to == nil {
		return
	}

	var isDirected = v.remove(label)

	if isDirected == Undirected {
		to.remove(label)
	}
}

func (v *Vertex) remove(label string) int {
	for y, edge := range v.edges {
		if edge.Label() == label {
			if edge.to == v {
				c := make([]Edge, len(v.edges)-1)
				v.edges = append(append(c, v.edges[:y]...), v.edges[y+1:]...)
				return edge.isDirected
			}
		}
	}
}
