package caudex

// An Edge connects two Vertex in a graph.
type Edge struct {
	from *Vertex
	to   *Vertex
	edge *edge
}

type edge struct {
	Weight float32
	Label  string
}

func (e *Edge) Weight() float32 {
	return e.edge.Weight
}

func (e *Edge) SetWeight(weight float32) {
	e.edge.Weight = weight
}

func (e *Edge) Label() string {
	return e.edge.Label
}
