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

func Weight(e Edge) float32 {
	return e.edge.Weight
}

func SetWeight(e Edge, weight float32) {
	e.edge.Weight = weight
}

func Label(e Edge) string {
	return e.edge.Label
}
