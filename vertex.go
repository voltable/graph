package graphs

// Vertex .
type Vertex struct {
	ID    string
	edges map[string]Edge
	Value interface{}
	label string
}

// Label vertex label type
func (v *Vertex) Label() string {
	return v.label
}

func (v *Vertex) Edges() []Edge {
	a := make([]Edge, 0, len(v.edges))

	for _, value := range v.edges {
		a = append(a, value)
	}

	return a
}

func (v *Vertex) removeRelationshipOnLabel(label string) Digraph {
	return v.removeRelationshipsF(func(id string, e Edge) bool {
		return e.label == label
	})
}

func (v *Vertex) removeRelationships() {
	v.removeRelationshipsF(func(id string, e Edge) bool {
		return true
	})
}

func (v *Vertex) removeRelationshipsOnVertex(to *Vertex) Digraph {
	return v.removeRelationshipsF(func(id string, e Edge) bool {
		return id == to.ID
	})
}

func (v *Vertex) removeRelationshipsF(f func(id string, e Edge) bool) Digraph {
	for id, edge := range v.edges {
		if f(id, edge) {
			delete(v.edges, edge.id)
			return edge.isDirected
		}
	}
	return Undirected
}
