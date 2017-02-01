package graphs

import "github.com/RossMerr/Caudex.Graph/graphs/internal"

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

func (v *Vertex) remove(label string) internal.Digraph {
	for y, edge := range v.edges {
		if edge.Label() == label {
			if edge.to == v {
				c := make([]Edge, len(v.edges)-1)
				v.edges = append(append(c, v.edges[:y]...), v.edges[y+1:]...)
				return edge.isDirected
			}
		}
	}

	return internal.Undirected
}
