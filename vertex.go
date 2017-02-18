package graphs

// Vertex .
type Vertex struct {
	ID    string
	edges map[string]Edge
	Value *interface{}
	label string
}

// Label vertex label type
func (v *Vertex) Label() string {
	return v.label
}

func (v *Vertex) removeRelationshipOnLabel(label string) Digraph {
	return v.removeRelationshipsF(func(e Edge) bool {
		return e.label == label
	})
}

func (v *Vertex) removeRelationships() {
	v.removeRelationshipsF(func(e Edge) bool {
		return true
	})
}

func (v *Vertex) removeRelationshipsF(f func(e Edge) bool) Digraph {
	for _, edge := range v.edges {
		if f(edge) {
			if edge.to == v {
				delete(edge.to.edges, edge.id)
				return edge.isDirected
			}
		}
	}
	return Undirected
}
