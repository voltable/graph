package graphs

type (
	// An Edge connects two Vertex in a graph.
	Edge struct {
		id         string
		isDirected Digraph
		weight     float32
		label      string
	}

	Edges []*Edge
)

// Weight of a path in a weighted graph
func (e *Edge) Weight() float32 {
	return e.weight
}

// SetWeight sets the edge weight
func (e *Edge) SetWeight(weight float32) {
	e.weight = weight
}

// Label or Type of the edge
func (e *Edge) Label() string {
	return e.label
}

func (e *Edge) SetLabel(label string) {
	e.label = label
}

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].weight > a[j].weight }
